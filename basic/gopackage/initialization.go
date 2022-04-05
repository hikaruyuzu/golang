package gopackage

// package init atau initialization adalah sebuah package dimana, ketika function di definisikan sebagai func init,
// dia otomatis akank di panggil secara langsung di file pemanggilnya
// ini sangat cocok untuk open connection ke database
// kamu bisa membuatnnya dengan cara membuat nama function dengan nama init

// contoh
var connection string

// func init, otomatis akan dipanggil ketika package di import
func init() {
	connection = "Kubernetes"
}

func GetConnection() string {
	return connection
}

//cara pakai

/**
/ dia akan otomatis dipanggil atau di eksekusi ketika
"basic/gopackage" di import

ignoree error jika package tidak dipakai dengan _
_"basic/gopackage"
*/
