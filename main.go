package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"github.com/brixel/MusicRestService/controllers"
)

func main() {
	// Create a value of a router
	r := httprouter.New()

	// Add a route for the router
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Send the message as a response through the response writer
		fmt.Fprint(w, "Hello Brixel!")
	})

	// Get an instance of the TrackController type
	tc := controllers.NewTrackController()

	// Add a route to serve a track
	r.GET("/tracks/", tc.GetAllTracks)
	r.GET("/tracks/:id", tc.GetTrack)
	r.POST("/tracks", tc.CreateTrack)
	r.DELETE("/tracks/:id", tc.DeleteTrack)

	// Give some feedback to the console
	fmt.Println("Serving on localhost:3000")

	// Start the server by listening on the right port
	err := http.ListenAndServe("localhost:3000", r)
	// If something goes wrong while serving, notify over the console
	if err != nil {
		fmt.Println("Error while trying to serve:", err)
	}
}
