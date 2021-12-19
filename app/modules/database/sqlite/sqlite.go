package database

import (
	"database/sql"
	"os"
	"web/config"
	"web/modules/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type BaseDb struct {
	*sqlx.DB
}

// DB接続のインスタンス
var sqlxDb *sqlx.DB

func init() {
	// DBファイルの存在確認
	_, err := os.Stat(config.GetEnv().DatabaseName)
	checkErr(err)

	// DBコネクション取得
	sqlxDb, err = sqlx.Connect("sqlite3", config.GetEnv().DatabaseName)
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		// 異常時は終了
		log.FatalObject(err) // スタックトレースも出力
	}
}

func (db *BaseDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := sqlxDb.Query(query, args...)

	return rows, err
}
