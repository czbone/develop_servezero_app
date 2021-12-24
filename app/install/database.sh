#!/bin/sh

sqlite3 zero.db  "CREATE TABLE user (id INTEGER PRIMARY KEY, account TEXT, password TEXT);"
sqlite3 zero.db  "INSERT INTO user (account, password) values ('admin@example.com', '\$2a\$10\$qc8rQ5f9NWL5FVKDUghr.ejS3sMiVT/.RFYwHVhfHiuSudSaBNnxa');" # admin
sqlite3 zero.db  "SELECT * FROM user";
