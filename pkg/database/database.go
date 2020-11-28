package database

import (
	"database/sql"
	"fmt"

	"github.com/nouzun/l2r/pkg/model"

	_ "github.com/go-sql-driver/mysql"
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

func (db *Database) GetWords() ([]model.Word, error) {
	// Execute the query
	results, err := db.database.Query("SELECT id, word, gender, plural, word_type, sentence, word_case FROM words")
	if err != nil {
		db.database.Close()
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var words []model.Word

	for results.Next() {
		var word model.Word

		// for each row, scan the result into our tag composite object
		err = results.Scan(&word.ID, &word.Word, &word.Gender, &word.Plural, &word.Type, &word.Sentence, &word.Case)
		if err != nil {
			db.database.Close()
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		words = append(words, word)
	}

	return words, nil
}
