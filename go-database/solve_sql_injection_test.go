package godatabase

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// ketika melakukan insert atau query dengan cara sebelumya kita selalu menggunakan hardcode input, itu sangat berbahaya
// disitu user akan bisa melakukan manipulasi input yang dapat menyebabkan sql injection
func TestInsertSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "Ranko'; DROP TABLE users'; #" // nah ini akan aman tidak akan terjadi sql injection
	password := "ayangranko"

	ctx := context.Background()

	sql := "INSERT INTO users(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, sql, username, password)
	if err != nil {
		panic(err)
	}

	log.Println("Success insert user ", username)

}

// sama saja untu select datanya kamu bisa memanfatkan ? / sql parameter untuk menghindari sql injection
func TestQuerySafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "Kaguya Shinomiya"
	password := "Salah"

	ctx := context.Background()

	sql := "SELECT username, password FROM users WHERE username= ? AND password= ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sql, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)

		} else {
			fmt.Println("Success login, welcome ", username)
		}

	} else {
		fmt.Println("Gagal login, username atau password anda salah")
	}
}
