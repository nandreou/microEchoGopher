package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.io/nicksauth/models"
)

type DB struct {
	SQL *sql.DB
}

var db = &DB{}

const (
	dbData = "host=192.168.1.17 port=5432 dbname=mydb user=postgres password=password"
)

func ConnectToDb() (*DB, error) {

	var err error

	db.SQL, err = sql.Open("pgx", dbData)

	if err != nil {
		return nil, err
	}

	err = db.SQL.Ping()

	if err != nil {
		return nil, err
	} else {
		log.Println("Connected To Database !!!")
	}

	return db, nil
}

func (db *DB) Auth(email string, password string) (bool, error) {

	query := "SELECT * FROM users WHERE email = $1 AND pass = $2"

	var users []*models.User

	rows, err := db.SQL.Query(query, email, password)

	if err != nil {
		log.Println("Error:", err)
	}

	defer rows.Close()

	for rows.Next() {
		var tmp models.User

		if err = rows.Scan(
			&tmp.ID,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.Email,
			&tmp.Username,
			&tmp.Password,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		); err != nil {
			return false, err
		} else {
			users = append(users, &tmp)
		}

	}

	return users != nil, nil
}
