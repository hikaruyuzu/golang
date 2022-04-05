package godatabase

import (
	"context"
	"log"
	"testing"
)

// untuk mendapatkan last insert id ketika execContext

func TestLastIdAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "Rankokanzakichan@gmail.com"
	coments := "Wakaranai desu"

	ctx := context.Background()

	sql := "INSERT INTO coments(email, comment) VALUES (?, ?)"
	results, err := db.ExecContext(ctx, sql, email, coments)
	if err != nil {
		panic(err)
	}

	// get last insert id, jika data auto incerement, jasi tidak perlu query 2 kali
	resultId, err := results.LastInsertId()
	if err != nil {
		panic(err)
	}

	log.Println("Success insert data with id ", resultId)
}
