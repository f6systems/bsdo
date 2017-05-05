package main

import (
//	"database/sql"
	"log"

	"github.com/f6systems/bsdo/dbsrv"
)

//Confi the structure for the application

func init() {
	log.Printf("[init()] Starting.\n")
}

func main() {
	log.Printf("Starting.\n")
	// move to mysql for testing
	/*
	dbconn := dbsrv.New( //"root:root@/weave?parseTime=true",
		"root:root@/ww?parseTime=true"
	)
	*/
	// db, err := sql.Open("mysql", "user:password@/dbname")
        // username:password@protocol(address)/dbname?param=value
        //  where address is For TCP and UDP networks, addresses have the form host:port
	//DSN format.  db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	dbconn := dbsrv.New("root:wiq@tcp(127.0.0.1:3306)/ww")
	//shareDSN := dsn + "?_busy_timeout=50000"
	// see = go/src/test
	//createDatabaseFile(dsn)
	/*
	dbc, err := sql.Open("sqlite3", shareDSN)
	if err = dbc.Ping(); err != nil {
		log.Panic("ERROR: Ping of sqlite database failed. ", err)
	}
	if err != nil {
		log.Printf("ERROR - opening database=%s\n", dsn)
	}
	*/
	//need an open
	err := dbconn.Open() //dbsrv.Open()
	if err != nil {
		log.Printf("ERROR: Issue Opening the database. (err=%v)\n",err)
	}

	dbconn.Mark("[bsdo:main()] Doing a Mark() in the mark table.")
	log.Printf("Completed.\n")
}
