#-*-coding: utf-8 -*-
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

Module that tests the operations of the _Base base test

"""

import site
site.addsitedir('C:\Users\Ciaran\Documents\pycotools')
site.addsitedir(r'/home/b3053674/Documents/pycotools')
import pycotools
from pycotools.Tests import _test_base
import os, glob
import pandas
import unittest
from lxml import etree




class TestBase(_test_base._BaseTest):
    def setUp(self):
        super(TestBase, self).setUp()
        self._base = pycotools._base._Base(A='a', B='b', key='MadeUpKey1')

    def test_num_kwargs(self):
        ## 4 kwargs defined. A and B in setUp
        ## and the kwargs dict
        self.assertEqual(len(self._base.__dict__), 4)

    def test_set_random_attribute(self):
        self.assertEqual(self._base.A, 'a')

    def test_set_random_attribute2(self):
        self.assertEqual(self._base.B, 'b')

    def test_setattr(self):
        self._base.A = 4
        self.assertEqual(self._base.A, 4)

    def test_str(self):
        kwarg_string = "A='a', B='b', key='MadeUpKey1'"
        self.assertEqual(kwarg_string,self._base.to_string() )

    def test_as_df_index(self):
        index = ['A','B','key']
        self.assertListEqual(list(self._base.to_df().index ), index)

    def test_as_df_values(self):
        self.assertListEqual(list(self._base.to_df()['Value'] ), ['a','b','MadeUpKey1'])
#
    def test_as_dict_values(self):
        keys = ['A', 'B', 'key']
        self.assertListEqual(keys, self._base.to_dict().keys())

    def test_as_dict_values(self):
        keys = ['a', 'b', 'MadeUpKey1']
        self.assertListEqual(keys, self._base.to_dict().values() )

    def test_update_properties(self):
        class New(pycotools._base._ModelBase):
            def __init__(self, model, **kwargs):
                super(New, self).__init__(model, **kwargs)

                options = {'A': 'not_a',
                           'B': 'b'}

                self.update_properties(options)
                #options.update(self.__dict__)
                self.list_of_output = []
                self.list_of_output.append(self.A )
                self.list_of_output.append(self.B )
                self.list_of_output.append(self.__dict__['A'] )
                self.list_of_output.append(self.__dict__['B'] )
        new_class = New(self.copasi_file, A='a')
        expected_output = ['a', 'b', 'a', 'b']
        self.assertListEqual(new_class.list_of_output, expected_output)

    def test_convert_bool_to_numeric(self):
        """

        :return:
        """
        class New(pycotools._base._ModelBase):
            def __init__(self, model, **kwargs):
                super(New, self).__init__(model, **kwargs)

                options = {'append': True,
                           'confirm_overwrite': False,
                           'output_event': False,
                           'scheduled': True,
                           'plot': True}

                options = self.convert_bool_to_numeric(options)
                self.update_properties(options)

        new_class = New(self.copasi_file, random_option=True)
        lst = [new_class.append, new_class.confirm_overwrite,
               new_class.plot]
        self.assertListEqual(lst, ['1', '0', True])


    def test_check_integrity(self):
        """

        :return:
        """
        class New(pycotools._base._ModelBase):
            def __init__(self, model, **kwargs):
                super(New, self).__init__(model, **kwargs)

                options = {'append': True,
                           'confirm_overwrite': False,
                           'output_event': False,
                           'scheduled': True,
                           'plot': True}

                options = self.convert_bool_to_numeric(options)
                self.update_properties(options)
                self.check_integrity(options.keys(), kwargs.keys())
        with self.assertRaises(pycotools.errors.InputError) as context:
            New(self.copasi_file, wrong_option=True)
        self.assertTrue(isinstance(context.exception, pycotools.errors.InputError))



class BaseModelTests(_test_base._BaseTest):
    def setUp(self):
        super(BaseModelTests, self).setUp()
        # self._model_base_from_string = pycotools._base._ModelBase(self.copasi_file, A='a')
        # self._model_base_from_element = pycotools._base._ModelBase(self._model_base_from_string.model, B='b')


    def test_from_path(self):
        """
        Test _ModelBase can get model from string to model
        path
        :return:
        """
        M = pycotools._base._ModelBase(self.copasi_file)
        self.assertTrue(isinstance(M.model, pycotools.model.Model) )


    def test_kwargs(self):
        M = pycotools._base._ModelBase(self.copasi_file, A='a')
        self.assertEqual(M.A, 'a')

    def test_to_string(self):
        M = pycotools._base._ModelBase(self.copasi_file)
        self.assertTrue(isinstance(M.to_string(), str))

    def test_inheritance(self):
        """

        :return:
        """

        class A(pycotools._base._ModelBase):
            def __init__(self, model, **kwargs):
                super(A, self).__init__(model, **kwargs)

                self.default_properties = {'a': 1,
                                           'b': 2}

                self.update_properties(self.default_properties)
                self.update_kwargs(self.default_properties)

            def __str__(self):
                return 'A({})'.format(self.to_string())

        class B(A):
            def __init__(self, model, **kwargs):
                super(B, self).__init__(model, **kwargs)

                self.default_properties = {'c': 3,
                                           'd': 4}

                self.update_properties(self.default_properties)
                self.update_kwargs(kwargs)

            def __str__(self):
                return "B({})".format(self.to_string())



        a = A(self.model)
        b = B(self.model, d=5)
        self.assertEqual(b.d, 5)


    def test_number_of_kwargs(self):
        """

        :return:
        """
        class A(pycotools._base._ModelBase):
            def __init__(self, model, **kwargs):
                super(A, self).__init__(model, **kwargs)

                self.default_properties = {'a': 1,
                                           'b': 2}

                self.update_properties(self.default_properties)
                self.update_kwargs(self.default_properties)

            def __str__(self):
                return 'A({})'.format(self.to_string())

        a = A(self.model, a=4)
        print a


            # def test_save(self):
    #     class New(pycotools._base._ModelBase):
    #         def __init__(self, model, **kwargs):
    #             super(New, self).__init__(model, **kwargs)
    #
    #             options = {'A': 'not_a',
    #                        'B': 'b'}
    #
    #             self.update_properties(options)
    #             for i in kwargs.keys():
    #                 assert i in options.keys(), '{} is not a keyword argument for Reports'.format(i)
    #             options.update(kwargs)
    #             self.kwargs = options
    #     new_class = New(self.copasi_file, A='a')
    #     new_class.save(self.copasi_file, )


if __name__ == '__main__':
    unittest.main()















