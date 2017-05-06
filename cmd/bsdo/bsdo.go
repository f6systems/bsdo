package main

import (
	"log"

	"github.com/f6systems/bsdo/dbsrv"
	"github.com/f6systems/bsdb"
)

func init() {
	log.Printf("[init()] Starting.\n")
}

func main() {
	log.Printf("Starting.\n")
	DBConn := dbsrv.New("root:wiq@tcp(127.0.0.1:3306)/ww")

	err := DBConn.Open()
	if err != nil {
		log.Printf("ERROR: Issue Opening the database. (err=%v)\n", err)
	}

	err = bsdb.BSCheck(DBConn.DB)
	DBConn.Mark("[bsdo:main()] Doing a Mark() in the mark table.")
	log.Printf("Completed.\n")
}
