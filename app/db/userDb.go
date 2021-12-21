package db

import (
	database "web/modules/database/sqlite"
)

type UserDb struct {
	*database.BaseDb
}

// ログインアカウントでユーザ情報取得
func (db *UserDb) GetUser(account string) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, account FROM user WHERE account = ?`,
		account,
	)
	return row
}
