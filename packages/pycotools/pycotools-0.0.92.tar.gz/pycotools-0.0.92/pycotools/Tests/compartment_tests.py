# -*-coding: utf-8 -*-
"""

 This file is part of pycotools.

 pycotools is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 pycotools is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License for more details.

 You should have received a copy of the GNU Lesser General Public License
 along with pycotools.  If not, see <http://www.gnu.org/licenses/>.


 $Author: Ciaran Welsh


"""

import site

# site.addsitedir('C:\Users\Ciaran\Documents\pycotools')
site.addsitedir('/home/b3053674/Documents/pycotools')

import pycotools

from pycotools.model import Compartment
from pycotools.Tests import _test_base
import os, glob
import pandas
import unittest


class TestCompartments(_test_base._BaseTest):
    def setUp(self):
        super(TestCompartments, self).setUp()
        self.comp = Compartment(self.model, name='Nucleus', initial_value=5,
                                simulation_type='fixed', key='compartment_1')

    def test_checks(self):
        """
        Doesnot work at present.
        :return:
        """
        pass
        # with self.assertRaises(Exception):
        #     pycotools.pycopi.Compartments()

    def test_comp_name(self):
        self.assertEqual(self.comp.name, 'Nucleus')

    def test_comp_size(self):
        self.assertEqual(self.comp.initial_value, 5)

    def test_compartment_str(self):
        ## TODO modify this test for updated compartment string
        pass
        # comp_str = "Compartment(key='compartment_1', name='Nucleus', type='fixed', value=5)"
        # self.assertEqual(comp_str, self.comp.__str__())

    def test_reference(self):
        ref = 'Vector=Compartments[Nucleus]'
        self.assertTrue(self.comp.reference == ref)


if __name__ == '__main__':
    unittest.main()
