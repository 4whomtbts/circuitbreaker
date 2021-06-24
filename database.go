package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

type BrakedHosts struct {
	hostname string `db:"hostname"`
	braked_at time.Time `db:"braked_at"`
}

type CBLog struct {
	hostname string `db:"hostname"`
	braked bool `db:"braked"`
	created_at time.Time `db:"created_at"`
}

func InitDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Can't connect to sqlite3 : %s", err.Error())
		panic(err)
	}

	drop1 := `DROP TABLE IF EXISTS BRAKED_HOSTS`
	drop2 := `DROP TABLE IF EXISTS CB_LOG`
	if _, err = db.Exec(drop1); err != nil {
		log.Fatalf("Failed to create sqlite3 table : %s", err.Error())
		panic(err)
	}

	if _, err = db.Exec(drop2); err != nil {
		log.Fatalf("Failed to create sqlite3 table : %s", err.Error())
		panic(err)
	}

	createBrakedHostsTable := `
		CREATE TABLE IF NOT EXISTS BRAKED_HOSTS (
			hostname string not null unique,
			brake_type string not null,
			braked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
		)
	`
	createCircuitBrakerLog := `
		CREATE TABLE IF NOT EXISTS CB_LOG (
			hostname string not null,
			brake_Type string not null,
			braked TINYINT not null,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
		)
	`
	if _, err = db.Exec(createBrakedHostsTable); err != nil {
		log.Fatalf("Failed to create sqlite3 table : %s", err.Error())
		panic(err)
	}

	if _, err = db.Exec(createCircuitBrakerLog); err != nil {
		log.Fatalf("Failed to create sqlite3 table : %s", err.Error())
		panic(err)
	}

	tHost := BrakedHosts{
		hostname: "localhost",
		braked_at: time.Now(),
	}

	if _, err := db.Exec("INSERT INTO BRAKED_HOSTS (hostname) VALUES (?)", tHost.hostname); err != nil {
		panic(err)
	}

	if _, err := db.Exec("DELETE FROM BRAKED_HOSTS WHERE hostname=?", tHost.hostname); err != nil {
		panic(err)
	}
	var rst BrakedHosts
	rows, err := db.Query("SELECT * FROM BRAKED_HOSTS")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&rst.hostname, &rst.braked_at)
		if err != nil {
			panic(err)
		}
		fmt.Println(rst.braked_at)
	}


	return db, nil
}