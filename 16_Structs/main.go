package main

import (
	"fmt"
	"time"
)

// Structs are like classees
type Book struct {
	name         string
	author       string
	is_available bool
	price        float32
	pages        int
	rating       float32
	publisher    string
	created_at   time.Time
}

type Car struct {
	name        string
	modleNo     string
	numberPlate int
	color       string
	companyName string
	createdAt   time.Time
	updatedAt   time.Time
}

// Constructor Function
func newBook(name string, author string, is_available bool, price float32, pages int, rating float32, publisher string, created_at time.Time) *Book {
	// Initial Setup
	myBook := Book{
		name:         name,
		author:       author,
		is_available: is_available,
		price:        price,
		pages:        pages,
		rating:       rating,
		publisher:    publisher,
		created_at:   created_at,
	}

	return &myBook
}

// Creatring Behaviour Functions.
// Whe have to use pointers to modify the struct
func (b *Book) change_statuc(is_available bool) {
	b.is_available = is_available
}

// Dont use Pointers to get values
func (b *Book) getAmount() float32 {
	return b.price
}

func main() {

	myBook := newBook("The Alchemist", "Paulo Coelho", true, 100, 200, 4.5, "HarperCollins", time.Now())
	fmt.Println("My Book :->", myBook)

	book := Book{
		name:         "The Alchemist",
		author:       "Paulo Coelho",
		is_available: true,
		price:        100,
		pages:        200,
		rating:       4.5,
		publisher:    "HarperCollins",
	}

	book.change_statuc(false)
	book.created_at = time.Now()
	fmt.Println("My Book :->", book)

	book2 := Book{
		name:         "Hello",
		author:       "Hello",
		is_available: true,
		price:        100,
		pages:        200,
		rating:       4.5,
		publisher:    "HarperCollins",
		created_at:   time.Now(),
	}

	fmt.Println("My Book :->", book2)

	// Adding new Property to the book
	fmt.Println(book2.getAmount())
}
