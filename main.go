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
		go func(id int){
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache:")
				b.Display()
			}
		}(id)

		// if not fetch from db	
		go func(id int){
			if b, ok := queryDB(id); ok {
				fmt.Println("from DB:")
				b.Display()
			}
		}(id)

		// fmt.Println("Book not found!")
		time.Sleep(150 * time.Millisecond)
	}
	
	time.Sleep(2 * time.Second)
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