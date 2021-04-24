#!/usr/bin/env python

"""Prenasec.py parses a word list from a source file to the SQLite3 database."""

__author__ = "Tommy Ch., Joly S."
__copyright__ = "Copyright 2021, SmartPasswd"
__license__ = "MIT"
__version__ = "1.0.0"
__maintainer__ = "Tommy Chu"
__email__ = "chutommy101@gmail.com"
__status__ = "Production"

import os
import sqlite3
import sys


def retrieve_arguments():
    args = sys.argv
    if len(args) != 2:
        sys.exit(f"""Invalid arguments: {args}
    USAGE: python3 prenasec.py [SOURCE]""")

    src = args[1]
    source_path = os.path.join("raw", src)

    trg = os.path.splitext(src)[0] + ".db"
    target_path = os.path.join("parsed", trg)

    return source_path, target_path


class WordList:
    """
    A class WordList represents a simple word list.

    Attributes
    ----------
    conn : Connection
        The database connection
    db : Cursor
        The cursor for the selected database
    table : str
        Name of the table which is being populated
    file: textIO
        The representation of the source text file
    """

    def __init__(self, source, target, table):
        """
        Parameters
        ----------
        source : str
            The path and name of the source file
        target : str
            The path and name of the target file
        table: str
            The name of the table
        """
        if not os.path.exists(source):
            sys.exit(f"Source file `{source}` does not exists.")

        # db connection
        conn = sqlite3.connect(target, timeout=3.0)
        self.conn = conn
        self.db = conn.cursor()

        self.table = table

        # source file
        file = open(source, "r")
        self.file = file

        # parse
        self.parse()
        self.close()

    def init_table(self):
        """Initialises the word list table in the database."""
        self.db.execute(f"DROP TABLE IF EXISTS {self.table}")
        self.db.execute(f"""CREATE TABLE {self.table}
            (id INTEGER PRIMARY KEY AUTOINCREMENT, word TEXT NOT NULL UNIQUE)""")

    def add_word(self, word):
        """Inserts a word into the word list table."""
        self.db.execute(f"INSERT INTO {self.table} ('word') VALUES ('{word}')")

    def parse(self):
        """Parses the word list table."""
        self.init_table()
        while True:
            line = self.file.readline()
            if not line:
                break
            self.add_word(line.strip())
        self.conn.commit()

    def close(self):
        """Close all files and in-memory connections."""
        self.db.close()
        self.conn.close()
        self.file.close()


sourcePath, targetPath = retrieve_arguments()
WordList(sourcePath, targetPath, "words")
