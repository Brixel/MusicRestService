package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"github.com/wim07101993/MusicRestService/models"
	"encoding/json"
)

func main() {
	// Create a value of a router
	r := httprouter.New()

	// Add a route for the router
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Send the message as a response through the response writer
		fmt.Fprint(w, "Hello Brixel!")
	})

	// Add a route to serve a track
	r.GET("/tracks/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Create an example track to serve
		t := models.Track{
			Title:"Song 2",
			Artist: "Blur",
			Album: "Blur",
			Genre: "Rock",
			Id: "1",
		}

		// Marshal the value to json
		tj, err := json.Marshal(t)
		// Check for errors
		if err!=nil {
			fmt.Println("Error while trying to marshal track:", err)
		}

		// Write the content-type, statuscode and content with the response writer
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", tj)
	})

	r.POST("/tracks", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Create empty value to fill with body of request
		t := models.Track{}

		// Fill the track's fields
		json.NewDecoder(r.Body).Decode(&t)

		// Add an id
		t.Id = "someId"

		// Marshal the value to json
		tj, err := json.Marshal(t)
		// Check for errors
		if err!=nil {
			fmt.Println("Error while trying to marshal track:", err)
		}

		// Write the content-type, statuscode and content with the response writer
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", tj)
	})

	r.DELETE("/tracks", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO: there is no database to remove from yet
		w.WriteHeader(http.StatusNotImplemented)
	})

	// Give some feedback to the console
	fmt.Println("Serving on localhost:3000")

	// Start the server by listening on the right port
	err := http.ListenAndServe("localhost:3000", r)
	// If something goes wrong while serving, notify over the console
	if err != nil {
		fmt.Println("Error while trying to serve:", err)
	}
}
