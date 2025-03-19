package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Dbclient *sql.DB

func InitDbConnection() {
	connStr := "postgresql://admin:linhporo1@localhost:5432/admin_movie?sslmode=disable"
	//drive and uri connnection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error at the db repository", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping to database: ", err)
		panic(err)
	} else {
		fmt.Println("Database connected")
	}

	Dbclient = db
}

func CloseDbConnection() {
	err := Dbclient.Close()
	if err != nil {
		fmt.Println("Error closing database: ", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(Dbclient)
}
