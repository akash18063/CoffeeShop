package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/akash18063/CoffeeShop/handlers"
)

func main() {
	// Basic HttpHandlers
	/*
		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Hello World")
			d, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(rw, "Oops", http.StatusBadRequest)
				return
			}
			fmt.Fprintf(rw, "Hellow %s", d)
		})

		http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Goodbye World")
		})
	*/

	//using handlers
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.ListenAndServe(":9000", hh) // second param handler
}
