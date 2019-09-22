package main

import (
	"log"
	"net/http"
)

func middlewareAuthenticate(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareAuthenticate")
		next.ServeHTTP(w, r)
	})
}

func middlewareAuthorize(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareAuthorize")
		next.ServeHTTP(w, r)
	})
}

func middlewareSuccess(w http.ResponseWriter, r *http.Request) {
	log.Println("Success")
	w.Write([]byte("OK"))
}

func main() {
	success := http.HandlerFunc(middlewareSuccess)
	http.Handle("/test", middlewareAuthenticate(middlewareAuthorize(success)))
	http.ListenAndServe(":3000", nil)
}