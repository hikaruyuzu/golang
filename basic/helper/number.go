package helper

// tipe data number di golang dibagi menjadi 2 yaitu int dan float
/*
integer dibagi menjadi int8, int16 , int 32, dan int64
float dibagi menjadi float32 dan float 64 , dan juga complex32 dan complex64, complex digunakan untuk operasi yang matematik sekali
adapun tipe data tanpa nilai minus - yaitu uint
default nilai jika kita tulis int atau float saja dia akan mengikuti berapa bit sistem operasi kita
*/

// function convert rupiah ke eth
func RupiahToEtherium(idr float64) float64 {
	eth := idr / 34087229
	// return nilai eth
	return float64(eth)
}

// function convert dolar ke rupiah
func DollarToRupiah(dollar float32) int64 {
	rupiah := (dollar * 14387.40)
	return int64(rupiah)
}
