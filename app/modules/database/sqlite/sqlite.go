package database

import (
	"web/config"
	"web/modules/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.Print("#db init")

	// DBコネクション取得
	//var db *sqlx.DB
	sqlxDb, err := sqlx.Connect("sqlite3", config.GetEnv().DatabaseName)
	defer sqlxDb.Close()
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		log.Print("#process exit by error")

		// 異常時は終了
		//log.Fatal(err)	// スタックトレースは出力しない =log.Fatal(err.Error())
		log.Error(err) // スタックトレースも出力
	}
}
