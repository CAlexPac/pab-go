package model

import ("fmt"
        "database/sql"
        _ "github.com/lib/pq"
)

const (
    DB_HOST = "localhost"
    DB_PORT = "5432"
    DB_USER = "pab"
    DB_PASS = "pab"
    DB_NAME = "pab"
    DB_SSL_MODE = "disable"
)

var dbconn *sql.DB


func dbconnect() *sql.DB {
    url  := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME, DB_SSL_MODE)
    conn, err := sql.Open("postgres", url)

    if err != nil {
        panic(err) 
    }

    return conn
}

func getdbconn() *sql.DB {
    if dbconn == nil {
        dbconn = dbconnect()
    }
    return dbconn
}
