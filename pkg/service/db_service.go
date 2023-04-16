package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.1.17"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

var dbConn Database

type Database struct {
	*sql.DB
}

func init() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password='%s'"+
		"dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error while connecting to db: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Db connected")

	dbConn.DB = db

}

func GetDbInstance() Database {
	return dbConn
}

func (db Database) TransactionalStep() {

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Println("Could not start transaction! ", err)
		return
	}

	insertDynStmt := `insert into "stu"("name", "roll") values($1,$2)`
	_, err = tx.Exec(insertDynStmt, "Jane", 1)

	if err != nil {
		log.Println("Error while inserting ", err)
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error while commiting transaction")
		tx.Rollback()
		return
	}

}

func (db Database) GetNextRuntime() int64 {

	fetchExecutionTimeStmt := `select timestamp from waitforit`

	var execTimeStamp int64

	err := db.QueryRow(fetchExecutionTimeStmt).Scan(&execTimeStamp)
	if err != nil {
		log.Fatal("Could not get next execution time ", err)
	}

	return execTimeStamp

}
