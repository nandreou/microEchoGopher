package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var db = &DB{}

const (
	dbData = "host=192.168.1.17 port=5432 dbname=mydb user=postgres password=password"
)

func ConnectToDB() (*DB, error) {
	var err error

	db.SQL, err = sql.Open("pgx", dbData)

	if err != nil {
		return nil, err
	}

	err = db.SQL.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) ApiKeyValidate(apikey string) (bool, error) {
	query := `SELECT * FROM apikeys WHERE apikey = $1`

	sqlRow := &struct {
		ApiKey  string
		Email   string
		Created time.Time
	}{}

	row := db.SQL.QueryRow(query, apikey)

	err := row.Scan(
		&sqlRow.ApiKey,
		&sqlRow.Email,
		&sqlRow.Created,
	)

	if err != nil {
		return false, err
	}

	//ADD Email Check too Here

	if err != nil {
		fmt.Println("Error parsing created timestamp:", err)
		return false, err
	}

	duration := 24 * time.Hour

	if sqlRow.Created.Sub(time.Now()) > duration {
		return false, nil
	}

	return true, nil

}

func (db *DB) WriteApiKeyToDB(apikey, email string) (sql.Result, error) {
	query := `INSERT INTO apikeys (apikey, email, created_at) 
	VALUES ($2, $1, $3);`

	return db.SQL.Exec(query,
		apikey,
		email,
		time.Now(),
	)
}
