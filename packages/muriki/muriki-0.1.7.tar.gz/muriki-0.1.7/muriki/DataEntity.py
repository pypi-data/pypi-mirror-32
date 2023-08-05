#!/usr/bin/env python
# =====================================================================
# Copyright 2018 Humberto Ramos Costa. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or 
# implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# =====================================================================

import csv
import re
import decimal
import traceback

from decimal import *
from gettext import gettext as _
from aenum import Enum, skip
from datetime import *
from muriki.DataProperty import *
#~ import ( DataProperty, generate_error_message, data_property
                            #~ , auto_data_property, Sql)


#import psycopg2 // will be needed if postgre database is used

#from .DataProperty import *

"""
This utilty is intended to use to ease the import of data from files and
the persisntence into databases. Options are made to make those
resources easy to insert into existent code (using decorators)
and writing a minimum of code.
"""


def data_entity(
        cls=None,
        database_engine=None,
        database_server_address=None,
        database_user=None,
        database_password=None,
        database_name=None,
        database_port=None,
        fl_regex=None):
    """
    This class decorator is to be used to allow an class to be easily
    used to import data from files and persist this data

    :param cls: The class to be decorated, passed automatically to the
        decorator
    :param database_engine: Database engine to be used to
        persist data, see the DataProperty.Sql.DatabaseEngine
    :param database_server_address: network address of the
        server that will be used to persist data
    :param database_user: user to persist data, if necessary
    :param database_password: password to grant access to
        the user on the database
    :param database_port: TCP port of the database server to
        be used in the connection
    :param fl_regex: Regex to be used to identify lines with
        registers to be imported into instances of the
        specified class

    """


    @classmethod
    def create_table(cls):
        """
        Simply use the execute_sql function to create int the database
        that correspond to the class
            :param cls: The class passed by the decorator
        """
        cls.execute_sql(cls.create_table_sql())

    
    @classmethod
    def execute_sql(cls, sql=None, values=None, commit=True):
        """
        Execute arbitrary sql code, the code is generated in other functions
        based on the sqlEngine parameter
            :param cls: The class passed by the decorator
        """
        try:
            result = cls._cursor.execute(sql, values)
        except Exception as e:
            raise Exception(
                'Sql Error:' + (
                    cls._cursor .mogrify(sql, values).decode('utf-8'))
                ) from e

        if commit:
            cls._connection.commit()
        return result

    @classmethod
    def sql_columns(cls):
        """
        search the properties decorated, get those with sql_column_type
        defined and put in the correspondent list. The separation is
        needed to create the columns in a certain order (keys, indexes
        ,data) to improve performance
            :param cls: The class passed by the decorator
        """
        columns = {
            Sql.ColumnType.AUTO_INCREMENT: [],
            Sql.ColumnType.PRIMARY_KEY: [],
            Sql.ColumnType.INDEX: [],
            Sql.ColumnType.SIMPLE: [],
        }
        for k, cls_member in vars(cls).items():
            if (not (cls_member is None)
                    and isinstance(cls_member, DataProperty)
                    and getattr( cls_member, 'sql_column_type', None) is not None
                    ):
                columns[cls_member.sql_column_type].append(
                    cls_member)
        return columns

    @classmethod
    def fl_properties(cls):
        """
        search the properties decorated, get those with fl_start
        member defined, order by this member and return the list.
        The order is important in a fixed length import
            :param cls: The class passed by the decorator
        """
        flp = [p for p in vars(cls).values()
               if ((p is not None) and isinstance(p, DataProperty)
                   and not(p.fl_start is None))
               ]
        flp.sort(key=lambda f: f.fl_start)
        return flp


    @classmethod
    def csv_properties(cls):
        """
        search the properties decorated, get those with csv_position
        member defined, order by this member and return the list.
        The order is important in a fixed length import
            :param cls: The class passed by the decorator
        """

        csvp = [p for p in vars(cls).values()
               if ((p is not None) and isinstance(p, DataProperty)
                   and not(p.csv_position is None))
               ]
        csvp.sort(key=lambda f: f.csv_position)

        return csvp

    @classmethod
    def create_table_sql(cls):
        """
        Generates the string (based on the properties passed to the
        decorator) used to create the corresponding table 
            :param cls: The class passed by the decorator
        """

        tmp = cls._database_engine.create_table.value + \
            ' ' + (cls.__name__.lower() + ' ( ')
        
        aditional_primary_key = ''
        aditional_index = ''
        
        ## The sql columns ared inserted in a specific order
        ## to improve performance
        for c in cls.sql_columns()[Sql.ColumnType.AUTO_INCREMENT]:
            aditionar_primary_key = aditionar_primary_key + c._name + ', '
            tmp = tmp + get_sql_definition(cls, c) + ', '
        for c in cls.sql_columns()[Sql.ColumnType.PRIMARY_KEY]:
            tmp = tmp + get_sql_definition(cls, c) + ', '
            aditional_primary_key = aditional_primary_key + c._name + ', '
        for c in cls.sql_columns()[Sql.ColumnType.INDEX]:
            tmp = tmp + get_sql_definition(cls, c) + ', '
            aditional_index = (aditional_index + (cls._database_engine.index
                                                  .value.format(cls.__name__
                                                    , c._name)))
        for c in cls.sql_columns()[Sql.ColumnType.SIMPLE]:
            tmp = tmp + get_sql_definition(cls, c) + ', '
            
        ## format the string inserting aditional keys and indexes
        ## clauses
        tmp = (tmp
               + ((cls._database_engine.primary_key.value
                   .format(aditional_primary_key))
                  if aditional_primary_key
                  else '')
               + ');'
               + aditional_index)
        #clean the unused extra commas
        tmp = re.sub(',\s*\)', ' )', tmp)
        return tmp

    @classmethod
    def insert_sql(cls):
        """
        Generates the string (based on the properties passed to the
        decorator) used to create one register on the database, this
        register will 'fit' in a database created by create_table_sql
        function
            :param cls: The class passed by the decorator
        """
        names = []
        auto_increment_field_name = ''
        for k, sql_columns_by_type in cls.sql_columns().items():
            for sql_column in sql_columns_by_type:
                if(sql_column.data_type != auto_increment):
                    names.append(sql_column._name)
                else:
                    auto_increment_field_name = sql_column._name

        insert_sql = cls._database_engine.insert.value.format(
            cls.__name__.lower(),
            ', '.join(names),
            '%s,' * len(names),
            ('RETURNING ' + auto_increment_field_name
                if auto_increment_field_name else ''))

        insert_sql = re.sub(',\s*\)', ' )', insert_sql)

        return insert_sql

    @classmethod
    def batch_insert(cls, instances=[]):
        """
        Creates a 'huge' string to insert a lot of registers in just one
        execution.
            :param cls: The class passed by the decorator
            :para instances: a list of the objects to be persisted
        """
        
        batch_sql = ''
        for instance in instances:
            batch_sql = batch_sql + instance.insert_sql_with_values()

        cls.execute_sql(batch_sql)

    #TODO - Define what to do when there is an error in a property (add parameter to choose?)
    @classmethod
    def create_from_csv(
            cls,
            file_name=None,
            headers=False,
            delimiter='\t',
            quotechar=None):
        """
        Creates a list of instances of objects from the defined entity from an csv file
            :param cls: The class passed by the decorator
            :param file_name: The name of the file (csv formatted)
            :param headers: Wheter the file has a first row with column names
            :param delimiter: The delimiter of the fields in the csv file
            :param quotechar: The char used to 'enclose' the values, it's 
                needed to allow the delimiter to occur inside the fields
        """
        #~ with (open( file_name ) as csv_file:
            #~ csv_reader = csv.DictReader( fieldnames= )
        entities = []
        csv_file = open(file_name)
        csv_reader = csv.reader(csv_file, delimiter=delimiter, quotechar=quotechar)
        for row in csv_reader:
            entity = cls()
            entity.properties_from_csv( row )
            entities.append( entity )
        return ( entities )

    def properties_from_csv(self, values=None):
        """
        Set the values of properties of an object (self) with a list of
            :param cls: The class passed by the decorator
            :param values: The list of values to be setted in the properties
        """
        
        for p in self.__class__.csv_properties():
            setattr(self, p._name, values[ p.csv_position] )
            
    @classmethod
    def create_from_fl(cls, file_name=None, encoding="utf-8"):
        """
        Create a list of objects from the type (cls) reading the values from
        a fixed length file
            :param cls: The class passed by the decorator
            :param file_name: The name (and location) of the file to be
                readed
            :param encoding: The encoding to be used to read the file
        """
        entities = []
        fl_file = open(file_name, 'r', encoding=encoding)
        for nr_linha, row in enumerate(fl_file):
            if (re.match(fl_regex, row)):
                try:
                    entity = cls()
                    entity.properties_from_fl(row)
                    entities.append(entity)
                except BaseException:
                    #~ print('error in line: ' + row)
                    traceback.print_exc()
                    pass
        return entities

    def properties_from_fl(self, string=None):
        """
        Set the values of the properties from a fixed length string
            :param self: The object which the values will be set
            :param string: The string where the values will be readed
        """
        for p in self.__class__.fl_properties():
            # print( str( type ( p.fl_start )))
            tmp = string[p.fl_start: (p.fl_start + p.fl_length)].strip()
            setattr(self, p._name, tmp)
            
    def insert_sql_with_values(self):
        """
        Generate an sql string with the values, read to be inserted in the database
        """

        values = []
        for k, sql_columns_by_type in self.__class__.sql_columns().items():
            for sql_column in sql_columns_by_type:
                if(sql_column.data_type != auto_increment):
                    values.append(self.get_data_sql_value(sql_column))

        return (
            self.__class__._cursor.mogrify(
                self.__class__.insert_sql(),
                values).decode('utf-8'))

    def insert(self, commit=True):
        """
        Insert the values of the object (self) in the database, note that
            :param commit: commit directly after executing the sql
        """
        values = []

        for k, sql_columns_by_type in self.__class__.sql_columns().items():
            for sql_column in sql_columns_by_type:
                if(sql_column.data_type != auto_increment):
                    values.append(self.get_data_sql_value(sql_column))

        _id = self.__class__.execute_sql(
            self.__class__.insert_sql(), values, commit)

        return _id

    def get_data_sql_value(self, data_property=None):
        """
        Read the value of a sql column, if can't read the value, and there
        is an default, the default is returned, None otherwise
            :param data_property: Property to be readed
        """
        value = getattr(self, data_property._name,
                        getattr(data_property, 'default_value', None)
                        )
        return value

    def get_sql_definition(cls=None, cls_member=None):
        """
        Return the sql definition of a givern member of the class
            :param cls: The class passed by the decorator
            :param cls_member: The member which the sql should be generated
        """

        if(cls_member.sql_column_type
                == Sql.ColumnType.AUTO_INCREMENT):
            return (cls._database_engine.auto_increment.value
                    .format(cls_member._name))
        else:
            python_data_type = getattr(cls_member, 'data_type', str)
            sql_data_type = (cls._database_engine[
                python_data_type.__name__].value)
            if(python_data_type == str):
                length = getattr(cls_member, 'max_length',
                                 getattr(cls_member, 'length', 0)
                                 )
                if(length >= BIG_TEXT_MIN_LEN):
                    sql_data_type = (cls._database_engine.bigtext.value)
                elif(length > 0):
                    sql_data_type = (sql_data_type + '('
                                     + str(length) + ')')
            elif(python_data_type == int):
                max_value = getattr(cls_member, 'max_exclusive',
                                    getattr(cls_member, 'max_inclusive', 0)
                                    )
                if(max_value > INT_MAX_VALUE):
                    sql_data_type = (cls._database_engine.biginteger.value)

                try:
                    if(cls_member.fl_length > 9):
                        # print( 'adjusting to biginteger' + str( cls_member) )
                        sql_data_type = (cls._database_engine.biginteger.value)
                except BaseException:
                    pass

            elif(python_data_type == Decimal):
                total_digits = getattr(
                    cls_member, 'total_digits', DEFAULT_DECIMAL_TOTAL_DIGITS)
                fraction_digits = getattr(
                    cls_member,
                    'fraction_digits',
                    DEFAULT_DECIMAL_FRACTION_DIGITS)
                sql_data_type = (
                    sql_data_type + '(' + str(total_digits) + ', '
                    + str(fraction_digits) + ')')
            return sql_data_type.format(cls_member._name)

    if (cls is None):
    ## This is the chunck that really decorates the class, inserting the
    ## various methods defined above

        def wrapper(cls):
            cls._database_engine = database_engine
            cls.sql_columns = sql_columns
            cls.create_table_sql = create_table_sql
            cls.create_table = create_table
            cls.execute_sql = execute_sql
            cls.insert = insert
            cls.get_data_sql_value = get_data_sql_value
            cls.insert_sql = insert_sql
            cls.batch_insert = batch_insert
            cls.insert_sql_with_values = insert_sql_with_values
            cls.create_from_csv = create_from_csv
            cls.properties_from_fl = properties_from_fl
            cls.create_from_fl = create_from_fl
            cls.fl_properties = fl_properties
            cls.create_from_csv = create_from_csv
            cls.properties_from_csv=properties_from_csv
            cls.csv_properties = csv_properties

            cls.fl_regex = fl_regex

            if not(database_name is None):
                if( cls._database_engine == Sql.Mysql ):
                    # TODO - Connect to the mysql database
                    pass
                elif( cls._database_engine == Sql.Postgresql ):
                    # TODO - Maybe not the best strategy
                    import psycopg2
                    cls._connection = psycopg2.connect(
                        dbname=database_name,
                        user=database_user,
                        password=database_password,
                        host=database_server_address,
                        port=database_port)
                    cls._cursor = cls._connection.cursor()
            return cls
        return wrapper
    return cls
