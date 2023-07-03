package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func connectDB(host string, port int, user string, password string, dbname string) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
func createDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP
	);
`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func insertDB(username string, password string, db *sql.DB) {
	createdAt := time.Now()

	_, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
}
func deleteFromDB(id int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
func printUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)
}
func getWhereIdEquals(id int, db *sql.DB) {
	var (
		username  string
		password  string
		createdAt time.Time
	)

	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}
