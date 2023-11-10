package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	panicOnError(err)

	_, err = os.Create(os.Getenv("DATABASE_NAME"))
	panicOnError(err)

	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_NAME"))
	panicOnError(err)

	err = db.Ping()
	panicOnError(err)

	defer db.Close()

	// Drop existing tables
	_, err = db.Exec("DROP TABLE IF EXISTS focuspoints_on_exercise")
	panicOnError(err)

	_, err = db.Exec("DROP TABLE IF EXISTS exercises")
	panicOnError(err)

	_, err = db.Exec("DROP TABLE IF EXISTS focuspoints")
	panicOnError(err)

	// Create new tables
	_, err = db.Exec(`CREATE TABLE "exercises" (
		"id" INTEGER,
		"description" TEXT,
		"duration" INTEGER,
		"section" TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)
	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
