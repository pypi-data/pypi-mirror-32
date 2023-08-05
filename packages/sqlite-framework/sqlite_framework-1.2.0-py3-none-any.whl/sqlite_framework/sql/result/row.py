import sqlite3

from sqlite_framework.sql.util.column import ColumnUtil


class ResultRow(sqlite3.Row):
    def __getitem__(self, item):
        item = ColumnUtil.name_if_column(item)
        return super().__getitem__(item)

    def map(self, func: callable):
        return func(self)
