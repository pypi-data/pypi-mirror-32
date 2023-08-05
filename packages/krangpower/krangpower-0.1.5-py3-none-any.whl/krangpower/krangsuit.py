import copy
import json
import re
from functools import singledispatch as _singledispatch


import networkx as nx
import numpy as np
import pandas
from tqdm import tqdm as _tqdm

from . import busquery as bq
from . import components as co
from .components import um, _pint_qty_type
from .enhancer import OpendssdirectEnhancer
from .enhancer import _clog
from . import aux_fcn as au

__all__ = ['Krang', 'from_json']
_elk = 'el'
_default_krang_name = 'Unnamed_Krang'
_cmd_log_newline_len = 60


class Krang:

    def __init__(self):
        self.id = _default_krang_name
        self.up_to_date = False
        self.last_gr = None
        self.named_entities = []
        self.ai_list = []
        self.com = ''
        self.brain = None

    def initialize(self, name=_default_krang_name, vsource=co.Vsource()):
        self.brain = OpendssdirectEnhancer(id=name)
        _clog.debug('\n' + '$%&' * _cmd_log_newline_len)
        self.command('clear')
        self.id = name
        master_string = 'new object = circuit.' + name + ' '
        vsource.name = 'source'  # we force the name, so an already named vsource can be used as source.
        main_source_dec = vsource.fcs(buses=('sourcebus', 'sourcebus.0.0.0'))
        main_dec = re.sub('New vsource\.source ', '', main_source_dec)
        main_dec = re.sub('bus2=[^ ]+ ', '', main_dec)
        master_string += ' ' + main_dec + '\n'

        self.command(master_string)
        self.set(mode='duty')
        self.command('makebuslist')

    def __getitem__(self, item):
        # OEShell['bus.nname'] gets the component or a list of components of the bus
        # OEShell[('bus.nname',)] gets a view
        # implemented outside for single dispatching
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
        try:
            assert other.isnamed()
            self.named_entities.append(other)
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
        return self

    def __bool__(self):
        return self.up_to_date

    def command(self, cmd_str: str, echo=True):

        rslt = self.brain.txt_command(cmd_str, echo)
        if echo:
            self.com += cmd_str + '\n'

        return rslt

    def set(self, **opts_vals):
        for option, value in opts_vals.items():
            if isinstance(value, _pint_qty_type):
                vl = value.to(um.parse_units(co.default_settings['units'][option])).magnitude
            else:
                vl = value
            self.command('set {0}={1}'.format(option, vl))

    def get(self, *opts):

        assert all([x in list(co.default_settings['values'].keys()) for x in opts])
        r_opts = {opt: self.command('get {0}'.format(opt), echo=False).split('!')[0] for opt in opts}

        for op in r_opts.keys():
            tt = type(co.default_settings['values'][op])
            if tt is list:
                r_opts[op] = eval(r_opts[op])  # lists are literals like [13]
            else:
                r_opts[op] = tt(r_opts[op])
            if op in co.default_settings['units'].keys():
                r_opts[op] *= um.parse_units(co.default_settings['units'][op])

        return r_opts

    def snap(self):
        self.command('set mode=snap\nsolve\nset mode=duty')

    def drag_solve(self):
        nmbr = self.brain.Solution.Number()
        self.brain.Solution.Number(1)
        v = pandas.DataFrame(
            columns=[x.lower() for x in self.brain.Circuit.YNodeOrder()])
        i = pandas.DataFrame(
            columns=[x.lower() for x in self.brain.Circuit.YNodeOrder()])

        self.brain.log_line('Commencing drag_solve of {0} points: the individual "solve" commands will be omitted.'
                            ' Wait for end message...'.format(nmbr))
        for _ in _tqdm(range(nmbr)):
            for ai_el in self.ai_list:
                self.command(ai_el.element.fus(self, ai_el.name))

            self.command('solve', echo=False)
            v = v.append(self.brain.Circuit.YNodeVArray(), ignore_index=True)
            i = i.append(self.brain.Circuit.YCurrents(), ignore_index=True)

        self.brain.log_line('Drag_solve ended')
        self.brain.Solution.Number(nmbr)

        return v, i

    def solve(self):
        self.command('solve')

    def make_json_dict(self):
        master_dict = {'cktname': self.name, 'elements': {}, 'settings': {}}

        for Nm in self.brain.Circuit.AllElementNames():
            nm = Nm.lower()
            master_dict['elements'][nm] = self[nm].unpack().jsonize()
            master_dict['elements'][nm]['topological'] = self[nm].topological

        for ne in self.named_entities:
            master_dict['elements'][ne.fullname] = ne.jsonize()

        # options
        opts = self.get(*list(co.default_settings['values'].keys()))
        for on, ov in opts.items():
            if isinstance(ov, _pint_qty_type):
                opts[on] = ov.to(um.parse_units(co.default_settings['units'][on])).magnitude
                if isinstance(opts[on], (np.ndarray, np.matrix)):
                    opts[on] = opts[on].tolist()

        master_dict['settings']['values'] = opts
        master_dict['settings']['units'] = co.default_settings['units']

        return master_dict

    def save_json(self, path):
        with open(path, 'w') as ofile:
            json.dump(self.make_json_dict(), ofile, indent=4)

    def save_dss(self, path):
        with open(path, 'w') as ofile:
            ofile.write(self.com)

    @property
    def name(self):
        return self.brain.Circuit.Name()

    @property
    def graph(self):

        def _update_node(self, gr, bs, name):
            try:
                exel = gr.nodes[bs][_elk]
            except KeyError:
                gr.add_node(bs, **{_elk: [self.brain[name]]})
                return
            exel.append(self.brain[name])
            return

        def _update_edge(self, gr, ed, name):
            try:
                exel = gr.edges[ed][_elk]
            except KeyError:
                gr.add_edge(*ed, **{_elk: [self.brain[name]]})
                return
            exel.append(self.brain[name])
            return

        if False:
            return self.last_gr
        else:
            gr = nx.Graph()
            ns = self.brain.Circuit.AllElementNames()
            for name in ns:
                try:
                    buses = self.brain[name].BusNames()
                except TypeError:
                    continue

                #todo encode term perms in the graph
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

            self.up_to_date = True
            self.last_gr = gr
        return gr

    @property
    def bus_coords(self):
        bp = {}
        for bn in self.brain.Circuit.AllBusNames():
            if self.brain['bus.' + bn].Coorddefined():
                bp[bn] = (self.brain[bn].X(), self.brain[bn].Y())
            else:
                bp[bn] = None
        return bp

    @staticmethod
    def _bus_resolve(bus_descriptor: str):

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
    """Simple extension of nx.Digraph created to reverse-walk a dependency tree in order to declare the entities in the
    right order."""
    @property
    def leaves(self):
        return [x for x in self.nodes() if self.out_degree(x) == 0]

    def trim(self):
        self.remove_nodes_from(self.leaves)


class _BusView:
    """_BusView is meant to be instantiated only by Krang and is returned by indicization like Krang['bus1', 'bus2'] or
    Krang['bus1']. The main uses of a _BusView are:

        -Adding elements to the corresponding bus/edge:
            Krang['bus1'] << kp.Load()  # adds a Load to bus1
            Krang['bus1', 'bus2'] << kp.Line()  # adds a Line between bus1 and bus2

        -Evaluating a function from the submodule 'busquery' on the corresponding bus/edge through attribute resolution:
            Krang['bus1'].voltage
            Krang['bus1'].totload

        -Getting a _PackedOpendssElement from the ones pertaining the corresponding bus/edge:
            Krang['bus1']['myload']
    """

    def __init__(self, oek: Krang, bustermtuples):
        self.btt = bustermtuples
        self.tp = dict(bustermtuples)
        self.buses = tuple(self.tp.keys())
        self.nb = len(self.buses)
        self.oek = oek

        buskey = 'buses'
        tkey = 'terminals'

        if len(self.buses) == 1:
            try:
                self.content = self.oek.graph.nodes[self.buses[0]][_elk]
                # it's ok that a KeyError by both indexes is caught in the same way
            except KeyError:
                self.content = []
        elif len(self.buses) == 2:
            self.content = list(nx.get_edge_attributes(self.oek.graph.subgraph(self.buses), _elk).values())
        else:
            raise ValueError

        self.fcs_kwargs = {buskey: self.buses, tkey: self.tp}

    def __lshift__(self, other):
        assert not other.isnamed()
        try:
            assert other.name != ''
        except AssertionError:
            raise ValueError('Did you name the element before adding it?')
        self.oek.command(other.fcs(**self.fcs_kwargs))

        # remember ai elements
        if other.isai():
            self.oek.ai_list.append(other)

        return _BusView(self.oek, self.btt)

    def __getitem__(self, item):
        return self.content[item]

    def __getattr__(self, item):
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
    l_ckt = Krang()
    l_ckt.initialize(master_dict['cktname'], au.dejsonize(master_dict['elements']['vsource.source']))
    del master_dict['elements']['vsource.source']

    # load and declare options
    opt_dict = master_dict['settings']
    for on, ov in opt_dict['values'].items():
        # we try in any way to see if the value is the same as the default and, if so, we continue
        # todo there is an edge case where the value of a measured quantity is the same, but the unit is different
        if ov == co.default_settings['values'][on]:
            continue
        try:
            if np.isclose(ov, co.default_settings['values'][on]):
                continue
        except ValueError:
            if np.isclose(ov, co.default_settings['values'][on]).all():
                continue
        except TypeError:
            pass

        try:
            if ov.lower() == co.default_settings['values'][on].lower():
                continue
        except AttributeError:
            pass

        if on in opt_dict['units'].keys():
            d_ov = ov * um.parse_units(opt_dict['units'][on])
        else:
            d_ov = ov

        l_ckt.set(**{on: d_ov})

    # reconstruction of dependency graph and declarations
    dep_graph = _DepGraph()
    for jobj in master_dict['elements'].values():
        # if the element has no dependencies, we just add a node with iths name
        if jobj['depends'] == {} or all([d == '' for d in jobj['depends'].values()]):
            dep_graph.add_node(jobj['type'] + '.' + jobj['name'])
        else:
            # if an element parameter depends on another name, or a list of other names, we create all the edges
            # necessary
            for dvalue in jobj['depends'].values():
                if isinstance(dvalue, list):
                    for dv in dvalue:
                        if dv != '':
                            dep_graph.add_edge(jobj['type'] + '.' +jobj['name'], dv)
                else:
                    if dvalue != '':
                        dep_graph.add_edge(jobj['type'] + '.' + jobj['name'], dvalue)

    # we cyclically consider all "leaves", add the objects at the leaves, then trim the leaves and go on with
    # the new leaves.
    # In this way we are sure that, whenever a name is mentioned in a fcs, its entity was already declared.
    while dep_graph.leaves:
        for nm in dep_graph.leaves:
            try:
                jobj = copy.deepcopy(master_dict['elements'][nm])
            except KeyError:
                mdmod = {k.split('.')[1]: v for k, v in master_dict['elements'].items()}
                jobj = copy.deepcopy(mdmod[nm])
            dssobj = au.dejsonize(jobj)
            if dssobj.isnamed():
                l_ckt << dssobj
            elif dssobj.isabove():
                l_ckt << dssobj.aka(jobj['name'])
            else:
                l_ckt[tuple(jobj['topological'])] << dssobj.aka(jobj['name'])
                # l_ckt.command(dssobj.aka(jobj['name']).fcs(buses=jobj['topological']))
        dep_graph.trim()

    return l_ckt


def _main():
    pass


if __name__ == '__main__':
    _main()
