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
		`SELECT id, account, password FROM user WHERE account = ?`,
	)
	return rows
}
