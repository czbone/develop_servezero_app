#!/bin/sh

sqlite3 test.db  "CREATE TABLE user (id INTEGER PRIMARY KEY, account TEXT, password TEXT);"
sqlite3 test.db  "INSERT INTO user (account, password) values ('admin@example.com', 'admin');"
sqlite3 test.db  "SELECT * FROM user";
