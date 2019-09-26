package main

import (
	"context"
	"errors"
	"log"
)

// --------- library code
type MiddlewareFunc func(HandlerFunc) HandlerFunc
type HandlerFunc func(context.Context) error

type Employer interface {
	Use(mw MiddlewareFunc)
	Run()
}

type Employee struct {
	middleware []MiddlewareFunc

	Firstname string
	Lastname  string
	Skills    []string
}

func NewEmployee(f, l string, sk []string) Employer {
	return &Employee{
		Firstname: f,
		Lastname:  l,
		Skills:    sk,
	}
}

func (e *Employee) Use(mw MiddlewareFunc) {
	e.middleware = append(e.middleware, mw)
}

func (e *Employee) Run() {
	h := func(c context.Context) error {
		var h HandlerFunc
		// for i := len(e.middleware) - 1; i >= 0; i-- {
		// 	h = e.middleware[i](h)
		// }
		for i := 0; i < len(e.middleware); i++ {
			h = e.middleware[i](h)
		}

		return h(c)
	}

	var c context.Context
	if err := h(c); err != nil {
		log.Println("Execute error", err)
	}

	log.Print("Hiring starting...")
}

// ---------

// --------- production code
// func Interview() MiddlewareFunc {
// 	return func(next HandlerFunc) HandlerFunc {
// 		return func(c context.Context) error {
// 			log.Println("interview...")
// 			return nil
// 		}

// 		return next
// 	}
// }

// func TechnicalInterview() MiddlewareFunc {
// 	return func(next HandlerFunc) HandlerFunc {
// 		log.Println("technical interview...")
// 		return func(c context.Context) error {

// 			return nil
// 		}

// 		return next
// 	}
// }

func Interview(next HandlerFunc) HandlerFunc {
	log.Println("interview...")
	return func(c context.Context) error {

		return nil
	}
}

func TechnicalInterview(next HandlerFunc) HandlerFunc {
	log.Println("technical...")
	return func(c context.Context) error {

		return errors.New("technical not ok")
	}
}

func main() {
	golf := NewEmployee("Teerapong", "Singthong", []string{"go", "js"})
	golf.Use(Interview)
	golf.Use(TechnicalInterview)
	golf.Run()
}
