# Copyright (c) 2017-2018 Fumito Hamamura <fumito.ham@gmail.com>

# This library is free software: you can redistribute it and/or
# modify it under the terms of the GNU Lesser General Public
# License as published by the Free Software Foundation version 3.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# Lesser General Public License for more details.
#
# You should have received a copy of the GNU Lesser General Public
# License along with this library.  If not, see <http://www.gnu.org/licenses/>.

import builtins
from types import MappingProxyType
from collections import OrderedDict
from textwrap import dedent
import pickle

import networkx as nx

from modelx.core.base import (
    get_interfaces,
    ImplDict,
    ImplChainMap,
    BaseMapProxy)
from modelx.core.cells import CellArgs
from modelx.core.space import (
    SpaceContainerImpl,
    SpaceContainer,
    SpaceMapProxy,
    RefDict)
from modelx.core.util import is_valid_name


class DependencyGraph(nx.DiGraph):
    """Directed Graph of ObjectArgs"""

    def clear_descendants(self, source, clear_source=True):
        """Remove all descendants of(reachable from) `source`.

        Args:
            source: Node descendants
            clear_source(bool): Remove origin too if True.
        Returns:
            set: The removed nodes.
        """
        desc = nx.descendants(self, source)
        if clear_source:
            desc.add(source)
        self.remove_nodes_from(desc)
        return desc

    def clear_obj(self, obj):
        """"Remove all nodes with `obj` and their descendants."""
        obj_nodes = self.get_nodes_with(obj)
        removed = set()
        for node in obj_nodes:
            if self.has_node(node):
                removed.update(self.clear_descendants(node))
        return removed

    def get_nodes_with(self, obj):
        """Return nodes with `obj`."""
        result = set()

        if nx.__version__[0] == '1':
            nodes = self.nodes_iter()
        else:
            nodes = self.nodes

        for node in nodes:
            if node.obj_ == obj:
                result.add(node)
        return result

    def fresh_copy(self):
        """Overriding Graph.fresh_copy"""
        return DependencyGraph()


class ModelImpl(SpaceContainerImpl):

    def __init__(self, *, system, name):
        SpaceContainerImpl.__init__(self, system, if_class=Model)

        self.cellgraph = DependencyGraph()
        self.currentspace = None

        if not name:
            self.name = system._modelnamer.get_next(system.models)
        elif is_valid_name(name):
            self.name = name
        else:
            raise ValueError("Invalid name '%s'." % name)

        data = {'__builtins__': builtins}
        self._global_refs = RefDict(self, data=data)
        self._spaces = ImplDict(self, SpaceMapProxy)
        self._namespace = ImplChainMap(
            self, BaseMapProxy, [self._spaces, self._global_refs])
        self.allow_none = False
        self.lazy_evals = self._namespace

    def rename(self, name):
        """Rename self. Must be called only by its system."""
        if is_valid_name(name):
            if name not in self.system.models:
                self.name = name
                return True     # Rename success
            else:               # Model name already exists
                return False
        else:
            raise ValueError("Invalid name '%s'." % name)

    def clear_descendants(self, source, clear_source=True):
        """Clear values and nodes calculated from `source`."""
        removed = self.cellgraph.clear_descendants(source, clear_source)
        for cell in removed:
            del cell.cells.data[cell.argvalues]

    def clear_obj(self, obj):
        """Clear values and nodes of `obj` and their dependants."""
        removed = self.cellgraph.clear_obj(obj)
        for cell in removed:
            del cell.cells.data[cell.argvalues]

    @property
    def repr_(self):
        format_ = dedent("""\
        name: %s
        spaces(%s): %s""")

        return format_ % (
            self.name,
            len(self.spaces), list(self.spaces.keys()))

    def __repr__(self):
        return self.name

    @property
    def _repr_self(self):
        return self.name

    @property
    def _repr_parent(self):
        return ''

    @property
    def fullname(self):
        return self.name

    @property
    def model(self):
        return self

    @property
    def global_refs(self):
        return self._global_refs.get_updated()

    @property
    def namespace(self):
        return self._namespace.get_updated()

    def close(self):
        self.system.close_model(self)

    def save(self, filepath):
        self.update_lazyevals()
        with open(filepath, 'wb') as file:
            pickle.dump(self.interface, file, protocol=4)

    def get_object(self, name):
        """Retrieve an object by a dotted name relative to the model."""
        parts = name.split('.')
        space = self.spaces[parts.pop(0)]
        if parts:
            return space.get_object('.'.join(parts))
        else:
            return space

    # ----------------------------------------------------------------------
    # Serialization by pickle

    state_attrs = ['name',
                   'cellgraph',
                   '_namespace',
                   '_global_refs'] + SpaceContainerImpl.state_attrs

    def __getstate__(self):

        state = {key: value for key, value in self.__dict__.items()
                 if key in self.state_attrs}

        mapping = {}
        for node in self.cellgraph:
            name = node.obj_.get_fullname(omit_model=True)
            mapping[node] = (name, node.argvalues)

        state['cellgraph'] = nx.relabel_nodes(self.cellgraph, mapping)

        return state

    def __setstate__(self, state):
        self.__dict__.update(state)

    def restore_state(self, system):
        """Called after unpickling to restore some attributes manually."""
        SpaceContainerImpl.restore_state(self, system)
        mapping = {}
        for node in self.cellgraph:
            name, argvalues = node
            cells = self.get_object(name)
            mapping[node] = CellArgs(cells, argvalues)

        self.cellgraph = nx.relabel_nodes(self.cellgraph, mapping)

    def del_space(self, name):
        space = self.spaces[name]
        if space.has_inherited():
            raise ValueError("%s has derived spaces" % repr(space.interface))
        else:
            self.spaces.del_item(name)

    def _set_space(self, space):

        if space.name in self.spaces:
            self.del_space(space.name)
        elif space.name in self.global_refs:
            raise KeyError("Name '%s' already already assigned" % self.name)

        self.spaces.set_item(space.name, space)

    def del_ref(self, name):
        self.global_refs.del_item(name)

    def get_attr(self, name):
        if name in self.spaces:
            return self.spaces[name].interface
        elif name in self.global_refs:
            return get_interfaces(self.global_refs[name])
        else:
            raise KeyError("Model '{0}' does not have '{1}'".\
                           format(self.name, name))

    def set_attr(self, name, value):
        if name in self.spaces:
            raise KeyError("Space named '%s' already exist" % self.name)

        self.global_refs.set_item(name, value)

    def del_attr(self, name):

        if name in self.spaces:
            self.del_space(name)
        elif name in self.global_refs:
            self.del_ref(name)
        else:
            raise KeyError("Name '%s' not defined" % name)


class Model(SpaceContainer):
    """Top-level container in modelx object hierarchy.

    Model instances are the top-level objects and directly contain
    :py:class:`Space <modelx.core.space.Space>` objects, which in turn
    contain other spaces or
    :py:class:`Cells <modelx.core.cells.Cells>` objects.

    A model can be created by
    :py:func:`new_model <modelx.core.model.Model>` API function.
    """

    def rename(self, name):
        """Rename the model itself"""
        self._impl.system.rename_model(new_name=name, old_name=self.name)

    def save(self, filepath):
        """Save the model to a file."""
        self._impl.save(filepath)

    def close(self):
        """Close the model."""
        self._impl.close()

    # ----------------------------------------------------------------------
    # Getting and setting attributes

    def __getattr__(self, name):
        return self._impl.get_attr(name)

    def __setattr__(self, name, value):
        if name in self.properties:
            object.__setattr__(self, name, value)
        else:
            self._impl.set_attr(name, value)

    def __delattr__(self, name):
        self._impl.del_attr(name)

    def __dir__(self):
        return self._impl.namespace.interfaces

    @property
    def cellgraph(self):
        """A directed graph of cells."""
        return self._impl.cellgraph

    @property
    def refs(self):
        """Return a mapping of global references."""
        return self._impl.global_refs.mproxy




