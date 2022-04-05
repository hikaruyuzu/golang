package godatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

// prepare statement akan membuat koneksi sendiri secara manual
// dimana sebenarnay di atas exect dan query ada prepare statement, akan tetapi jika dibuat secara default dari exect atau query
// dia akan selalau menyakan konseksinya ke db poo;, ini akan sangat buruk jika kita ingin memasuka data dalam jumlah banyak
// akan tetapi parameternya(isi) saja yang berbeda, disini agar dia tidak menanyakan koneksinya berulang kali dan hanya bind ke
// 1 koneksi saja kamu bisa menggunakan prepare statemenet

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "INSERT INTO coments(email, comment) VALUES (?, ?)"

	statement, err := db.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	// always close statement
	defer statement.Close()

	// prepae statement digunakan untuk bind ke satu koneksi saja dan ini sangat cocok diunakan juka kita ingin mrmasukan banyak data
	// dengan query yang sama akan tetaimemiliki data parameter yang berbeda, misal data ratusan baris excel
	// drngasn 1 koneski saja, agar lebih cepat dan tidak perlu menanyakan koneksinya terus terusan ke db pool
	for i := 0; i < 10; i++ {
		email := "kaguyasan" + strconv.Itoa(i) + "@gmail.com"
		comments := "coments from " + email
		// disini kamu bisa menggunakan query apapun derngan statement, ga usah masukin querynya lagi karena sudah di bind sama PrepareStatement
		result, err := statement.ExecContext(ctx, email, comments)
		if err != nil {
			panic(err)
		}

		// get last id
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert data with id ", id)

	}
}
