#!/bin/sh

sqlite3 zero.db  "CREATE TABLE user (id INTEGER PRIMARY KEY, account TEXT, password TEXT);"
sqlite3 zero.db  "CREATE TABLE domain (id INTEGER PRIMARY KEY, name TEXT, dir_name TEXT, created_dt TEXT);"
sqlite3 zero.db  "INSERT INTO user (account, password) values ('admin@example.com', '\$2a\$10\$qc8rQ5f9NWL5FVKDUghr.ejS3sMiVT/.RFYwHVhfHiuSudSaBNnxa');" # admin
sqlite3 zero.db  "SELECT * FROM user";
