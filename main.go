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
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache:")
			b.Display()
			continue
		}

		// if not fetch from db	
		if b, ok := queryDB(id); ok {
			fmt.Println("from DB:")
			b.Display()
			continue
		}

		fmt.Println("Book not found!")
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

	for _, b := range books.BookDB {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	
	return books.Book{}, false
}