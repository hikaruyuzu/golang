package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
)

// insert ke data sellers
func TestInsertDataSellers(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlInsert := "INSERT INTO sellers(id, name, email, birth_date, balance, rating, married) VALUES ('A0003', 'Kato Megumi', 'megimichan@gmail.com', '2003-08-10', 80000000, 9.8, true)"

	_, err := db.ExecContext(ctx, sqlInsert)
	if err != nil {
		panic(err)
	}

	log.Println("Success insert data into table sellers")
}

// untuk query complex
func TestQueryComplexSellers(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name, email, birth_date, balance, rating, married, created_at FROM sellers"
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	// iterate data
	var id, name, email string
	var birth_date, created_at time.Time
	var balance int32
	var rating float64
	var married bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &birth_date, &balance, &rating, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("id ", id)
		fmt.Println("name ", name)
		fmt.Println("email ", email)
		fmt.Println("birthday ", birth_date)
		fmt.Println("balance ", balance)
		fmt.Println("rating ", rating)
		fmt.Println("married ", married)
		fmt.Println("created at ", created_at)
		fmt.Println("========================")
	}

	log.Println("Success get data from table sellers")

}

// jika data yang di kembalikan query bisa null misal
func TestQueryComplexIfDataNullable(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name, email, birth_date, balance, rating, married, created_at FROM sellers"
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	// iterate data
	var id, name string
	var email sql.NullString // jika data di table databse bisa null kamu bisa menggunakan tipe data nullable
	var created_at time.Time
	var birth_date sql.NullTime
	var balance int32
	var rating float64
	var married bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &birth_date, &balance, &rating, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("id ", id)
		fmt.Println("name ", name)
		if email.Valid {
			fmt.Println("email ", email.String)
		}
		if birth_date.Valid {
			fmt.Println("birthday ", birth_date.Time)
		}
		fmt.Println("balance ", balance)
		fmt.Println("rating ", rating)
		fmt.Println("married ", married)
		fmt.Println("created at ", created_at)
		fmt.Println("========================")
	}

	log.Println("Success get data from table sellers")

}
