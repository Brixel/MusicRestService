package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func main() {
	r := httprouter.New()

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello Brixel!")
	})

	fmt.Println("Serving on localhost:3000")

	err := http.ListenAndServe("localhost:3000", r)
	if err != nil {
		fmt.Println("Error while trying to serve:", err)
	}
}