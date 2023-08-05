# ext/orderinglist.py
# Copyright (C) 2005-2018 the SQLAlchemy authors and contributors
# <see AUTHORS file>
#
# This module is part of SQLAlchemy and is released under
# the MIT License: http://www.opensource.org/licenses/mit-license.php

"""A custom list that manages index/position information for contained
elements.

:author: Jason Kirtland

``orderinglist`` is a helper for mutable ordered relationships.  It will
intercept list operations performed on a :func:`.relationship`-managed
collection and
automatically synchronize changes in list position onto a target scalar
attribute.

Example: A ``slide`` table, where each row refers to zero or more entries
in a related ``bullet`` table.   The bullets within a slide are
displayed in order based on the value of the ``position`` column in the
``bullet`` table.   As entries are reordered in memory, the value of the
``position`` attribute should be updated to reflect the new sort order::


    Base = declarative_base()

    class Slide(Base):
        __tablename__ = 'slide'

        id = Column(Integer, primary_key=True)
        name = Column(String)

        bullets = relationship("Bullet", order_by="Bullet.position")

    class Bullet(Base):
        __tablename__ = 'bullet'
        id = Column(Integer, primary_key=True)
        slide_id = Column(Integer, ForeignKey('slide.id'))
        position = Column(Integer)
        text = Column(String)

The standard relationship mapping will produce a list-like attribute on each
``Slide`` containing all related ``Bullet`` objects,
but coping with changes in ordering is not handled automatically.
When appending a ``Bullet`` into ``Slide.bullets``, the ``Bullet.position``
attribute will remain unset until manually assigned.   When the ``Bullet``
is inserted into the middle of the list, the following ``Bullet`` objects
will also need to be renumbered.

The :class:`.OrderingList` object automates this task, managing the
``position`` attribute on all ``Bullet`` objects in the collection.  It is
constructed using the :func:`.ordering_list` factory::

    from sqlalchemy.ext.orderinglist import ordering_list

    Base = declarative_base()

    class Slide(Base):
        __tablename__ = 'slide'

        id = Column(Integer, primary_key=True)
        name = Column(String)

        bullets = relationship("Bullet", order_by="Bullet.position",
                                collection_class=ordering_list('position'))

    class Bullet(Base):
        __tablename__ = 'bullet'
        id = Column(Integer, primary_key=True)
        slide_id = Column(Integer, ForeignKey('slide.id'))
        position = Column(Integer)
        text = Column(String)

With the above mapping the ``Bullet.position`` attribute is managed::

    s = Slide()
    s.bullets.append(Bullet())
    s.bullets.append(Bullet())
    s.bullets[1].position
    >>> 1
    s.bullets.insert(1, Bullet())
    s.bullets[2].position
    >>> 2

The :class:`.OrderingList` construct only works with **changes** to a
collection, and not the initial load from the database, and requires that the
list be sorted when loaded.  Therefore, be sure to specify ``order_by`` on the
:func:`.relationship` against the target ordering attribute, so that the
ordering is correct when first loaded.

.. warning::

  :class:`.OrderingList` only provides limited functionality when a primary
  key column or unique column is the target of the sort.  Operations
  that are unsupported or are problematic include:

    * two entries must trade values.  This is not supported directly in the
      case of a primary key or unique constraint because it means at least
      one row would need to be temporarily removed first, or changed to
      a third, neutral value while the switch occurs.

    * an entry must be deleted in order to make room for a new entry.
      SQLAlchemy's unit of work performs all INSERTs before DELETEs within a
      single flush.  In the case of a primary key, it will trade
      an INSERT/DELETE of the same primary key for an UPDATE statement in order
      to lessen the impact of this limitation, however this does not take place
      for a UNIQUE column.
      A future feature will allow the "DELETE before INSERT" behavior to be
      possible, allevating this limitation, though this feature will require
      explicit configuration at the mapper level for sets of columns that
      are to be handled in this way.

:func:`.ordering_list` takes the name of the related object's ordering
attribute as an argument.  By default, the zero-based integer index of the
object's position in the :func:`.ordering_list` is synchronized with the
ordering attribute: index 0 will get position 0, index 1 position 1, etc.  To
start numbering at 1 or some other integer, provide ``count_from=1``.


"""
from ..orm.collections import collection, collection_adapter
from .. import util

__all__ = ['ordering_list']


def ordering_list(attr, count_from=None, **kw):
    """Prepares an :class:`OrderingList` factory for use in mapper definitions.

    Returns an object suitable for use as an argument to a Mapper
    relationship's ``collection_class`` option.  e.g.::

        from sqlalchemy.ext.orderinglist import ordering_list

        class Slide(Base):
            __tablename__ = 'slide'

            id = Column(Integer, primary_key=True)
            name = Column(String)

            bullets = relationship("Bullet", order_by="Bullet.position",
                                    collection_class=ordering_list('position'))

    :param attr:
      Name of the mapped attribute to use for storage and retrieval of
      ordering information

    :param count_from:
      Set up an integer-based ordering, starting at ``count_from``.  For
      example, ``ordering_list('pos', count_from=1)`` would create a 1-based
      list in SQL, storing the value in the 'pos' column.  Ignored if
      ``ordering_func`` is supplied.

    Additional arguments are passed to the :class:`.OrderingList` constructor.

    """

    kw = _unsugar_count_from(count_from=count_from, **kw)
    return lambda: OrderingList(attr, **kw)


# Ordering utility functions


def count_from_0(index, collection):
    """Numbering function: consecutive integers starting at 0."""

    return index


def count_from_1(index, collection):
    """Numbering function: consecutive integers starting at 1."""

    return index + 1


def count_from_n_factory(start):
    """Numbering function: consecutive integers starting at arbitrary start."""

    def f(index, collection):
        return index + start
    try:
        f.__name__ = 'count_from_%i' % start
    except TypeError:
        pass
    return f


def _unsugar_count_from(**kw):
    """Builds counting functions from keyword arguments.

    Keyword argument filter, prepares a simple ``ordering_func`` from a
    ``count_from`` argument, otherwise passes ``ordering_func`` on unchanged.
    """

    count_from = kw.pop('count_from', None)
    if kw.get('ordering_func', None) is None and count_from is not None:
        if count_from == 0:
            kw['ordering_func'] = count_from_0
        elif count_from == 1:
            kw['ordering_func'] = count_from_1
        else:
            kw['ordering_func'] = count_from_n_factory(count_from)
    return kw


class OrderingList(list):
    """A custom list that manages position information for its children.

    The :class:`.OrderingList` object is normally set up using the
    :func:`.ordering_list` factory function, used in conjunction with
    the :func:`.relationship` function.

    """

    def __init__(self, ordering_attr=None, ordering_func=None,
                 reorder_on_append=False):
        """A custom list that manages position information for its children.

        ``OrderingList`` is a ``collection_class`` list implementation that
        syncs position in a Python list with a position attribute on the
        mapped objects.

        This implementation relies on the list starting in the proper order,
        so be **sure** to put an ``order_by`` on your relationship.

        :param ordering_attr:
          Name of the attribute that stores the object's order in the
          relationship.

        :param ordering_func: Optional.  A function that maps the position in
          the Python list to a value to store in the
          ``ordering_attr``.  Values returned are usually (but need not be!)
          integers.

          An ``ordering_func`` is called with two positional parameters: the
          index of the element in the list, and the list itself.

          If omitted, Python list indexes are used for the attribute values.
          Two basic pre-built numbering functions are provided in this module:
          ``count_from_0`` and ``count_from_1``.  For more exotic examples
          like stepped numbering, alphabetical and Fibonacci numbering, see
          the unit tests.

        :param reorder_on_append:
          Default False.  When appending an object with an existing (non-None)
          ordering value, that value will be left untouched unless
          ``reorder_on_append`` is true.  This is an optimization to avoid a
          variety of dangerous unexpected database writes.

          SQLAlchemy will add instances to the list via append() when your
          object loads.  If for some reason the result set from the database
          skips a step in the ordering (say, row '1' is missing but you get
          '2', '3', and '4'), reorder_on_append=True would immediately
          renumber the items to '1', '2', '3'.  If you have multiple sessions
          making changes, any of whom happen to load this collection even in
          passing, all of the sessions would try to "clean up" the numbering
          in their commits, possibly causing all but one to fail with a
          concurrent modification error.

          Recommend leaving this with the default of False, and just call
          ``reorder()`` if you're doing ``append()`` operations with
          previously ordered instances or when doing some housekeeping after
          manual sql operations.

        """
        self.ordering_attr = ordering_attr
        if ordering_func is None:
            ordering_func = count_from_0
        self.ordering_func = ordering_func
        self.reorder_on_append = reorder_on_append

    # More complex serialization schemes (multi column, e.g.) are possible by
    # subclassing and reimplementing these two methods.
    def _get_order_value(self, entity):
        return getattr(entity, self.ordering_attr)

    def _set_order_value(self, entity, value):
        setattr(entity, self.ordering_attr, value)

    def reorder(self):
        """Synchronize ordering for the entire collection.

        Sweeps through the list and ensures that each object has accurate
        ordering information set.

        """
        for index, entity in enumerate(self):
            self._order_entity(index, entity, True)

    # As of 0.5, _reorder is no longer semi-private
    _reorder = reorder

    def _order_entity(self, index, entity, reorder=True):
        have = self._get_order_value(entity)

        # Don't disturb existing ordering if reorder is False
        if have is not None and not reorder:
            return

        should_be = self.ordering_func(index, self)
        if have != should_be:
            self._set_order_value(entity, should_be)

    def append(self, entity):
        super(OrderingList, self).append(entity)
        self._order_entity(len(self) - 1, entity, self.reorder_on_append)

    def _raw_append(self, entity):
        """Append without any ordering behavior."""

        super(OrderingList, self).append(entity)
    _raw_append = collection.adds(1)(_raw_append)

    def insert(self, index, entity):
        super(OrderingList, self).insert(index, entity)
        self._reorder()

    def remove(self, entity):
        super(OrderingList, self).remove(entity)

        adapter = collection_adapter(self)
        if adapter and adapter._referenced_by_owner:
            self._reorder()

    def pop(self, index=-1):
        entity = super(OrderingList, self).pop(index)
        self._reorder()
        return entity

    def __setitem__(self, index, entity):
        if isinstance(index, slice):
            step = index.step or 1
            start = index.start or 0
            if start < 0:
                start += len(self)
            stop = index.stop or len(self)
            if stop < 0:
                stop += len(self)

            for i in range(start, stop, step):
                self.__setitem__(i, entity[i])
        else:
            self._order_entity(index, entity, True)
            super(OrderingList, self).__setitem__(index, entity)

    def __delitem__(self, index):
        super(OrderingList, self).__delitem__(index)
        self._reorder()

    def __setslice__(self, start, end, values):
        super(OrderingList, self).__setslice__(start, end, values)
        self._reorder()

    def __delslice__(self, start, end):
        super(OrderingList, self).__delslice__(start, end)
        self._reorder()

    def __reduce__(self):
        return _reconstitute, (self.__class__, self.__dict__, list(self))

    for func_name, func in list(locals().items()):
        if (util.callable(func) and func.__name__ == func_name and
                not func.__doc__ and hasattr(list, func_name)):
            func.__doc__ = getattr(list, func_name).__doc__
    del func_name, func


def _reconstitute(cls, dict_, items):
    """ Reconstitute an :class:`.OrderingList`.

    This is the adjoint to :meth:`.OrderingList.__reduce__`.  It is used for
    unpickling :class:`.OrderingList` objects.

    """
    obj = cls.__new__(cls)
    obj.__dict__.update(dict_)
    list.extend(obj, items)
    return obj
