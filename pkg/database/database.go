package db

import (
	"database/sql"
	"fmt"

	"github.com/nouzun/l2r/model"
)

type Database struct {
	database *sql.DB
}

func ConnectDatabase() (*Database, error) {
	serverName := "localhost:3306"
	user := "root"
	password := "456852"
	dbName := "learn2remember"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return &Database{database: db}, nil
}

func (db *Database) GetWords() error {
	// Execute the query
	results, err := db.database.Query("SELECT id, word, gender, plural, type, sentence, case FROM words")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var word model.Word

		// for each row, scan the result into our tag composite object
		err = results.Scan(&word.ID, &word.Word, &word.Gender)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return nil
}
