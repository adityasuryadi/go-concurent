package test

import (
	"fmt"
	"testing"
	"time"
)

// :flow
// - buatkan channel - done
// - buatkan simulasi pengiriman data lewat channel
// - buatkan simulasi penerimaan data lewat channel
// - tunggu sampai program selesai
func TestChannel(t *testing.T) {
	fmt.Println("start")
	messageCH := make(chan string)

	test := make(chan int)
	// receiver
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("menerima pesan")
			messageData := <-messageCH
			fmt.Printf("data di terima: %s\n", messageData)
			fmt.Printf("test var %d\n", test)
		}
	}()

	// time.Sleep(5 * time.Second)

	// sender
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("mengirim pesan")
			messageCH <- "ini ada pesan dari go routine"
			test <- i
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}

// channel default buffer nya 1
// kadang jika si penerima channel lambat maka si penerima juga lambat,jadi bisa di set bufered channel nya
// unbuffered channel digunakan jika saat kita program yang running nya sedikit dan tidak memberatkan ram dan process nya
// buffered channel digunakan utk data yg banyak dan di jalankan saat concurent
// dapat menimbulkan penggunaan ram dan processor yang begitu baynyak
// sehingga kita dapat melimit nya menggunkan buffered channel
func TestBufferedChannel(t *testing.T) {
	fmt.Println("start")

	// set buffered channel jadi 3
	messageCH := make(chan string, 3)

	// receiver
	go func() {
		for {
			messageData := <-messageCH
			fmt.Printf("data di terima: %s\n", messageData)
		}
	}()

	// sender
	go func() {
		for i := 1; i <= 12; i++ {
			fmt.Printf("data ke :%d dikirim\n", i)
			messageCH <- fmt.Sprintf("data dari goroutine ke %d \n", i)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
