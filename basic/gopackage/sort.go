package gopackage

import "sort"

// digunakan untuk melakukan pengurutan
// disini untuk melakukan pengurutan kamu harus mengikuti kontarak dari package sort
// lihat dokumentasi lengkapnya di https://pkg.go.dev/sort

// misal kita punya struct Anime dan kita akan mengurutkannya berdasrkan rating
type AnimeList struct {
	Title  string
	Rating float32
}

// alias untuk []AnimeList
type ArrayAnimeList []AnimeList

// membuat kontrak package sort
func (anime ArrayAnimeList) Len() int {
	return len(anime)
}

// compare
func (anime ArrayAnimeList) Less(i, j int) bool {
	return anime[i].Rating < anime[j].Rating
}

// swap
func (anime ArrayAnimeList) Swap(i, j int) {
	anime[i], anime[j] = anime[j], anime[i]
}

// call di amin function
func SortingAnimeByRating(anime ArrayAnimeList) {
	sort.Sort(anime)
}

/**
/ cara penggunaan

anime := []gopackage.AnimeList{
	{"Sword art online", 7.8},
	{"Attack on titan", 8.8},
	{"Non non Biyori", 8.3},
	{"Kaguya sama love is war", 8.5},
}

// sebelum di sort
fmt.Println(anime)

gopackage.SortingAnimeByRating(anime)
// konversi tipe data
fmt.Println(gopackage.ArrayAnimeList(anime))

*/
