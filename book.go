package main

import "fmt"

// 1. Book struct

type Book struct{
	ID int
	Title string
}

// 2. Printing method for struct

func (b Book) Display() {
	fmt.Printf(
		"Title:\t\t%q\n"+
			"ID:\t%v\n", b.Title, b.ID)
}

// 3. Array of books (db)

var books = []Book{
	Book{
		ID: 1,
		Title: "Harry Potter 1",
	},
	Book{
		ID: 2,
		Title: "Harry Potter 2",
	},
	Book{
		ID: 3,
		Title: "Harry Potter 3",
	},
	Book{
		ID: 4,
		Title: "Harry Potter 4",
	},
	Book{
		ID: 5,
		Title: "Harry Potter 5",
	},
	Book{
		ID: 6,
		Title: "Harry Potter 6",
	},
	Book{
		ID: 7,
		Title: "Harry Potter 7",
	},
	Book{
		ID: 8,
		Title: "Harry Potter 8",
	},
	Book{
		ID: 9,
		Title: "Harry Potter 9",
	},
	Book{
		ID: 10,
		Title: "Harry Potter 10",
	},
}