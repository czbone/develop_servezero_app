package db

import (
	database "web/modules/database/sqlite"
	"web/modules/log"
)

type UserDb struct {
	*database.BaseDb
}

func (db *UserDb) Test() {
	rows, err := db.Query(
		`SELECT id, account FROM USER`,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var account string

		// カーソルから値を取得
		if err := rows.Scan(&id, &account); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}

		log.Printf("id: %d, accunt: %s\n", id, account)
	}
}
