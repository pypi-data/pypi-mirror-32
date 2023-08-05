# encoding: utf-8

"""
.. codeauthor:: Tsuyoshi Hombashi <tsuyoshi.hombashi@gmail.com>
"""

from __future__ import absolute_import, print_function, unicode_literals

from collections import OrderedDict
from textwrap import dedent

import pytest
import sqliteschema as ss
from sqliteschema._text_extractor import (
    SqliteSchemaTextExtractorV0, SqliteSchemaTextExtractorV1, SqliteSchemaTextExtractorV2,
    SqliteSchemaTextExtractorV3, SqliteSchemaTextExtractorV4, SqliteSchemaTextExtractorV5)

from .fixture import database_path


def patch_attr(self, table_name, schema_type):
    return ["'Primary Key ID' INTEGER", "real_attr REAL", "'AA BB CC' TEXT"]


class Test_SqliteSchemaTextExtractorV0(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV0

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()
        expected = dedent("""\
            testdb0
            testdb1
            constraints
            """)

        print("[expected]\n{}".format(expected))
        print("[actual]\n{}".format(output))

        assert output == expected

    def test_normal_get_table_schema(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema("testdb1")

        assert output == ['foo', 'bar', 'hoge']

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb0")

        assert output == "testdb0\n"

    def test_normal_get_table_schema_w_space(self, monkeypatch, database_path):
        monkeypatch.setattr(
            self.EXTRACTOR_CLASS, "_get_attr_schema", patch_attr)

        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema("testdb1")

        assert output == ['Primary Key ID', 'real_attr', 'AA BB CC']


class Test_SqliteSchemaTextExtractorV1(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV1

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()
        expected = dedent("""\
            testdb0 (attr_a, attr_b)
            testdb1 (foo, bar, hoge)
            constraints (primarykey_id, notnull_value, unique_value)
            """)

        print("[expected]\n{}".format(expected))
        print("[actual]\n{}".format(output))

        assert output == expected

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb0")

        assert output == """testdb0 (attr_a, attr_b)\n"""

    def test_normal_get_table_schema_w_space(self, monkeypatch, database_path):
        monkeypatch.setattr(
            self.EXTRACTOR_CLASS, "_get_attr_schema", patch_attr)

        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema("testdb1")

        assert output == ['Primary Key ID', 'real_attr', 'AA BB CC']


class Test_SqliteSchemaTextExtractorV2(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV2

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()
        expected = dedent("""\
            testdb0 (attr_a INTEGER, attr_b INTEGER)
            testdb1 (foo INTEGER, bar REAL, hoge TEXT)
            constraints (primarykey_id INTEGER, notnull_value REAL, unique_value INTEGER)
            """)

        print("[expected]\n{}".format(expected))
        print("[actual]\n{}".format(output))

        assert output == expected

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb0")

        assert output == """testdb0 (attr_a INTEGER, attr_b INTEGER)\n"""

    def test_normal_get_table_schema(self, monkeypatch, database_path):
        monkeypatch.setattr(
            self.EXTRACTOR_CLASS, "_get_attr_schema", patch_attr)

        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema("testdb1")

        assert output == OrderedDict([
            ('Primary Key ID', 'INTEGER'),
            ('real_attr', 'REAL'),
            ('AA BB CC', 'TEXT'),
        ])


class Test_SqliteSchemaTextExtractorV3(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV3

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()
        expected = dedent("""\
            testdb0 (attr_a INTEGER, attr_b INTEGER)
            testdb1 (foo INTEGER, bar REAL, hoge TEXT)
            constraints (primarykey_id INTEGER PRIMARY KEY, notnull_value REAL NOT NULL, unique_value INTEGER UNIQUE)
            """)

        print("[expected]\n{}".format(expected))
        print("[actual]\n{}".format(output))

        assert output == expected

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb1")

        assert output == """testdb1 (foo INTEGER, bar REAL, hoge TEXT)\n"""

    def test_normal_get_table_schema(self, monkeypatch, database_path):
        monkeypatch.setattr(
            self.EXTRACTOR_CLASS, "_get_attr_schema", patch_attr)

        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema("testdb1")

        assert output == OrderedDict([
            ('Primary Key ID', 'INTEGER'),
            ('real_attr', 'REAL'),
            ('AA BB CC', 'TEXT'),
        ])


class Test_SqliteSchemaTextExtractorV4(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV4

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()
        expected = dedent("""\
            testdb0 (
                attr_a INTEGER,
                attr_b INTEGER
            )

            testdb1 (
                foo INTEGER,
                bar REAL,
                hoge TEXT
            )

            constraints (
                primarykey_id INTEGER PRIMARY KEY,
                notnull_value REAL NOT NULL,
                unique_value INTEGER UNIQUE
            )

            """)

        print("[expected]\n{}".format(expected))
        print("[actual]\n{}".format(output))

        assert output == expected

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb1")

        assert output == dedent("""\
            testdb1 (
                foo INTEGER,
                bar REAL,
                hoge TEXT
            )
            """)


class Test_SqliteSchemaTextExtractorV5(object):
    EXTRACTOR_CLASS = SqliteSchemaTextExtractorV5

    def test_normal_dumps(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.dumps()

        print("[actual]\n{}".format(output))

        assert len(output) > 180

    def test_normal_get_table_schema_text(self, database_path):
        extractor = self.EXTRACTOR_CLASS(database_path)
        output = extractor.get_table_schema_text("testdb1")

        assert len(output) > 50


class Test_SqliteSchemaTextExtractor_error(object):

    @pytest.mark.parametrize(["extractor_class"], [
        [SqliteSchemaTextExtractorV0],
        [SqliteSchemaTextExtractorV1],
        [SqliteSchemaTextExtractorV2],
        [SqliteSchemaTextExtractorV3],
        [SqliteSchemaTextExtractorV4],
        [SqliteSchemaTextExtractorV5],
    ])
    def test_exception_constructor(self, extractor_class):
        with pytest.raises(ValueError):
            extractor_class(None)

        with pytest.raises(IOError):
            extractor_class("not_exist_path")

    @pytest.mark.parametrize(["extractor_class"], [
        [SqliteSchemaTextExtractorV0],
        [SqliteSchemaTextExtractorV1],
        [SqliteSchemaTextExtractorV2],
        [SqliteSchemaTextExtractorV3],
        [SqliteSchemaTextExtractorV4],
        [SqliteSchemaTextExtractorV5],
    ])
    def test_exception_get_table_schema_text(
            self, extractor_class, database_path):
        extractor = extractor_class(database_path)

        with pytest.raises(ss.DataNotFoundError):
            print(extractor.get_table_schema_text("not_exist_table"))
