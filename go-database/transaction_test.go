package godatabase

import (
	"context"
	"log"
	"strconv"
	"testing"
)

// ini adalah representasi dari transaction di mysql
// kamu bisa menggunakan db.Begin() untuk memulai dimana disini dia mengembalikan tx, dan err
// setelah melakukan transaction kamu bisa memngguakan commit/ rolleback

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// start tarnsaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// query
	sql := "INSERT INTO coments(email, comment) VALUES (?, ?)"
	// prepare statement lock 1 connection
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "shiyzu" + strconv.Itoa(i) + "@gmail.com"
		comments := "comments form email " + email

		// start insert data
		result, err := stmt.ExecContext(ctx, email, comments)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		// console log
		log.Println("Success insert data into table comments with id ", id)
	}

	// bisa kamu commit/ rolleback
	err = tx.Rollback()
	if err != nil {
		defer panic(err)
	}

}
