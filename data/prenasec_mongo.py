#!/usr/bin/env python

"""prenasec_mongo.py parses a wordlist from a source file to the Mongo database."""

__author__ = "Tommy Ch., Joly S."
__copyright__ = "Copyright 2021, SmartPasswd"
__license__ = "MIT"
__version__ = "1.0.1"
__maintainer__ = "Tommy Chu"
__email__ = "chutommy101@gmail.com"
__status__ = "Production"

import os
import sys

import pymongo


def retrieve_arguments():
    args = sys.argv
    if len(args) != 3:
        sys.exit(f"""Invalid arguments: {args}
    USAGE: python3 prenasec_sqlite.py [SOURCE FILE] [CONNECTION STRING]""")

    return args[1], args[2]


class WordList:
    """
    A class WordList represents a simple wordlist.

    Attributes
    ----------
     col : Collection
         The database collection to insert
     file: textIO
         The representation of the source text file
     """

    def __init__(self, source, conn):
        """
        Parameters
        ----------
        source : str
            The path and name of the source file
        conn : str
            The connection string of the Mongo database
        """
        if not os.path.exists(source):
            sys.exit(f"Source file `{source}` does not exists.")

        # source file
        with open(source, "r") as file:
            self.file = file

            # db connection
            with pymongo.MongoClient(conn) as client:
                db = client.get_default_database()

                self.col = db["words"]
                self.col.drop()
                self.col.create_index("word")

                # parse
                self.parse()

    def parse(self):
        """Parses the wordlist table."""
        words = []
        while True:
            line = self.file.readline()
            if not line:
                break
            words.append({"word": line.strip()})
        self.col.insert_many(words)


sourcePath, connectionStr = retrieve_arguments()
WordList(sourcePath, connectionStr)
