package main

import (
	"fmt"
)

type handleFunc func(n string)

func hello() handleFunc{
	text := "Hello"
	return func(s string) {
		fmt.Println(text + " " + s)
	}
}

func name(n string, next handleFunc) handleFunc{
	text := "my name is " + n
	return func(s string) {
		next(text + " " + s)
	}
}

func company(next handleFunc) handleFunc {
	text := "from"
	return func(s string) {
		next(text + " " + s)
	}
}

func main()  {
	// Hello my name is Brown from LINE

	brown := company(name("Brown", hello()))
	brown("LINE")
}