package db

import (
	"web/modules/log"
)

type UserDb struct{}

func init() {
	log.Println("#init userDb")
}

func (db *UserDb) Test() {
	log.Println("#Test()")
}
