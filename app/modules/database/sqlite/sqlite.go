package database

import (
	"os"
	"web/config"
	"web/modules/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.Print("#db init")

	// DBファイルの存在確認
	_, err := os.Stat(config.GetEnv().DatabaseName)
	checkErr(err)

	// DBコネクション取得
	//var db *sqlx.DB
	sqlxDb, err := sqlx.Connect("sqlite3", config.GetEnv().DatabaseName)
	checkErr(err)
	defer sqlxDb.Close()
}
func checkErr(err error) {
	if err != nil {
		// 異常時は終了
		log.FatalObject(err) // スタックトレースも出力
	}
}
