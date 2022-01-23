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
		`SELECT id, name, dir_name, domain_id, created_dt FROM domain ORDER BY id`,
	)
	return rows
}

// IDでドメイン取得
func (db *DomainDb) GetDomain(id int) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, name, dir_name, domain_id, created_dt FROM domain WHERE id = ?`,
		id,
	)
	return row
}

// ドメイン取得
func (db *DomainDb) GetDomainByName(name string) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, name, dir_name, domain_id, created_dt FROM domain WHERE name = ?`,
		name,
	)
	return row
}

// ドメイン追加
func (db *DomainDb) AddDomain(name string, dir string, domainId string) bool {
	_, err := db.Exec(
		`INSERT INTO domain (name, dir_name, domain_id, created_dt) VALUES (?, ?, ?, datetime('now', 'localtime'))`,
		name,
		dir,
		domainId,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}

// ドメイン削除
func (db *DomainDb) DelDomain(id int) bool {
	_, err := db.Exec(
		`DELETE FROM domain WHERE id = ?`,
		id,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
