package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func stopIfError(err error, when string) {
	if err != nil {
		print(when + "\n")
		panic(err.Error())
	}
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err := db.Exec(query)

	stopIfError(err, "while creating table")
}

func insertUser(db *sql.DB) {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	stopIfError(err, "while inserting user")
}

func queryUserById(db *sql.DB, queryId int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	query := `SELECT id, username, password, created_at FROM users WHERE id = 4`

	err := db.QueryRow(query).Scan(&id, &username, &password, &createdAt)

	stopIfError(err, "while querying by id")

	fmt.Println(id, " "+username+" "+password)
}

func queryAllUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
	stopIfError(err, "while quering all rows")
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
		stopIfError(err, "while scanning single row from queried data")

		fmt.Println(u)
	}

}

func deleteUser(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 6) // check err
    stopIfError(err,"while deleting")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	db, err := sql.Open("mysql", "root:password@(127.0.0.1:23306)/laravel?parseTime=true")
    
	stopIfError(err, "while connecting")
	stopIfError(db.Ping(), "while pinging")

	createTable(db)

	insertUser(db)

	// queryUserById(db, 1)
	queryAllUsers(db)

	// deleteUser(db)

	http.ListenAndServe(":8080", nil)
}
