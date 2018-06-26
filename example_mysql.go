package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
)

type (
	Context struct {
		ID          int
		TypeSearch  string
		TypeContext string
		Request     string
		Directlink  string
		Filename    string
		Result      bool
	}
)

var (
	createTable = "CREATE TABLE IF NOT EXISTS context(" +
		"id INTEGER PRIMARY KEY," +
		"typesearch TEXT," +
		"typecontext TEXT," +
		"request TEXT," +
		"directlink TEXT," +
		"filename TEXT," +
		"result BOOL"

	insertRecord = "INSERT INTO context (" +
		"typesearch, typecontext,request,directlink,filename,result)" +
		" values (?, ?, ?, ?, ?,?)"
)

func main() {
	//open/create database sqlite3
	dbs, err := sql.Open("sqlite3", "dbssqlite.dbs")
	if err != nil {
		log.Panic(err)
	}
	defer dbs.Close()
	log.Printf("DBS instance: %v\n", dbs)
	//create if not exists new table
	stm, _ := dbs.Prepare(createTable)
	//creating
	result, _ := stm.Exec()
	//insert new record
	stm, _ = dbs.Prepare(insertRecord)
	for x:=0; x < 10; x ++ {
		stm.Exec("google","image","http://somereguestgoole","somefilename.jpg","0")
	}


	log.Printf("resultexec: %v\n", result)
}
