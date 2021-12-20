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

// Channel
var cacheCh = make(chan books.Book)
var dbCh = make(chan books.Book)

// Random number
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func waitGroupMutex() {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	// Loop 1 to 10 random ids
	for i:=0; i<10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		// Look for book with id in cache
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex){
			if b, ok := queryCache(id, mu); ok {
				fmt.Println("from cache:")
				b.Display()
			}
			wg.Done()
		}(id, wg, mu)

		// if not fetch from db
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex){
			if b, ok := queryDB(id, mu); ok {
				fmt.Println("from DB:")
				b.Display()
			}
			wg.Done()
		}(id, wg, mu)

		// fmt.Println("Book not found!")
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

// Fetch from cache func

func queryCache(id int, mu *sync.RWMutex) (books.Book, bool) {
	mu.RLock()
	b, ok := cache[id]
	mu.RUnlock()
	return b, ok
}

// Fetch from db func

func queryDB(id int, mu *sync.RWMutex) (books.Book, bool) {

	time.Sleep(100 * time.Millisecond)

	for _, b := range books.BookDB {
		if b.ID == id {
			mu.Lock()
			cache[id] = b
			mu.Unlock()
			return b, true
		}
	}

	return books.Book{}, false
}

func channelsIf () {
	wg := sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func (ch <-chan int) {
		msg, ok := <-ch
		fmt.Println(msg, ok)
		wg.Done()
	}(ch)

	go func (ch chan<- int) {
		close(ch)
		wg.Done()
	}(ch)
	
	wg.Wait()
}

func channelsLoop () {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func (ch chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	
	go func (ch chan int, wg *sync.WaitGroup) {
		for i :=0; i<10; i++ {
			ch<-i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func channelsSelect() {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	// Loop 1 to 10 random ids
	for i:=0; i<10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		// Look for book with id in cache
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex, ch chan <- books.Book){
			if b, ok := queryCache(id, mu); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, mu, cacheCh)

		// if not fetch from db
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex, ch chan <- books.Book){
			if b, ok := queryDB(id, mu); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, mu, dbCh)

		go func(cacheCh, dbCh <-chan books.Book){
			select {
			case b := <- cacheCh:
				fmt.Println("from cache")
				b.Display()
				<- dbCh
			case b := <- dbCh:
				fmt.Println("from db")
				b.Display()
			}
		}(cacheCh, dbCh)

		// fmt.Println("Book not found!")
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func main() {
	// channelsIf()
	// channelsLoop()
	// waitGroupMutex()
	channelsSelect()
}