package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID      string
	Name    string
	Phone   string
	Created time.Time
}

func DB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalln(err)
	} else if db == nil {
		log.Fatalln("Could not connect to the database.")
	}
	return db
}

func createTable(db *sql.DB) {
	q := `
		CREATE TABLE IF NOT EXISTS users(
			ID TEXT NOT NULL PRIMARY KEY,
			Name TEXT,
			Phone TEXT,
			Created DATETIME
		);
	`
	_, err := db.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

func insertUsers(db *sql.DB, users []User) {
	q := `
		INSERT OR REPLACE INTO users(
			ID, Name, Phone, Created
		) 
		values(
			?, ?, ?, CURRENT_TIMESTAMP
		)
	`

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	for _, user := range users {
		_, err = stmt.Exec(user.ID, user.Name, user.Phone)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getUsers(db *sql.DB) []User {
	q := `
		SELECT
			ID, Name, Phone, Created
		FROM
			users
		ORDER BY datetime(Created) DESC
	`

	rows, err := db.Query(q)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Created)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	return users
}

func main() {
	const path = "users.db"

	db := DB(path)
	defer db.Close()

	createTable(db)

	users := []User{
		User{"1", "John Doe", "555 123 1234", time.Now()},
		User{"2", "Jane Doe", "555 321 4321", time.Now()},
	}
	insertUsers(db, users)

	u := getUsers(db)
	log.Printf("%#v\n", u)

	moreUsers := []User{
		User{"1", "Bob Doe", "555 987 6543", time.Now()},
		User{"3", "Mary Lamb", "555 789 3456", time.Now()},
	}
	insertUsers(db, moreUsers)

	u = getUsers(db)
	log.Printf("%#v\n", u)
}
