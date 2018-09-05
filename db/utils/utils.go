package utils

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gchaincl/dotsql"
)

func OpenMySQL(uname, pass string) *sql.DB {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)" +
                        "/romanwing?parseTime=true", uname, pass))
    if err != nil {
        panic(err.Error())
    }
    return db
}

func Migrate(uname, pass, mFName string, tNames []string) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/",
                        uname, pass))
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    db.Exec("DROP DATABASE romanwing")
    db.Exec("CREATE DATABASE romanwing")
    db.Exec("USE romanwing")
    dot, err := dotsql.LoadFromFile(mFName)
    if err != nil {
        panic(err.Error())
    }
    for _, tName := range tNames {
        fmt.Println("Migrating table " + tName + "...")
        _, err := dot.Exec(db, tName)
        if err != nil {
            fmt.Printf("Failed to migrate table " + tName + ": ")
            fmt.Println(err)
        } else {
            fmt.Println("Successfully migrated " + tName + ".")
        }
    }
}
