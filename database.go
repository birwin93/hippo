package hippo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Transaction interface {
	Query() string
	QueryVals() []interface{}
	Scan(row *sql.Row) error
}

type Database struct {
	DB *sql.DB
}

func NewDB(c DatabaseConfig) (*Database, error) {
	initStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", c.User, c.Name)
	database, err := sql.Open("postgres", initStr)
	if err != nil {
		return nil, err
	}
	return &Database{DB: database}, nil
}

func (db *Database) Select(tx Transaction) error {
	statement, err := db.DB.Prepare(tx.Query())
	if err != nil {
		return err
	}
	defer statement.Close()
	return tx.Scan(statement.QueryRow(tx.QueryVals()...))
}

func (db *Database) Insert(tx Transaction) error {
	statement, err := db.DB.Prepare(tx.Query())
	if err != nil {
		return err
	}
	log.Println(statement)
	defer statement.Close()
	return tx.Scan(statement.QueryRow(tx.QueryVals()...))
}

func (db *Database) Destroy(tx Transaction) error {
	statement, err := db.DB.Prepare(tx.Query())
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(tx.QueryVals()...)
	return err
}
