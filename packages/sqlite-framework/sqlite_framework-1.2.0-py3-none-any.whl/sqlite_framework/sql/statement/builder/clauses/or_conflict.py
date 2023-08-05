from sqlite_framework.sql.item.constants.conflict_resolution import ConflictResolution
from sqlite_framework.sql.statement.builder.clauses.base import BaseClause


class OrClause(BaseClause):
    def __init__(self):
        super().__init__()
        self._or = None

    def or_(self, conflict_resolution: ConflictResolution):
        self._or = "or {conflict_resolution}".format(conflict_resolution=conflict_resolution.str())
        return self
