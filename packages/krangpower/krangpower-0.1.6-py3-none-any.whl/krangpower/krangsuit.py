import copy
import json
import re
import xml.etree.ElementTree as ET
from csv import reader as csvreader
from functools import singledispatch as _singledispatch

import networkx as nx
import numpy as np
import pandas
from tqdm import tqdm as _tqdm

import krangpower.components
from krangpower import busquery as bq
from krangpower import components as co
from krangpower.aux_fcn import get_help_out
from krangpower.config_loader import _PINT_QTY_TYPE, _ELK, _DEFAULT_KRANG_NAME, _CMD_LOG_NEWLINE_LEN, UM, DSSHELP, \
    _TMP_PATH
from krangpower.enhancer import OpendssdirectEnhancer
from krangpower.logging_init import _clog

__all__ = ['Krang', 'from_json']


def _helpfun(config, section):
    """This decorator adds a 'help' submethod to a function or method, and is meant to read a series of help (
    property, desc) pairs from a config file. The help method is invoked by function.help(), just prints stuff to the
    console and returns nothing. """

    def _real_helpfun(f):

        def ghelp():
            print('\nPARAMETERS HELP FOR FUNCTION {0} ({1})\n'.format(f.__name__, section))
            print(get_help_out(config, section))

        f.help = ghelp
        return f

    return _real_helpfun


class Krang:
    """The Krang is the main class of krangpower. It has facilities to drive the OpenDSS engine, retrieve and query
    elements and options, saving and loading circuit data as json, representing the circuit as graph, perform repeated
    solutions and reporting the results as DataFrames.

        >>> myKrang = Krang('myckt', co.Vsource(basekv=15.0 * UM.kV).aka('source'))
        >>> UM = myKrang.get_unit_register()
        >>> myKrang.set(number=20, stepsize=15 * UM.min)
        >>> myKrang['sourcebus', 'a'] << co.Line(length=120 * UM.unitlength).aka('line_1')
        <BusView('sourcebus', 'a')>
        >>> l2 = co.Line(units='m', length=0.72 * UM.m).aka('line_2')
        >>> myKrang[('a.1.3.2', 'b.3.2.1')] << l2
        <BusView('a', 'b')>
        >>> myKrang['line.line_2']['length'] = 200 * UM.yard
        >>> vls = np.matrix([15.0, 7.0]) * UM.kV
        >>> trf = co.Transformer(windings=2, kvs=vls).aka('trans')
        >>> myKrang[('b', 'c')] << trf
        <BusView('b', 'c')>
        >>> print(myKrang['line.line_2']['rmatrix'])
        [[ 0.09813333  0.04013333  0.04013333] [ 0.04013333  0.09813333  0.04013333] [ 0.04013333  0.04013333  0.09813333]] ohm / meter
        >>> bp_load = co.Load(kw=18.2345 * UM.kW, kv=15 * UM.kV)  # * lish
        >>> myKrang[('a',)] << bp_load.aka('load_a')
        <BusView('a',)>
        >>> myKrang[('b',)] << bp_load.aka('load_b')
        <BusView('b',)>
        >>> u_load_a= myKrang['load_a'].unpack(verbose=False)
        >>> myKrang.command('makebuslist')
        ''
        >>> i, v = myKrang.drag_solve()
        >>> np.isclose(i['a.1'][0].to('V').magnitude, 8672.156356-30.933278j)
        True
        >>> myKrang.save_json(r'.\krang.json')
        >>> myKrang.save_dss(r'.\krang.dss')
        >>> cs = from_json(r'.\krang.json')
        >>> -cs['line.line_2']
        >>> i, v = cs.drag_solve()
        >>> print(i['b.1'][0])
        0j volt
    """

    def __init__(self, *args):
        self.id = _DEFAULT_KRANG_NAME
        self._up_to_date = False
        self._last_gr = None
        self._named_entities = []
        self._ai_list = []
        self.com = ''
        self.brain = None
        self._coords_linked = dict()
        self._initialize(*args)

    def _initialize(self, name=_DEFAULT_KRANG_NAME, vsource=co.Vsource(), source_bus_name='sourcebus'):
        self.brain = OpendssdirectEnhancer(oe_id=name)
        _clog.debug('\n' + '$%&' * _CMD_LOG_NEWLINE_LEN)
        self.id = name
        self.command('clear')
        master_string = self._form_newcircuit_string(name, vsource, source_bus_name)
        self.command(master_string)
        self.set(mode='duty')
        self.command('makebuslist')  # in order to make sourcebus recognizable
        self.brain.Basic.DataPath(_TMP_PATH)  # redirects all file output to the temp folder

    @staticmethod
    def get_unit_registry():
        """Retrieves krangpower's UnitRegistry."""
        return UM

    @staticmethod
    def _form_newcircuit_string(name, vsource, source_bus_name):
        master_string = 'new object = circuit.' + name + ' '
        vsource.name = 'source'  # we force the name, so an already named vsource can be used as source.
        main_source_dec = vsource.fcs(buses=(source_bus_name, source_bus_name + '.0.0.0'))
        main_dec = re.sub('New vsource\.source ', '', main_source_dec)
        main_dec = re.sub('bus2=[^ ]+ ', '', main_dec)
        master_string += ' ' + main_dec + '\n'
        return master_string

    def __getitem__(self, item):
        """Krang['bus.nname'], Krang['load.load1'], Krang['load1'] gets the component or a list of components of the bus;
        Krang[('bus.nname',)], Krang['bus.b1','bus.b2'] gets a BusView to which you can add components with <<.
        Note that in order to get a BusView of a single bus, one needs to explicitly pack it in a tuple.
        """
        # implemented outside for rigid single dispatching between strings and tuples
        return _oe_getitem(item, self)

    def __getattr__(self, item):
        """Krang.item, aside from retrieving the builtin attributes, wraps by default the calls to opendssdirect's
        'class_to_dataframe' utility function. These are accessible via capital letter calls. Both singular and plural
        are accepted. (e.g., 'Line' or 'Lines', 'RegControl' , 'Regcontrol', 'RegControls', but not 'transformers')"""
        try:
            assert item[0].isupper()
            dep_item = re.sub('s$', '', item).lower()
            return self.brain.utils.class_to_dataframe(dep_item)
        except (AssertionError, NotImplementedError):
            raise AttributeError('{0} is neither a valid attribute nor a valid identifier for the class-views.'.format(
                item
            ))

    def __lshift__(self, other):
        """The left-shift operator << adds components to the Krang. Components that can and have to be directly added
        to the krang are those that represent data (WireData, TSData, LineGeometry...) or those that are in the circuit
        but above the topology (such as Monitors and all the Controllers).
        The names of these elements must not be blank when they are added."""
        try:
            assert other.isnamed()
            self._named_entities.append(other)
        except AssertionError:
            try:
                assert other.isabove()
            except AssertionError:
                raise TypeError('The object could not be directly added to the Krang. Is it a bus object?')

        try:
            assert other.name != ''
        except AssertionError:
            raise ValueError('Tried to add an object with a blank name')

        self.command(other.fcs())
        self.brain._names_up2date = False
        return self

    def __bool__(self):
        return self._up_to_date

    # def start_console(self):
    #
    #     import sys
    #     sys.ps1 = '<<>'
    #     name = self.name
    #
    #     class KrangSole(code.InteractiveConsole):
    #
    #         def __init__(self, locals=None, filename="<console>"):
    #             super().__init__(locals, filename)
    #
    #         def interact(self, banner=None, exitmsg=None):
    #             """Closely emulate the interactive Python console.
    #
    #             The optional banner argument specifies the banner to print
    #             before the first interaction; by default it prints a banner
    #             similar to the one printed by the real Python interpreter,
    #             followed by the current class name in parentheses (so as not
    #             to confuse this with the real interpreter -- since it's so
    #             close!).
    #
    #             The optional exitmsg argument specifies the exit message
    #             printed when exiting. Pass the empty string to suppress
    #             printing an exit message. If exitmsg is not given or None,
    #             a default message is printed.
    #
    #             """
    #             banner = 'aye aye!'
    #             prompt1 = '[krang]>'
    #             prompt2 = '[krang]...'
    #             self.write("%s\n" % str(banner))
    #             more = 0
    #             while 1:
    #                 try:
    #                     if more:
    #                         prompt = prompt1
    #                     else:
    #                         prompt = prompt2
    #                     try:
    #                         line = self.raw_input(prompt)
    #                     except EOFError:
    #                         self.write("\n")
    #                         break
    #                 except KeyboardInterrupt:
    #                     self.write("\nKeyboardInterrupt\n")
    #                     self.resetbuffer()
    #                     more = 0
    #                 else:
    #                     more = self.push(line)
    #             if exitmsg is None:
    #                 self.write('now exiting %s...\n' % self.__class__.__name__)
    #             elif exitmsg != '':
    #                 self.write('%s\n' % exitmsg)
    #
    #         def push(self, line):
    #             if line.startswith('['):
    #                 super().push(name + line)
    #             else:
    #                 super().push(name + '.' + line)
    #
    #     conzol = KrangSole(locals={self.name: self, 'load': co.Load})
    #     conzol.interact()

    @_helpfun(DSSHELP, 'EXECUTIVE')
    def command(self, cmd_str: str, echo=True):
        """Performs an opendss textual command and adds the commands to the record Krang.com if echo is True."""

        rslt = self.brain.txt_command(cmd_str, echo)
        self._up_to_date = False
        if echo:
            self.com += cmd_str + '\n'
        return rslt

    @_helpfun(DSSHELP, 'OPTIONS')
    def set(self, **opts_vals):
        """Sets circuit options according to a dict. Option that have a physical dimensionality (such as stepsize) can
        be specified as pint quantities; otherwise, the default opendss units will be used."""
        for option, value in opts_vals.items():
            if isinstance(value, _PINT_QTY_TYPE):
                vl = value.to(UM.parse_units(co.DEFAULT_SETTINGS['units'][option])).magnitude
            else:
                vl = value
            self.command('set {0}={1}'.format(option, vl))
        self._up_to_date = False

    @_helpfun(DSSHELP, 'OPTIONS')
    def get(self, *opts):
        """Takes a list of circuit options and returns them as dict of name, value."""
        assert all([x in list(co.DEFAULT_SETTINGS['values'].keys()) for x in opts])
        r_opts = {opt: self.command('get {0}'.format(opt), echo=False).split('!')[0] for opt in opts}

        for op in r_opts.keys():
            tt = type(co.DEFAULT_SETTINGS['values'][op])
            if tt is list:
                r_opts[op] = eval(r_opts[op])  # lists are literals like [13]
            else:
                r_opts[op] = tt(r_opts[op])
            if op in co.DEFAULT_SETTINGS['units'].keys():
                r_opts[op] *= UM.parse_units(co.DEFAULT_SETTINGS['units'][op])

        return r_opts

    def snap(self):
        """Solves a circuit snapshot."""
        self.set(mode='snap')
        self.solve()
        self.set(mode='duty')

    @_helpfun(DSSHELP, 'EXPORT')
    def export(self, object_descriptor):
        """Wraps the OpenDSS export command. Returns a DataFrame for csv exports, an ElementTree for xml exports.
        Has a help attribute."""
        tmp_filename = self.command('export ' + object_descriptor)
        if tmp_filename.lower().endswith('csv'):
            return pandas.read_csv(tmp_filename)
        elif tmp_filename.lower().endswith('xml'):
            return ET.parse(tmp_filename)
        else:
            raise ValueError('Unknown format for export file {0}, contact the developer'.format(tmp_filename))

    @_helpfun(DSSHELP, 'PLOT')
    def plot(self, object_descriptor):
        """Wraps the OpenDSS export command. Has a help attribute."""
        self.command('plot ' + object_descriptor)

    @_helpfun(DSSHELP, 'SHOW')
    def show(self, object_descriptor):
        """Wraps the OpenDSS show command. Has a help attribute."""
        self.command('show ' + object_descriptor)
        # tmp_filename = self.command('show ' + object_descriptor)
        # with open(tmp_filename) as show_file:
        #     content = show_file.read()
        #
        # if rtrn:
        #     return content
        # else:
        #     print(content)

    def drag_solve(self):
        """Instead of launching a monolithic duty solve, This command solves one step at a time and saves node currents
        and voltages in the two DataFrames returned."""
        nmbr = self.brain.Solution.Number()
        self.brain.Solution.Number(1)
        v = pandas.DataFrame(
            columns=[x.lower() for x in self.brain.Circuit.YNodeOrder()])
        i = pandas.DataFrame(
            columns=[x.lower() for x in self.brain.Circuit.YNodeOrder()])

        self.brain.log_line('Commencing drag_solve of {0} points: the individual "solve" commands will be omitted.'
                            ' Wait for end message...'.format(nmbr))
        for _ in _tqdm(range(nmbr)):
            for ai_el in self._ai_list:
                self.command(ai_el.element.fus(self, ai_el.name))

            self.solve(echo=False)
            v = v.append(self.brain.Circuit.YNodeVArray(), ignore_index=True)
            i = i.append(self.brain.Circuit.YCurrents(), ignore_index=True)

        self.brain.log_line('Drag_solve ended')
        self.brain.Solution.Number(nmbr)
        self._up_to_date = True

        return v, i

    def solve(self, echo=True):
        """Imparts the solve command to OpenDSS."""
        self._declare_buscoords()
        self.command('solve', echo)
        self._up_to_date = True

    def _declare_buscoords(self):
        """Meant to be called just before solve, so that all buses are already mentioned"""
        self.command('makebuslist')
        for busname, coords in self._coords_linked.items():
            if coords is not None:
                try:
                    self[busname].X(coords[0])
                    self[busname].Y(coords[1])
                except KeyError:
                    continue  # we ignore buses present in coords linked, but not generated
            else:
                continue

    def _preload_buscoords(self, path):
        """Loads in a local dict the buscoords from path."""
        with open(path, 'r') as bc_file:
            bcr = csvreader(bc_file)
            try:
                while True:
                    try:
                        row = next(bcr)
                        float(row[1])
                        break
                    except ValueError:
                        continue
            except StopIteration:
                return
            while row:
                self._coords_linked[str(row[0])] = [float(row[1]), float(row[2])]
                try:
                    row = next(bcr)
                except StopIteration:
                    return

    def link_coords(self, csv_path):
        """Notifies Krang of a csv file with bus coordinates to use."""
        # todo sanity check
        self._preload_buscoords(csv_path)

    def make_json_dict(self):
        """Returns a complete description of the circuit and its objects as a json.dumpable-dict."""
        master_dict = {'cktname': self.name, 'elements': {}, 'settings': {}}

        # elements
        for Nm in _tqdm(self.brain.Circuit.AllElementNames()):
            nm = Nm.lower()
            master_dict['elements'][nm] = self[nm].unpack().jsonize()
            master_dict['elements'][nm]['topological'] = self[nm].topological

        for ne in _tqdm(self._named_entities):
            master_dict['elements'][ne.fullname] = ne.jsonize()

        # options
        opts = self.get(*list(co.DEFAULT_SETTINGS['values'].keys()))
        for on, ov in _tqdm(opts.items()):
            if isinstance(ov, _PINT_QTY_TYPE):
                opts[on] = ov.to(UM.parse_units(co.DEFAULT_SETTINGS['units'][on])).magnitude
                if isinstance(opts[on], (np.ndarray, np.matrix)):
                    opts[on] = opts[on].tolist()

        master_dict['settings']['values'] = opts
        master_dict['settings']['units'] = co.DEFAULT_SETTINGS['units']

        # coordinates
        master_dict['buscoords'] = self.bus_coords

        return master_dict

    def save_json(self, path):
        """Saves in path a complete description of the circuit and its objects"""
        with open(path, 'w') as ofile:
            json.dump(self.make_json_dict(), ofile, indent=4)

    def save_dss(self, path):
        """Saves a file with the text commands that were imparted by the Krang.command method aside from those for which
        echo was False. The file output should be loadable and runnable in traditional OpenDSS with no modifications."""
        with open(path, 'w') as ofile:
            ofile.write(self.com)

    @property
    def name(self):
        """The name of the circuit.."""
        return self.brain.Circuit.Name()

    @property
    def graph(self):
        """Krang.graph is a Networkx.Graph that contains a description of the circuit. The elements are stored as
        _PackedOpendssElement's in the edge/node property '{0}'""".format(_ELK)

        def _update_node(self, gr, bs, name):
            try:
                exel = gr.nodes[bs][_ELK]
            except KeyError:
                gr.add_node(bs, **{_ELK: [self.brain[name]]})
                return
            exel.append(self.brain[name])
            return

        def _update_edge(self, gr, ed, name):
            try:
                exel = gr.edges[ed][_ELK]
            except KeyError:
                gr.add_edge(*ed, **{_ELK: [self.brain[name]]})
                return
            exel.append(self.brain[name])
            return

        if self._up_to_date:
            return self._last_gr
        else:
            gr = nx.Graph()
            ns = self.brain.Circuit.AllElementNames()
            for name in ns:
                try:
                    buses = self.brain[name].BusNames()
                except TypeError:
                    continue

                # todo encode term perms in the graph
                if len(buses) == 1:
                    bs, _ = self._bus_resolve(buses[0])

                    _update_node(self, gr, bs, name)

                    # gr.add_node(bs, **{_elk: self.oe[name]})
                elif len(buses) == 2:
                    bs0, _ = self._bus_resolve(buses[0])
                    bs1, _ = self._bus_resolve(buses[1])

                    _update_edge(self, gr, (bs0, bs1), name)

                    # gr.add_edge(bs0, bs1, **{_elk: self.oe[name]})
                else:
                    raise IndexError('Buses were > 2. This is a big problem.')

            self._last_gr = gr
        return gr

    @property
    def bus_coords(self):
        """Returns a dict with the bus coordinates already loaded in Opendss. Beware that coordinates loaded through
        a link_kml or link_csv are not immediately loaded in the Opendss, but they are just before a solution is
        launched. """
        bp = {}
        for bn in self.brain.Circuit.AllBusNames():
            if self.brain['bus.' + bn].Coorddefined():
                bp[bn] = (self.brain[bn].X(), self.brain[bn].Y())
            else:
                bp[bn] = None
        return bp

    @staticmethod
    def _bus_resolve(bus_descriptor: str):
        """
        >>> Krang()._bus_resolve('bus2.3.1.2')
        ('bus2', (3, 1, 2))
        >>> Krang()._bus_resolve('bus2.33.14.12323.2.3.3')
        ('bus2', (33, 14, 12323, 2, 3, 3))
        """

        bus_descriptor.replace('bus.', '')
        tkns = bus_descriptor.split('.')

        bus = tkns[0]
        terminals = tuple(int(x) for x in tkns[1:])

        return bus, terminals


@_singledispatch
def _oe_getitem(item, oeshell):
    # no default implementation
    raise TypeError('Invalid identificator passed. You can specify fully qualified element names as str, or bus/'
                    'couples of buses as tuples of str.')


@_oe_getitem.register(str)
def _(item, krg):
    return krg.brain[item]


@_oe_getitem.register(tuple)
def _(item, krg):
    assert len(item) <= 2
    bustermtuples = map(krg._bus_resolve, item)
    return _BusView(krg, list(bustermtuples))


class _DepGraph(nx.DiGraph):
    """Simple extension of nx.Digraph created to reverse-walk a dependency branching in order to declare the entities
    in the right order."""

    @property
    def leaves(self):
        return [x for x in self.nodes() if self.out_degree(x) == 0]

    def trim(self):
        self.remove_nodes_from(self.leaves)

    def recursive_prune(self):

        # dependency non-circularity check
        cy = list(nx.simple_cycles(self))
        try:
            assert len(cy) == 0
        except AssertionError:
            print('Found circular dependency(s) in the graph!')
            for cycle in cy:
                print(cycle)
            raise ValueError('Cannot recursively prune a directed graph with cycles')

        # actual prune generator
        while self.leaves:
            yield self.leaves
            self.trim()


class _BusView:
    """_BusView is meant to be instantiated only by Krang and is returned by indicization like Krang['bus1', 'bus2'] or
    Krang['bus1']. The main uses of a _BusView are:

        - Adding elements to the corresponding bus/edge:
            Krang['bus1'] << kp.Load()  # adds a Load to bus1
            Krang['bus1', 'bus2'] << kp.Line()  # adds a Line between bus1 and bus2

        - Evaluating a function from the submodule 'busquery' on the corresponding bus/edge through attribute resolution:
            Krang['bus1'].voltage
            Krang['bus1'].totload

        - Getting a _PackedOpendssElement from the ones pertaining the corresponding bus/edge:
            Krang['bus1']['myload']
    """

    def __init__(self, oek: Krang, bustermtuples):
        self.btt = bustermtuples
        self.tp = dict(bustermtuples)
        self.buses = tuple(self.tp.keys())
        self.nb = len(self.buses)
        self.oek = oek
        self._content = None

        buskey = 'buses'
        tkey = 'terminals'

        self.fcs_kwargs = {buskey: self.buses, tkey: self.tp}

    def __lshift__(self, other):
        """Adds a component to the BusView, binding the component added to the buses referenced by the BusView."""
        assert not other.isnamed()
        try:
            assert other.name != ''
        except AssertionError:
            raise ValueError('Did you name the element before adding it?')
        self.oek.command(other.fcs(**self.fcs_kwargs))

        # remember ai elements
        if other.isai():
            self.oek._ai_list.append(other)

        return self

    @property
    def content(self):
        """A list containing the PackedOpendssElements bound to the BusView's buses."""
        if self._content is None:
            # CONTENT IS NOT UPDATED AFTER FIRST EVAL
            if len(self.buses) == 1:
                try:
                    self._content = self.oek.graph.nodes[self.buses[0]][_ELK]
                    # it's ok that a KeyError by both indexes is caught in the same way
                except KeyError:
                    self._content = []
            elif len(self.buses) == 2:
                self._content = list(nx.get_edge_attributes(self.oek.graph.subgraph(self.buses), _ELK).values())
            else:
                raise ValueError
        return self._content

    def __getattr__(self, item):
        """Calculates a quantity from the submodule busquery on the BusView."""
        try:
            # attributes requested via getattr are searched in busquery
            f = bq.get_fun(item)
        except KeyError:
            raise AttributeError('Attribute/query function {0} is not implemented')

        if self.nb == 1:
            return f(self.oek, self, self.buses[0])
        elif self.nb == 2:
            return f(self.oek, self, self.buses)
        else:
            raise AttributeError

    def __str__(self):
        return '<BusView' + str(self.buses) + '>'

    def __repr__(self):
        return self.__str__()


def from_json(path):
    """Loads circuit data from a json structured like the ones returned by Krang.save_json. Declaration precedence due
    to dependency between object is automatically taken care of."""

    # load all entities
    with open(path, 'r') as ofile:
        master_dict = json.load(ofile)

    # init the krang with the source, then remove it
    l_ckt = Krang(master_dict['cktname'], krangpower.components.dejsonize(master_dict['elements']['vsource.source']))
    del master_dict['elements']['vsource.source']

    # load and declare options
    opt_dict = master_dict['settings']
    for on, ov in opt_dict['values'].items():
        # we try in any way to see if the value is the same as the default and, if so, we continue
        # todo there is an edge case where the value of a measured quantity is the same, but the unit is different
        if ov == co.DEFAULT_SETTINGS['values'][on]:
            continue
        try:
            if np.isclose(ov, co.DEFAULT_SETTINGS['values'][on]):
                continue
        except ValueError:
            try:
                if np.isclose(ov, co.DEFAULT_SETTINGS['values'][on]).all():
                    continue
            except ValueError:
                pass
        except TypeError:
            pass

        try:
            if ov.lower() == co.DEFAULT_SETTINGS['values'][on].lower():
                continue
        except AttributeError:
            pass

        if on in opt_dict['units'].keys():
            d_ov = ov * UM.parse_units(opt_dict['units'][on])
        else:
            d_ov = ov

        l_ckt.set(**{on: d_ov})

    # reconstruction of dependency graph and declarations
    dep_graph = _DepGraph()
    for jobj in master_dict['elements'].values():
        vname = jobj['type'].split('_')[0] + '.' + jobj['name']
        # if the element has no dependencies, we just add a node with iths name
        if jobj['depends'] == {} or all([d == '' for d in jobj['depends'].values()]):
            dep_graph.add_node(vname)
        else:
            # if an element parameter depends on another name, or a list of other names, we create all the edges
            # necessary
            for dvalue in jobj['depends'].values():
                if isinstance(dvalue, list):
                    for dv in dvalue:
                        if dv != '':
                            dep_graph.add_edge(vname, dv)
                else:
                    if dvalue != '':
                        dep_graph.add_edge(vname, dvalue)

    # we cyclically consider all "leaves", add the objects at the leaves, then trim the leaves and go on with
    # the new leaves.
    # In this way we are sure that, whenever a name is mentioned in a fcs, its entity was already declared.
    for trimmed_leaves in dep_graph.recursive_prune():
        for nm in trimmed_leaves:
            try:
                jobj = copy.deepcopy(master_dict['elements'][nm.lower()])
            except KeyError:
                mdmod = {k.split('.')[1]: v for k, v in master_dict['elements'].items()}
                jobj = copy.deepcopy(mdmod[nm.lower()])
            dssobj = krangpower.components.dejsonize(jobj)
            if dssobj.isnamed():
                l_ckt << dssobj
            elif dssobj.isabove():
                l_ckt << dssobj.aka(jobj['name'])
            else:
                l_ckt[tuple(jobj['topological'])] << dssobj.aka(jobj['name'])
                # l_ckt.command(dssobj.aka(jobj['name']).fcs(buses=jobj['topological']))

    l_ckt._coords_linked = master_dict['buscoords']

    return l_ckt


def _main():
    pass


if __name__ == '__main__':
    _main()
