package database

import (
	"database/sql"
	"os"
	"strconv"
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

// テーブルから複数のレコードを取得
// 連想配列の配列でデータを取得する。エラーの場合はnilが返る。
func (db *BaseDb) QueryRows(query string, args ...interface{}) []map[string]interface{} {
	rs, err := sqlxDb.Query(query, args...)
	if err != nil {
		log.ErrorObject(err)
		return nil
	}

	defer rs.Close()

	col, colErr := rs.Columns()
	if colErr != nil {
		log.ErrorObject(colErr)
		return nil
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		log.ErrorObject(err)
		return nil
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			setColVarType(&colVar, i, typeVal[i].DatabaseTypeName())
		}

		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			log.ErrorObject(err)
			return nil
		}

		for j := 0; j < len(col); j++ {
			setResultValue(&result, col[j], colVar[j], typeVal[j].DatabaseTypeName())
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		log.ErrorObject(err)
		return nil
	}

	return results
}
func setColVarType(colVar *[]interface{}, i int, typeName string) {
	switch typeName {
	case "INT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "TINYINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "MEDIUMINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "SMALLINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "BIGINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "FLOAT":
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case "DOUBLE":
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case "DECIMAL":
		var s []uint8
		(*colVar)[i] = &s
	case "DATE":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TIME":
		var s sql.NullString
		(*colVar)[i] = &s
	case "YEAR":
		var s sql.NullString
		(*colVar)[i] = &s
	case "DATETIME":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TIMESTAMP":
		var s sql.NullString
		(*colVar)[i] = &s
	case "VARCHAR":
		var s sql.NullString
		(*colVar)[i] = &s
	case "MEDIUMTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "LONGTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TINYTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	default:
		var s interface{}
		(*colVar)[i] = &s
	}
}
func setResultValue(result *map[string]interface{}, index string, colVar interface{}, typeName string) {
	switch typeName {
	case "INT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "TINYINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "MEDIUMINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "SMALLINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "BIGINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "FLOAT":
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case "DOUBLE":
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case "DECIMAL":
		if len(*(colVar.(*[]uint8))) < 1 {
			(*result)[index] = nil
		} else {
			(*result)[index], _ = strconv.ParseFloat(string(*(colVar.(*[]uint8))), 64)
		}
	case "DATE":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TIME":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "YEAR":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "DATETIME":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TIMESTAMP":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "VARCHAR":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "MEDIUMTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "LONGTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TINYTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	default:
		(*result)[index] = colVar
	}
}
