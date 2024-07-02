package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWmutex(t *testing.T) {
	// flow:
	// 1. create bookstore
	bookstore := NewBookStore()

	// 2. simulate multiple readers
	for i := 0; i < 10; i++ {
		go func() {
			book := bookstore.getBookDetail(1)
			if book != nil {
				fmt.Printf("reader ke %d buku dengan id %d ditemukan quantity: %d \n", i, 1, book.Quantity)
			} else {
				fmt.Printf("reader ke %d, buku tidak ditemukan \n", 1)
			}
		}()
	}
	// 3. simulate write
	go func() {
		fmt.Println("Writter: Mengupdate jumlah buku..")
		bookstore.updateBookQueantity(1, -1)
		fmt.Println("writter: berhasil mengupdate jumlah buku")
		time.Sleep(3 * time.Second)
	}()
}

type Book struct {
	Title    string
	Quantity int
}

type BookStore struct {
	books   map[int]*Book
	rwMutex sync.RWMutex
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: map[int]*Book{
			1: {
				Title:    "Go Programming",
				Quantity: 10,
			},
			2: {
				Title:    "Concurency in Go",
				Quantity: 5,
			},
		},
	}
}

// read lock
func (bs *BookStore) getBookDetail(id int) *Book {
	bs.rwMutex.RLock()
	defer bs.rwMutex.RUnlock()
	book, found := bs.books[id]
	if !found {
		return nil
	}

	return &Book{
		Title:    book.Title,
		Quantity: book.Quantity,
	}
}

func (bs *BookStore) updateBookQueantity(id int, change int) {
	bs.rwMutex.Lock()
	defer bs.rwMutex.Lock()
	book, found := bs.books[id]
	if !found {
		fmt.Printf("Book With Id %d Not Found", id)
		return
	}
	book.Quantity += change
}
