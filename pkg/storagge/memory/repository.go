package database

import (
	"database/sql"
	//A MySQL-Driver for Go's database/sql package
	_ "github.com/go-sql-driver/mysql"
)

// TestConection is 
func TestConection() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:ffff000000ff@/BattleHelper")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO BattleHelper VALUES ( 10, 'hyunji' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}