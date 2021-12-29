package db

import (
	database "web/modules/database/sqlite"
)

type DomainDb struct {
	*database.BaseDb
}

// ドメイン一覧取得
func (db *DomainDb) GetDomainList() []map[string]interface{} {
	rows := db.QueryRows(
		`SELECT id, name, dir_name FROM domain ORDER BY id`,
	)
	return rows
}

// ドメイン取得
func (db *DomainDb) GetDomainByName(name string) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, name, dir_name FROM domain WHERE name = ?`,
		name,
	)
	return row
}
