package hippo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDB(config *Config) *sql.DB {
	db := config.Database
	initStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", db.Name, db.User)
	database, err := sql.Open("postgres", initStr)
	if err != nil {
		panic(err)
	}
	return database
}
