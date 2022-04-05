package godatabase

import (
	"context"
	"log"
	"testing"
)

func TestExectContextSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// jika ingin mengubahnya menjadi insert dekete dan update tinggal ubah querynya saja.
	queryInsert := "INSERT INTO costumers(id, name) VALUES ('AA122', 'Chino Kaffu')"
	_, err := db.ExecContext(ctx, queryInsert)
	if err != nil {
		panic(err)
	}

	log.Println("Success insert data into database")
}

func TestExectUpdate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlUpdate := "UPDATE costumers SET name = 'Kaguya Shinomiya Chan' WHERE id = 'AA123'"

	_, err := db.ExecContext(ctx, sqlUpdate)
	if err != nil {
		panic(err)
	}

	log.Println("Success upgrade data")

}

// kamu bisa menggunakan hal yang sama untuk delete data
