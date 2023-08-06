#
# Copyright 2017 Import.io
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
import logging

import petl
import pymysql

from importio2.commands import AdBase

logger = logging.getLogger(__name__)


class CsvToDatabase(AdBase):
    def __init__(self):
        super(CsvToDatabase, self).__init__()
        self._db_user = None
        self._db_password = None
        self._db_database = None
        self._db_host = None
        self._db_table = None
        self._csv_path = None
        self._table = None
        self._create = None
        self._append = None

    def cli_description(self):
        return 'Loads a CSV to specified file in a database'

    def handle_arguments(self):
        self._parser.add_argument('-u', '--user', action='store', dest='db_user', required=True, metavar='user',
                                  help='User name to use for authentication to the database')
        self._parser.add_argument('-p', '--password', action='store', dest='db_password', required=True,
                                  metavar='password',
                                  help='Password to use for authentication to the database')
        self._parser.add_argument('-d', '--database', action='store', dest='db_database', required=True,
                                  metavar='hostname',
                                  help='Database to use')
        self._parser.add_argument('-i', '--host', action='store', dest='db_host', required=True, metavar='hostname',
                                  help='Hostname or IP address of the database')
        self._parser.add_argument('-t', '--table', action='store', dest='db_table', required=True, metavar='table_name',
                                  help='Name of the table to insert the data into')
        self._parser.add_argument('-f', '--csv-path', action='store', dest='csv_path', required=True, metavar='path',
                                  help='Path to CSV file')
        group = self._parser.add_mutually_exclusive_group(required=False)
        group.add_argument('-a', '--append', action='store_true', dest='append', default=False,
                           help='Flag to append data')
        group.add_argument('-c', '--create', action='store_true', dest='create', default=False,
                           help='Flag to create data')

        super(CsvToDatabase, self).handle_arguments()

    def get_arguments(self):
        super(CsvToDatabase, self).get_arguments()
        self._db_user = self._args.db_user
        self._db_password = self._args.db_password
        self._db_database = self._args.db_database
        self._db_host = self._args.db_host
        self._db_table = self._args.db_table
        self._csv_path = self._args.csv_path

        self._append = self._args.append
        self._create = self._args.create

    def load_data(self):
        """
        Loads data from CSV into specified database
        :return: None
        """
        logger.info("loading csv file: {0}".format(self._csv_path))
        table = petl.fromcsv(self._csv_path)
        logger.info("Connecting to database: {0}".format(self._db_host))
        connection = pymysql.connect(host=self._db_host,
                                     user=self._db_user,
                                     password=self._db_password,
                                     database=self._db_database,
                                     charset='utf8')
        connection.cursor().execute('SET SQL_MODE=ANSI_QUOTES')

        # If append option is set the add the CSV file to existing table
        if self._append:
            logger.info("Appending data to table: {0}".format(self._db_table))
            petl.appenddb(table, connection, self._db_table)
        else:
            # Pass in flag that indicates creating the table or replacing the contents
            if self._create:
                logger.info("Creating and adding data to table: {0}".format(self._db_table))
            else:
                logger.info("Replacing data for table: {0}".format(self._db_table))
            petl.todb(table, connection, self._db_table, create=self._create)

    def run(self, user, password, database, host, table, csv_path, append=False, create=False):
        """
        Call to perform action to load a CSV file to the database
        :param user: User to authenticate against the database
        :param password: Password to use to authenticate against the database
        :param database: Database to move to after login
        :param host: Location of the database host
        :param table: Table to add CSV file to
        :param csv_path: Path to CSV file to load
        :param append: Flag that indicates to add the data to the existing table
        :param create: Create the table and populate with the CSV data. Error if table already exists
        :return: None
        """
        self._db_user = user
        self._db_password = password
        self._db_database = database
        self._db_host = host
        self._db_table = table
        self._csv_path = csv_path
        self._append = append
        self._create = create

        self.load_data()

    def execute(self):
        """
        Performs the command via the CLI
        :return:
        """
        self.handle_arguments()
        self.load_data()


def main():
    """
    Functioned called by the entry point
    :return:
    """
    o = CsvToDatabase()
    o.execute()


if __name__ == '__main__':
    main()
