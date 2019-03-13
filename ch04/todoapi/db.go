package todoapi

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

func CreateDbMap(dbURL string) (*gorp.DbMap, error) {

	ds, err := createDatasource(dbURL)
	if err != nil {
		return nil, err
	}

	db := &gorp.DbMap{
		Db: ds,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}

	db.AddTableWithName(Todo{}, "todo").SetKeys(true, "ID")
	return db, nil
}

func createDatasource(dbURL string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(2)
	return db, nil
}

type Todo struct {
	ID      uint      `db:"id"      json:"id"`
	Title   string    `db:"title"   json:"title"`
	Content string    `db:"content" json:"content"`
	Status  string    `db:"status"  json:"status"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}
