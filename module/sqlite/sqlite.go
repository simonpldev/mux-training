package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect() {

	var err error

	// Get connection to database
	db, err = sql.Open("sqlite3", "./data/todos.db")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare todo table
	_, err = db.Exec("create table if not exists todos (id integer primary key autoincrement, title nvarchar(50), status bit)")
	if err != nil {
		log.Fatal(err)
	}

}

func GetDatabase() *sql.DB {
	return db
}
