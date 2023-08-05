from sqlite_framework.sql.item.constants.order_mode import OrderMode
from sqlite_framework.sql.item.expression.parser import EXPRESSION_TYPE, ExpressionParser
from sqlite_framework.sql.statement.builder.clauses.base import BaseClause


class OrderByClause(BaseClause):
    def __init__(self):
        super().__init__()
        self._order_by = None

    def order_by(self, expr: EXPRESSION_TYPE, mode: OrderMode = None):
        expr = ExpressionParser.parse(expr)
        self._order_by = "order by {expr}".format(expr=expr.str())
        if mode is not None:
            self._order_by += " {mode}".format(mode=mode.str())
        return self
