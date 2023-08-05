#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Utility datastructures."""


class EList(list):
    """
    Enhanced List.

    This class supports every operation a normal list supports. Additionally,
    you can call it with a list as an argument.

    Examples
    --------
    >>> l = EList([2, 1, 0])
    >>> l[2]
    0
    >>> l[[2, 0]]
    [0, 2]
    >>> l[l]
    [0, 1, 2]
    """

    def __init__(self, *args):
        list.__init__(self, *args)

    def __getitem__(self, key):
        if isinstance(key, list):
            return EList([self[index] for index in key])
        else:
            return list.__getitem__(self, key)

    def remove_indices(self, indices):
        """
        Remove rows by which have the given indices.

        Parameters
        ----------
        indices : list

        Returns
        -------
        filtered_list : EList
        """
        new_list = []
        for index, el in enumerate(self):
            if index not in indices:
                new_list.append(el)
        return EList(new_list)
