package dbsrv

import (
	"log"
        "database/sql"
	"os"

        _ "github.com/go-sql-driver/mysql"
        _ "github.com/mattn/go-sqlite3"
)

func init() {
	log.Printf("[dbsrv:init()] - Starting.\n")
}

//Service  - odot:(hopley) something more applicable ? less generic
type DBService struct {
	Driver  string
        DB   *sql.DB
        Path string
}

// Build will completely build the database. Not a good thing to run unless you really need to.
//TODO:(hopley) Some testing for if no structure & define db schema, evolution 'pattern'
func (s *DBService) Build() error {
        return nil
}

// Open initializes the database session and gives you back an error if failed.
func (s *DBService) Open() error { 
	// db, err := sql.Open("mysql", "user:password@/dbname")
	// username:password@protocol(address)/dbname?param=value  
	//  where address is For TCP and UDP networks, addresses have the form host:port
	//mysql
        //db, err := sql.Open("mysql", s.Path)
	//sqlite -        DBC, err = sql.Open("sqlite3", dsn)
	//db, err := sql.Open("sqlite3",s.Path)
	db,err := sql.Open(s.Driver,s.Path)
        if err != nil {
		log.Printf("[dbsrv:Open()] ERROR: <TEMP> STOP here (dbsrv.go:31)\n")
		os.Exit(128)
                //return err
        }
	//dh 
	//defer db.Close()

        s.DB = db
        return nil
}

//New() Return a *DBService type to caller (includes the database pointer).
func New(d,p string) *DBService {
	//TODO:(hopley) upd so driver and dsn are passed
        s := &DBService{Driver: d,Path: p}
        return s
}

//
//mark() - func to put a message in the database ; Globa
func (s *DBService) Mark(msg string) {
        //TODO:(hopley) - For the Testing put a 'WT' mark in and then query for it
        result, err := s.DB.Exec("INSERT into mark(MarkLabel) VALUES('TestMark')")
        if err != nil {
		log.Printf("[dbsrv.Mark()] - ERROR - doing INSERT for mark.(err=%v)\n",err)
                os.Exit(128)
        }
        id, err := result.RowsAffected()
        if id > 1 {
                log.Printf("[db.makeMark()] - NOTE - Check , more than (1) row inserted?  In dbase.makeMark() id result.RowsAffected ()=%d\n", id)
        }
}

/*
 create the mark table
 # Drop TABLE mark;
CREATE TABLE mark (
  MarkLabel VARCHAR(140),
  UpdatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
#
*/


