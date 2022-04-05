package helper

import "fmt"

// function as parameter artinya kita bisa memoassing sebuah function sebagai paramater untuk fnction lainnya
// hal ini snagatlahpoerfull
// jika dirasa function ang di inputkan terlalu panjangb kamu bisa menggunakan alias

// kamu bisa membuat alias untuk function dan dimasukan kedalam parameter
type FilterVTuber func(string) string

// contoh
func FilterBannedVtuber(name string, filters func(string) string) string {
	myvtuber := filters(name)

	if len(myvtuber) != 20 {
		return fmt.Sprintf("Hello %s", myvtuber)
	} else {
		return myvtuber
	}
}

// ini adalah function untuk filternya
func BannedVtuber(name string) string {
	var results string
	var vtuber = []string{"Pekora", "Gura", "Ollie", "Mona", "Fuubuki"}
	for i := 0; i < len(vtuber); i++ {
		if vtuber[i] == name {
			results = "VTuber banned, karena terlalu op"
		} else {
			results = name
		}
	}

	return results
}

// dalam pemanggilan nantinya kamu bisa passing BannedVtuber function kedalam FilterBanned vtuber, baik menggunakan type declaration(alias)
// atau secara langsung

// contoh pemanggilan

/*
	var YourName = helper.FilterBannedVtuber("Gura", helper.BannedVtuber)
	fmt.Println(YourName)
*/
