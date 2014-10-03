// functionfirstcitizen project main.go
package main

import (
	"fmt"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", <-nextID) // HL
}

func main() {
	http.HandleFunc("/next", handler)
	go func() {
		for i := 0; ; i++ {
			nextID <- i // HL
		}
	}()
	http.ListenAndServe("localhost:3000", nil)
}
