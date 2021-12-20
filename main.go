package main

import (
	"concurrent/books"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Cache
var cache = map[int]books.Book{}

// Random number
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	// Loop 1 to 10 random ids
	for i:=0; i<10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		// Look for book with id in cache
		go func(id int, wg *sync.WaitGroup){
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache:")
				b.Display()
			}
			wg.Done()
		}(id, wg)

		// if not fetch from db	
		go func(id int, wg *sync.WaitGroup){
			if b, ok := queryDB(id); ok {
				fmt.Println("from DB:")
				b.Display()
			}
			wg.Done()
		}(id, wg)

		// fmt.Println("Book not found!")
		// time.Sleep(150 * time.Millisecond)
	}
	
	wg.Wait()
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
			// cache[id] = b
			return b, true
		}
	}
	
	return books.Book{}, false
}