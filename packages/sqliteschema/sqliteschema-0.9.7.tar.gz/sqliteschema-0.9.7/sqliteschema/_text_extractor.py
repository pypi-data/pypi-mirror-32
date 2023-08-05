#!/usr/bin/env python
# encoding: utf-8

"""
.. codeauthor:: Tsuyoshi Hombashi <tsuyoshi.hombashi@gmail.com>
"""

from __future__ import absolute_import, unicode_literals

import six
import simplesqlite as sqlite
import typepy

from ._error import DataNotFoundError
from ._interface import AbstractSqliteSchemaExtractor


class SqliteSchemaTextExtractorV0(AbstractSqliteSchemaExtractor):

    @property
    def verbosity_level(self):
        return 0

    def get_table_schema_text(self, table_name):
        self._validate_table_existence(table_name)

        return "{:s}\n".format(table_name)

    def _write_database_schema(self):
        for table_name in self.get_table_name_list():
            if table_name in sqlite.SQLITE_SYSTEM_TABLE_LIST:
                continue

            self._stream.write(self.get_table_schema_text(table_name))


class SqliteSchemaTextExtractorV1(SqliteSchemaTextExtractorV0):

    @property
    def verbosity_level(self):
        return 1

    def get_table_schema_text(self, table_name):
        return "{:s} ({:s})\n".format(
            table_name, ", ".join(self.get_table_schema(table_name)))


class SqliteSchemaTextExtractorV2(AbstractSqliteSchemaExtractor):

    @property
    def verbosity_level(self):
        return 2

    def get_table_schema(self, table_name):
        return self._get_table_schema_v1(table_name)

    def get_table_schema_text(self, table_name):
        attr_list = []
        for key, value in six.iteritems(self.get_table_schema(table_name)):
            attr_list.append("{:s} {:s}".format(key, value))

        return "{:s} ({:s})\n".format(table_name, ", ".join(attr_list))

    def _write_table_schema(self, table_name):
        self._stream.write(self.get_table_schema_text(table_name))

    def _write_database_schema(self):
        for table_name in self.get_table_name_list():
            self._write_table_schema(table_name)


class SqliteSchemaTextExtractorV3(SqliteSchemaTextExtractorV2):

    @property
    def verbosity_level(self):
        return 3

    def get_table_schema(self, table_name):
        return self._get_table_schema_v2(table_name)


class SqliteSchemaTextExtractorV4(SqliteSchemaTextExtractorV3):

    @property
    def verbosity_level(self):
        return 4

    def get_table_schema_text(self, table_name):
        attr_list = []
        for key, value in six.iteritems(self.get_table_schema(table_name)):
            attr_list.append("{:s} {:s}".format(key, value))

        return "\n".join([
            "{:s} (".format(table_name),
        ] + [
            ",\n".join([
                "    {:s}".format(attr)
                for attr in attr_list
            ])
        ] + [
            ")\n",
        ])

    def _write_table_schema(self, table_name):
        super(SqliteSchemaTextExtractorV4, self)._write_table_schema(
            table_name)
        self._stream.write("\n")


class SqliteSchemaTextExtractorV5(SqliteSchemaTextExtractorV4):
    __ENTRY_TYPE_LIST = ["table", "index"]

    @property
    def verbosity_level(self):
        return 5

    def get_table_schema_text(self, table_name):
        schema_text = super(
            SqliteSchemaTextExtractorV5,
            self
        ).get_table_schema_text(table_name)

        try:
            index_schema = self._get_index_schema(table_name)
        except DataNotFoundError:
            return schema_text

        index_schema_list = [
            "{}".format(index_entry)
            for index_entry in index_schema
            if typepy.is_not_null_string(index_entry)
        ]

        if typepy.is_empty_sequence(index_schema_list):
            return schema_text

        return "{:s}{:s}\n".format(schema_text, "\n".join(index_schema_list))
