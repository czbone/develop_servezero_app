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

// ドメイン追加
func (db *DomainDb) AddNewDomain(name string, dir string) bool {
	_, err := db.Exec(
		`INSERT INTO domain (name, dir_name) VALUES (?, ?)`,
		name,
		dir,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
