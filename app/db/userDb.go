package db

import (
	database "web/modules/database/sqlite"
)

type UserDb struct {
	*database.BaseDb
}

// ユーザIDでユーザ情報取得
func (db *UserDb) GetUser(id int) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, account, password FROM user WHERE id = ?`,
		id,
	)
	return row
}

// ログインアカウントでユーザ情報取得
func (db *UserDb) GetUserByAccount(account string) map[string]interface{} {
	row := db.QueryRow(
		`SELECT id, account, password FROM user WHERE account = ?`,
		account,
	)
	return row
}
