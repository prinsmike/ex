package main

import (
	"database/sql"
	"log"
	"net/url"
	"os"
)

func initDB() {
	_ = os.Remove("./users.db")
	w, err := os.Create("./users.db")
	if err != nil {
		log.Fatalln("Could not create the database file.")
	}
	defer w.Close()

	db := DB("./users.db")
	defer db.Close()

	createTables(db)
	insertTestData(db)
}

func createTables(db *sql.DB) {

	// Companies
	q := `
		CREATE TABLE IF NOT EXISTS companies(
			id INTEGER PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			parent INT,
			contact INT,
			domain TEXT,
			published INT NOT NULL
		)
	`

	_, err := db.Exec(q)
	if err != nil {
		log.Fatalf("Could not create companies table. Error: %s\n", err)
	} else {
		log.Println("Created companies table.")
	}

	// Job Positions
	q = `
		CREATE TABLE IF NOT EXISTS job_positions(
			id INTEGER PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			published INT NOT NULL
		);
	`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatalf("Could not create job_positions table. Error: %s\n", err)
	} else {
		log.Println("Created job_positions table.")
	}

	// Users
	q = `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY NOT NULL,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			passhash TEXT NOT NULL,
			role_id INT NOT NULL,
			firstname TEXT,
			lastname TEXT,
			company INT,
			parent INT,
			job_position INT,
			sex CHAR(1) NOT NULL,
			birthdate INT,
			active INT NOT NULL,
			FOREIGN KEY(company) REFERENCES companies(id),
			FOREIGN KEY(parent) REFERENCES users(id),
			FOREIGN KEY(job_position) REFERENCES job_positions(id)
		);
	`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatalf("Could not create users table. Error: %s\n", err)
	} else {
		log.Println("Created users table.")
	}
}

func insertTestData(db *sql.DB) {
	insertTestCompanies(db)
	insertTestJobPositions(db)
	insertTestUsers(db)
}

func insertTestCompanies(db *sql.DB) {
	example_url, err := url.Parse("http://www.example.com/about")
	if err != nil {
		log.Fatalln(err)
	}
	comps := []Company{
		Company{Title: stringPtr("Comp 1"), Parent: &Company{ID: int64Ptr(0)}, Contact: &User{ID: int64Ptr(0)}, Domain: example_url, Published: boolPtr(true)},
		Company{Title: stringPtr("Comp 2"), Parent: &Company{ID: int64Ptr(0)}, Contact: &User{ID: int64Ptr(0)}, Domain: example_url, Published: boolPtr(true)},
		Company{Title: stringPtr("Comp 3"), Parent: &Company{ID: int64Ptr(0)}, Contact: &User{ID: int64Ptr(0)}, Domain: example_url, Published: boolPtr(false)},
	}

	q := `
		INSERT OR REPLACE INTO companies(
			title, parent, contact, domain, published
		)
		VALUES(
			?, ?, ?, ?, ?
		)
	`

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	for _, company := range comps {
		_, err = stmt.Exec(
			company.Title,
			company.Parent.ID,
			company.Domain.String(),
			company.Contact.ID,
			company.Published,
		)
		if err != nil {
			log.Println(err)
		}
	}
}

func insertTestJobPositions(db *sql.DB) {
	jps := []JobPosition{
		JobPosition{Title: stringPtr("JP 1"), Published: boolPtr(true)},
		JobPosition{Title: stringPtr("JP 2"), Published: boolPtr(true)},
		JobPosition{Title: stringPtr("JP 3"), Published: boolPtr(false)},
	}

	q := `
		INSERT OR REPLACE INTO job_positions(
			Title, Published
		)
		VALUES(
			?, ?
		)
	`

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	for _, jp := range jps {
		_, err = stmt.Exec(jp.Title, jp.Published)
		if err != nil {
			log.Println(err)
		}
	}
}

func insertTestUsers(db *sql.DB) {
	users := []User{
		User{
			Username: stringPtr("User 1"), Email: stringPtr("user1@example.com"),
			PassHash: stringPtr("1234567890"), RoleID: intPtr(0),
			Sex: stringPtr("M"), Active: boolPtr(true),
		},
		User{
			Username: stringPtr("User 2"), Email: stringPtr("user2@example.com"),
			PassHash: stringPtr("1234567890"), RoleID: intPtr(0),
			Sex: stringPtr("F"), Active: boolPtr(true),
		},
		User{
			Username: stringPtr("User 3"), Email: stringPtr("user3@example.com"),
			PassHash: stringPtr("1234567890"), RoleID: intPtr(0),
			Sex: stringPtr("M"), Active: boolPtr(false),
		},
	}

	q := `
		INSERT OR REPLACE INTO users(
			username, email, passhash, role_id, sex, active
		)
		VALUES(
			?, ?, ?, ?, ?, ?
		)
	`

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	for _, user := range users {
		_, err = stmt.Exec(user.Username, user.Email, user.PassHash, user.RoleID, user.Sex, user.Active)
		if err != nil {
			log.Println(err)
		}
	}
}
