package godatabase

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// digunakan khusus untuk melakukan query yang mengembalikan data dari databse seperti select

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	querySelect := "SELECT id,name FROM costumers"

	rows, err := db.QueryContext(ctx, querySelect)
	if err != nil {
		panic(err)
	}

	// iterasi data dari rows, membaca semua data
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID ", id)
		fmt.Println("Name ", name)
	}

	log.Println("Success get data from databse")
}
