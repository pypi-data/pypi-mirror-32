import os
import re
from . import sql
from .sql import SQL, SQLTemplateRederer, SQLComposer
from .sqlmap import SQLMap
from .q import Q, V
from .d import D
from .f import F
from . import db3, pg2
from .utils import *

__version__ = "0.4.6"

sqlite3 = db3.open
postgres = pg2.open

class SQLPhile (SQLMap):
	def __init__ (self, dir = None, auto_reload = False, engine = "postgresql", conn = None):
		self._dir = dir
		self._auto_reload = auto_reload
		self._engine = engine
		self._conn = conn
		self._ns = {}
		self._dir and self._load_sqlmaps ()		
		
	def __getattr__ (self, name):
		try:
			return self._ns [name]
		except KeyError:
			return getattr (self._ns ["default"], name)
		
	def _load_sqlmaps (self):
		for fn in os.listdir (self._dir):
			if fn[0] == "#":
				continue
			ns = fn.split (".") [0]
			if ns == "ops":
				raise NameError ('ops cannot be used SQL map file name')
			self._ns [ns] = SQLMap (os.path.join (self._dir, fn), self._auto_reload, self._engine, self._conn)
	
			
def Template (backend = "postgresql", path = None, auto_reload = False):
	assert backend in ("postgresql", "sqlite3")
	if path and os.path.isdir (path):
		return SQLPhile (path, auto_reload, backend)
	return SQLMap (path, auto_reload, backend)

