from sqlite_framework.sql.item.expression.parser import EXPRESSION_TYPE, ExpressionParser
from sqlite_framework.sql.statement.builder.clauses.base import BaseClause


class WhereClause(BaseClause):
    def __init__(self):
        super().__init__()
        self._where = None

    def where(self, expr: EXPRESSION_TYPE):
        expr = ExpressionParser.parse(expr)
        self._where = "where {expr}".format(expr=expr.str())  # unsafe formatting
        return self
