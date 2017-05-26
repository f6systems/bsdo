package main

import (
	"log"
	"os"

	"github.com/f6systems/bsdb"
	"github.com/f6systems/bsdo/dbsrv"
)

func init() {
	log.Printf("[init()] Starting.\n")
}

func main() {
	log.Printf("Starting.\n")
	//TODO:(hopley)  Logic around ... DB_DRIVER and DB_DSN from ENV or default to : 'sqlite3' and 'BSDefault.sqlite'...
	//DBConn := dbsrv.New("mysql", "root:wiq@tcp(127.0.0.1:6033)/ww")
	//TODO:(hopley) From environment get or use defaults (need to get for path) ; and BS_SQLDIR
	DB_DRIVER := os.Getenv("DB_DRIVER")
	if DB_DRIVER == "" {
		DB_DRIVER = "sqlite3"
	}
	DB_DSN := os.Getenv("DB_DSN")
	if DB_DSN == "" {
		DB_DSN = "db/bootstrapTest.db"
	}
	DBConn := dbsrv.New(DB_DRIVER,DB_DSN)

	err := DBConn.Open()
	if err != nil {
		log.Printf("ERROR: Issue Opening the database. (err=%v)\n", err)
	}

	err = bsdb.BSCheck(DBConn.DB)
	if err != nil {
		log.Printf("ERROR: Some issue with database. (err=%v)\n", err)
	}
	release := bsdb.BSRelease(DBConn.DB)
	log.Printf("<INFO> This database should move to the current release value = %d\n", release)
	DBConn.Mark("[bsdo:main()] Doing a Mark() in the mark table.")
	//TODO:(hopley) - make process to get the newer SQL files run in. 1) list of what to process
	log.Printf("Completed.\n")
}
