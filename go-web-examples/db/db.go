package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DB *sql.DB

func init() {
    var err error
    DB, err = sql.Open("mysql", "myuser:mypassword@tcp(127.0.0.1:3306)/mydb?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Database connected!")
}
