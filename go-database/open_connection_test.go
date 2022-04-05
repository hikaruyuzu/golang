package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:neko@tcp(localhost:3306)/learn_golang_database")
	if err != nil {
		panic(err)
	}
	// always close
	defer db.Close()

	// use DB

}
