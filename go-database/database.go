package godatabase

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:neko@tcp(localhost:3306)/learn_golang_database?parseTime=true") // parseTime = true, agar jika query data berupa time tidak error
	if err != nil {
		panic(err)
	}

	// management databse pool
	db.SetMaxIdleConns(10)                  // default miinimum database connction
	db.SetMaxOpenConns(100)                 // maximal connection
	db.SetConnMaxIdleTime(5 * time.Minute)  // jika database tidak digunakan selama 5 menit maka akan di close ke nilai default minimumnya
	db.SetConnMaxLifetime(60 * time.Minute) // refresh connection setiap 60 menit

	return db

}
