package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"gopkg.in/gorp.v1"
	"log"
	"github.com/go-gorp/gorp"
)

var dbmap = initDb()

type Instruction struct {
	Id int64 `db:"id" json:"id"`
	EventStatus string `db:"event_status" json:"event_status"`
	EventName string `db:"event_name" json:"event_name"`
}

func initDb() *gorp.DbMap  {
	db, err := sql.Open("mysql", "root:pass@/instructions")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db:db, Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Instruction{}, "Instruction").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")
	return dbmap
}

func checkErr(err error, msg string)  {
	if err != nil {
		log.Fatal(msg, err)
	}
}

