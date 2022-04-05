package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// channnel digunakan untuk melakukan komunikasi antar goroutines
// digunakan untuk mengirim dan menerima data secara syncronus, yang artinya blocking dalam proeses pengiriman datanya
// jika kita menirim data dari groutines kedalam channel maka kita harus menunggu sampai datanya di ambil
// dalam channel kita harus punya sender dan receiver , jika salah satunya tidak ada maka akan terjadi error
// channel digunakan untuk menerima data yang memiliki return value dari goroutines dan menerimanya

// jika ada receiver tapi tidak ada sender yang mengirim data ke channel maka akan terjadi deadlock! begitupun sebaliknya

func TestCreateChannel(t *testing.T) {
	// gunakan make(chan dataType) untuk membuat channel
	channel := make(chan string)
	defer close(channel) // selalu close channel, agar tidak terkena memory leak(menggantung)

	// channel <- "something" -> mengirim data ke channel
	// data := <- channel -> menerima data dari channel
	// fmt.Pritln(<-channel) -> sperti ini juga bisa, digunakan untuk menerima data

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Ara ara ~ Kawaii senpaii"
		fmt.Println("Success send data to channel")
	}()

	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

// channel as parameter
// secara default channel di parameter akan pas by reverence, tidak butuh pointer
func GiveMeResponse(channelParameter chan string) {
	time.Sleep(2 * time.Second)
	channelParameter <- "Ara ara ~ Kawaii"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// call function, parameter
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

// Only in & Only Out
// Secara default channel di parameter bisa menerima dan mengirim data ke channel
// jika ingin hanya bisa melakukan salah satu saja bisa menggunakan channel OnlyIn dan OnlyOut
// Only In dapat dibuat dengan channel chan <- string (hanya dapat menirim data)
// ONly Out dapat dibuat dengan channel <- chan string (hanya dapat menerima data)

// implementasii
func OnlyIn(channelParameter chan<- string) {
	time.Sleep(2 * time.Second)
	// data := <- channelParameter -> akan error
	channelParameter <- "Ara ara ~ Kawaii"
	fmt.Println("Success send data")
}

func OnlyOut(channelParameter <-chan string) {
	// channelParameter <- "Ara ara ~ Kawaii" -> akan terjadi error
	data := <-channelParameter
	fmt.Println(data)
}

func TestOnlyInAndOut(t *testing.T) {
	channel := make(chan string)

	// call function
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

// buffered channel
// ini sangat cocok digunakan ketika receiver lebih lambat daripda sender
// secara default channel hanya memiliki satu buffer, hal ini akan membuat sender menjadi lambat jika receiver lambat mengambil data
// dan juga akan menyebabkan error deadlock jika data dalam channel tidak di ambil
// kamu bsa menggunakan buffer untuk mengatasi masalah ini
// buffer adalah penampungan data sementara dalam channel

// contoh
func TestCreateBufferedChannel(t *testing.T) {
	// kamu bisa memasukan kapasitas bufer(panampung) setelah deklarasi chan, misal 3
	channel := make(chan string, 2)
	// nah jadi misal walau tidak ada receiver/ receivernya lambat data akan dikirim, dan tidak akan terjadi deadlock
	// akan tetapi jika data yang ada di buffer sudah lebih dari kapasitas buffer dia akan enunggu sampai datanya di ambil
	go func() {
		channel <- "Kaguya shinomiya"
		channel <- "Ranko kanzaki"
		channel <- "Kato Megumi"

	}()

	// fmt.Println(len(channel)) // melihat panjang buffer yang terisi

	go func() {
		fmt.Println(<-channel) // mengambil data dari channel buffer
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(cap(channel)) // melihat panjag dari buffer
}

// for range
// kadang kita tidak tahu ada berapa data yang harus kita kirim dan terimaka di dalam channel(data sebanyak banyaknya)
// nah jika msalahnya seperti ini kita bisa menggunakan for range untuk mengatasinya
// kamu bisa menggunakan perulangan ini dalam penerimaan data
// dengan syarat waktu pengiriman data selsesai harus meng close channelnya sat itu juga agar tidak dilakukan perulangan pengambilan data secara etrus menerus (deadlock)

func TestCreateForRange(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- "Data ke " + strconv.Itoa(i)
		}
		// harus di close ketika selesai agar tidak terkena deadlock
		close(channel)

	}()

	go func() {
		// kita bisa terima dengan for range, jadi jangan manual
		for data := range channel {
			fmt.Println(data)
		}
		fmt.Println("Selesai")
	}()

	time.Sleep(2 * time.Second)

}

// kadang kita butuh mengambil bebrapa data dari bebrapa channel seklaigus
// jika kasusnya sperti ini kamu bisa menggunakan select channel
// dan ini juga dapat dikombinasikan dengan perulangan for agar kita tidak capek membuat se;ect ebbrapa kali
// kita kita menggunakna selct maka data yang paling cepat diterima yang akan di eksekusi terlebih dahulu

// contoh
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)
	//give response func sudah dibuat di atas
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// nah untuk menerima dua data dari channel ini kita bisa menggunakan select
	// disini daia akn di eksekusi yang lebih cepat terlebih dahulu
	/**
	/ cara seperti ini kurang efektip, akan bikin cape, jadi kita bisa mengguanakn perulangan for dengan kondisi berhenti
	/ agar tidak terjadi deadlock

	select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
	}

	select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
	}

	*/
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		}
		// cek kondisi counter jika data sudah tereksekusi semua maka break
		if counter == 2 {
			break
		}

	}

}

// default select
// jika ingin melakukan sesuatu jika data belum tersedia kamu bisa menggunakan default select

func TestDeafaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)
	//give response func sudah dibuat di atas
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// nah untuk menerima dua data dari channel ini kita bisa menggunakan select
	// disini daia akn di eksekusi yang lebih cepat terlebih dahulu
	/**
	/ cara seperti ini kurang efektip, akan bikin cape, jadi kita bisa mengguanakn perulangan for dengan kondisi berhenti
	/ agar tidak terjadi deadlock

	select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
	}

	select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
	}

	*/
	counter := 0
	test := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++

			// kamu bisa menggunakakn default di select
		default:
			test = test + 1
			fmt.Println("Chotomatte Onichan, your data is being proccesed > ~ <", test)
		}
		// cek kondisi counter jika data sudah tereksekusi semua maka break
		if counter == 2 {
			break
		}

	}

}
