package main

import "fmt"

type Booker interface {
	SetName(n string) *BookRepository
	SetDescription(d string) *BookRepository
	SetSerial(s uint64) *BookRepository
	SetPrice(p float64) *BookRepository
}

type BookRepository struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Serial      uint64  `json:"serial"`
	Price       float64 `json:"price"`
	Popular     bool    `json:"popular"`
}

type HandlerFunc func(reader uint32) bool

func (br *BookRepository) SetName(n string) *BookRepository {
	br.Name = n
	return br
}

func (br *BookRepository) SetDescription(d string) *BookRepository {
	br.Description = d
	return br
}

func (br *BookRepository) SetSerial(s uint64) *BookRepository {
	br.Serial = s
	return br
}

func (br *BookRepository) SetPrice(p float64) *BookRepository {
	br.Price = p
	return br
}

func (br *BookRepository) SetPopular(n uint32, handler HandlerFunc) *BookRepository {
	br.Popular = handler(n)
	return br
}

func (br *BookRepository) Build() *BookRepository {
	return br
}

func main() {
	var userReaderAmount HandlerFunc
	userReaderAmount = func(n uint32) bool {
		return n > 500
	}

	netflixBook := &BookRepository{}
	netflixBook.Build().SetName("Netflix by patty mccord").SetDescription("Netflix Culture").SetPrice(200.00).SetPopular(600, userReaderAmount)

	fmt.Printf("%+v", netflixBook)
}
