#!/bin/sh

sqlite3 init.sqlite3  "CREATE TABLE user (id INTEGER PRIMARY KEY, account TEXT, password TEXT);"
sqlite3 init.sqlite3  "CREATE TABLE domain (id INTEGER PRIMARY KEY, name TEXT, dir_name TEXT, domain_id TEXT, created_dt TEXT);"
sqlite3 init.sqlite3  "INSERT INTO user (account, password) values ('admin@example.com', '\$2a\$10\$qc8rQ5f9NWL5FVKDUghr.ejS3sMiVT/.RFYwHVhfHiuSudSaBNnxa');" # admin
sqlite3 init.sqlite3  "SELECT * FROM user";
