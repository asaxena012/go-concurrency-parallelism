package main

import (
	"concurrent/books"
	"fmt"
	"math/rand"
	"time"
)

// Cache
var cache = map[int]books.Book{}

// Random number
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// Loop 1 to 10 random ids

	for i:=0; i<10; i++ {
		id := rnd.Intn(10) + 1

		// Look for book with id in cache
		b1, ok := queryCache(id)

		if ok {
			// if found print from cache
			fmt.Println("from cache:")
			b1.Display()
			continue
		}

		// if not fetch from db
		b2, ok := queryDB(id)

		if ok {
			fmt.Println("from DB:")
			b2.Display()
		} else {
			fmt.Println("Book not found!")
		}

		time.Sleep(150 * time.Millisecond)
	}
	

}

// Fetch from cache func

func queryCache(id int) (books.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

// Fetch from db func

func queryDB(id int) (books.Book, bool) {

	time.Sleep(100 * time.Millisecond)

	for i := 0; i< len(books.BookDB); i++ {
		b := books.BookDB[i]
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return books.Book{}, false
}