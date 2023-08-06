# This file is a part of physics
#
# Copyright (c) 2018 The physics Authors (see AUTHORS)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice
# shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

"""
physics.proportionality
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
It contains the Proportionality class.

It could be used to define a proportionality relation
between numbers.
"""

cdef class Proportionality:

    """
    The proportionality class is used to
    calculate and use proportionality using
    numbers.
    percentage error and do arithmetic
    """

    cdef readonly float constant
    cdef readonly str relation
    cdef readonly str formula
    constant_formulas = {'direct': lambda x, y: y/x, 'inverse': lambda x, y: x*y, 'square': lambda x, y: y/x**2, 'inverse_square': lambda x, y: y*x**2}

    def __init__(self, **options):
        r"""
        It initializes the object and
        it checks **options parameter,
        and then get the constant of the
        proportionality.

        :param \**options: The constant and relation or a dict of numbers (x is the key, y the value)
        :type \**options: dict
        :raises MissingNeededParameters: It throws an exception if some parameters are missing.
        :raises NoRelationError: It throws an exception if there are no relations between numbers.
        :raises LessThanTwoNumbersError: It throws an exception if there are less numbers than 2.
        """

        cdef dict relations = {
            'direct': 'k*x',
            'inverse': 'k/x',
            'square': 'k*(x**2)',
            'inverse_square': 'k/(x**2)'}
        cdef int numbers_options

        if 'constant' in options and 'relation' in options:
            self.constant = float(options['constant'])
            if options['relation'] in relations:
                self.relation = options['relation'][0].title()
                self.formula = relations.get(options['relation'])
                return
            else:
                raise NoRelationError()
        elif 'numbers' in options:
            if 0 in options['numbers']:
                numbers_options = len(options['numbers']) - 1
            else:
                numbers_options = len(options['numbers'])
            if numbers_options > 1:
                self.relation, self.constant = self.search_proportionality(options['numbers'])
                self.formula = relations.get(self.relation)
                return
            else:
                raise LessThanTwoNumbersError()
        else:
            raise MissingNeededParameters()

    cdef tuple search_proportionality(self, dict numbers):
        cdef list propor = self.constant_formulas.keys()
        cdef str pr
        cdef int constant
        cdef bint right

        for pr in propor:
            right, constant = self.is_proportionality(pr, numbers)
            if right:
                return pr, constant

        raise NoRelationError()

    cdef tuple is_proportionality(self, str proportionality, dict numbers):
        cdef int constant = 0
        cdef int last_constant = 0
        cdef bint right
        cdef:
            int x
            int y

        for x, y in numbers.items():
            if x is 0 or y is 0:
                continue

            constant = self.constant_formulas[proportionality](x, y)

            if last_constant == constant:
                right = True
            elif last_constant == 0:
                right = True
            else:
                right = False
                break
            last_constant = constant

        return right, constant

    cpdef float calculate(self, float x):
        r"""
        Calculate the y using
        the formula created
        during proportionality
        check.

        :param x: The number you want to calculate.
        :type x: float
        """

        cdef float k = self.constant
        return eval(self.formula)

    def __str__(self) -> str:
        """
        Return the relation
        and the constant.

        :returns: The relation and its constant.
        :rtype: str
        """
        return ("Relation: " + self.relation +
                "\nConstant: " + self.constant)


cdef class LessThanTwoNumbersError(Exception):
    """
    This exception is called when
    number of parameters are less
    than 2. 0 is not counted.
    """

    def __init__(self):
        Exception.__init__(self, "Numbers parameters are less than 2")


cdef class NoRelationError(Exception):
    """
    This exception is called when
    there's no relation.
    """

    def __init__(self):
        Exception.__init__(self, "There's no relation")


cdef class MissingNeededParameters(Exception):
    """
    This exception is called when
    constant and proportionality aren't
    in the parameters and numbers is missing.
    """

    def __init__(self):
        Exception.__init__(self, "There's no relation")
