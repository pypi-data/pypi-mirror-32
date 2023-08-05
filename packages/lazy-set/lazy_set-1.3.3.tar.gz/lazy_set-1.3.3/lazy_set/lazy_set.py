from itertools import chain
from typing import Set, Iterable

from collections import deque
from collections.abc import MutableSet


class LazySet(MutableSet):
    """
    A collection that tries to imitate a "lazy" difference and union of sets.
    """

    def __init__(self, base_set: Set = frozenset(),
                 negative_sets: Iterable[Set] = list([frozenset()]),
                 positive_sets: Iterable[Set] = list([frozenset()])):
        """
        Initializes the LazySet.
        :param base_set: the base set from which all the negative sets are "removed" and to which
        all positive sets are "added". Note that the order is important! The resulting set is identical to:
        base_set.difference(*negative_sets).union(*positive_sets)
        :param negative_sets: sets that contain items that are to be "removed" from the *base* set.
        :param positive_sets: sets that contain items that are to be "added" to the base set
        *after* the "removal" of the negative items.
        """
        self._sets = []
        self._negative_indices = deque()
        self._positive_indices = deque()

        self.lazy_update(base_set)
        for negative_set in negative_sets:
            self.lazy_difference_update(negative_set)
        for positive_set in positive_sets:
            self.lazy_update(positive_set)

    def lazy_update(self, other) -> "LazySet":
        """
        Update the LazySet, adding elements from other.
        """
        self._positive_indices.appendleft(len(self._sets))
        self._sets.append(other)
        return self

    def update(self, *others):
        """
        Update the LazySet, adding elements from all others.
        """
        return self.lazy_update(set(*others))

    def __ior__(self, *others):
        """
        Update the LazySet, adding elements from all others.
        """
        return self.update(*others)

    def union(self, *others):
        """
        :return: a new LazySet with elements from the set and all others.
        """
        return self.copy().update(*others)

    def __or__(self, *others):
        """
        :return: a new LazySet with elements from the set and all others.
        """
        return self.union(*others)

    def lazy_difference_update(self, other) -> "LazySet":
        """
        Update the LazySet, removing elements found in other.
        """
        self._negative_indices.appendleft(len(self._sets))
        self._sets.append(other)
        return self

    def difference_update(self, *others):
        """
        Update the LazySet, removing elements found in others.
        """
        return self.lazy_difference_update(set(*others))

    def __isub__(self, *others):
        """
        Update the LazySet, removing elements found in others.
        """
        return self.difference_update(*others)

    def difference(self, *others):
        """
        :return: a new LazySet with elements in the set that are not in the others.
        """
        return self.copy().difference_update(*others)

    def __sub__(self, *others):
        """
        :return: a new LazySet with elements in the set that are not in the others.
        """
        return self.difference(*others)

    def _intersection(self, as_set, *others):
        """
        :return: a new LazySet/set (according to as_set) with elements common to the set and all others.
        """
        intersection_set = self.copy_to_set()
        intersection_set.intersection_update(*others)
        if as_set:
            return intersection_set
        return LazySet(base_set=intersection_set)

    def intersection(self, *others):
        """
        :return: a new LazySet with elements common to the set and all others.
        """
        return self._intersection(False, *others)

    def __and__(self, *others):
        """
        :return: a new LazySet with elements common to the set and all others.
        """
        return self.intersection(*others)

    def intersection_update(self, *others):
        """
        Update the LazySet, keeping only elements found in it and all others.
        """
        intersection_set = self._intersection(True, *others)
        self.clear()
        return self.lazy_update(intersection_set)

    def __iand__(self, *others):
        """
        Update the LazySet, keeping only elements found in it and all others.
        """
        return self.intersection_update(*others)

    def symmetric_difference_update(self, other):
        """
        Update the LazySet, keeping only elements found in either set, but not in both.
        """
        elements_in_common = set()
        for elem in chain(self, other):
            if elem in other and elem in self:
                elements_in_common.add(elem)
        self.lazy_update(other)
        self.lazy_difference_update(elements_in_common)
        return self

    def __ixor__(self, other):
        """
        Update the LazySet, keeping only elements found in either set, but not in both.
        """
        return self.symmetric_difference_update(other)

    def symmetric_difference(self, other):
        """
        :return: a new LazySet with elements in either the set or other but not both.
        """
        return self.copy().symmetric_difference_update(other)

    def __xor__(self, other):
        """
        :return: a new LazySet with elements in either the set or other but not both.
        """
        return self.symmetric_difference(other)

    def issubset(self, other):
        """
        :return: True iff every element in the LazySet is in other.
        """
        for item in self:
            if item not in other:
                return False
        return True

    def __le__(self, other):
        """
        :return: True iff every element in the LazySet is in other.
        """
        return self.issubset(other)

    def __lt__(self, other):
        """
        :return: True iff every element in the LazySet is in other and self != other.
        """
        for item in other:
            if item not in self:
                # there is one item missing in self, so definitely self != other
                return self <= other
        return False

    def issuperset(self, other):
        """
        :return: True iff every element in other is in the LazySet.
        """
        return other.issubset(self)

    def __ge__(self, other):
        """
        :return: True iff every element in other is in the LazySet.
        """
        return self.issuperset(other)

    def __gt__(self, other):
        """
        :return: True iff every element in other is in the LazySet and self != other.
        """
        for item in self:
            if item not in other:
                # there is one item missing in other, so definitely self != other
                return other <= self
        return False

    def __eq__(self, other):
        """
        :return: True iff both sets contain exactly the same elements.
        """
        return self <= other <= self

    def __ne__(self, other):
        """
        :return: True iff one set contains at least one element that the other doesn't.
        """
        return not self == other

    def _in_negative(self, item, positive_index=-1):
        """
        :return: True iff the given item is in one of the negative sets.
        """
        for negative_index in self._negative_indices:
            if negative_index < positive_index:
                break
            if item in self._sets[negative_index]:
                return True
        return False

    def __contains__(self, item):
        for positive_index in self._positive_indices:
            if item in self._sets[positive_index]:
                return not self._in_negative(item, positive_index)
        return False

    def __iter__(self):
        # Note: can iterate on same item a number of times! So need to keep track on already yielded items.
        # Note2: doesn't support modification while iterating!
        already_yielded = set()
        for positive_index in self._positive_indices:
            for item in self._sets[positive_index]:
                if item not in already_yielded and not self._in_negative(item, positive_index):
                    already_yielded.add(item)
                    yield item
        already_yielded.clear()

    def __len__(self):
        """
        :return: the number of items in the LazySet.
        """
        # make note - this is SLOW!
        count = 0
        for item in self:
            count += 1
        return count

    def add(self, elem) -> "LazySet":
        """
        Add element elem to the LazySet.
        """
        return self.update({elem})

    def remove(self, elem) -> "LazySet":
        """
        Remove element elem from the LazySet. Raises KeyError if elem is not contained in the LazySet.
        """
        if elem not in self:
            raise KeyError(elem)
        return self.discard(elem)

    def discard(self, elem) -> "LazySet":
        """
        Remove element elem from the LazySet if it is present.
        """
        return self.difference_update({elem})

    def clear(self) -> "LazySet":
        """ Remove all elements from the LazySet. """
        self._sets.clear()
        self._positive_indices.clear()
        self._negative_indices.clear()
        return self

    def copy(self) -> "LazySet":
        """
        Return a shallow copy of this LazySet.
        """
        shallow_copy = LazySet()
        shallow_copy._sets = self._sets.copy()
        shallow_copy._positive_indices = self._positive_indices.copy()
        shallow_copy._negative_indices = self._negative_indices.copy()
        return shallow_copy

    def copy_to_set(self) -> Set:
        """
        Shallow copies this LazySet to a regular set.
        :rtype: set
        :return: a shallow copy of this LazySet as a set.
        """
        return set(self)
