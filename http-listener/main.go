package main

// defining main package is necessary because that tells the compiler that this should
// be compiled as a standalone executable program.

// without it the compiler won't make an executable binary

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// entry point of the executable !

	// problem statement: Create a web server!
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")

		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			//w.WriteHeader(http.StatusBadRequest)
			//w.Write([]byte("Oops"))
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data :%s", d)

		fmt.Fprintf(w, "Hello %s", d)

	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})
	http.ListenAndServe(":9090", nil)

}
