package db

// DB Helper
import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Successfully set up a new posgres db connection with the conn string and ping tested it, so return
	return db, nil
}
