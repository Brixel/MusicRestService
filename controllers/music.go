package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/wim07101993/MusicRestService/models"
)

// TrackController represents the controller to modify the music collection
type TrackController struct{}

// NewTrackController creates a pointer to a new value of the TrackController type
func NewTrackController() *TrackController {
	return &TrackController{}
}

// GetTrack retrieves a single track from the music collection
func (tc TrackController) GetTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Create an example track to serve
	t := models.Track{
		Title:  "Song 2",
		Artist: "Blur",
		Album:  "Blur",
		Genre:  "Rock",
		Id:     "1",
	}

	// Marshal the value to json
	tj, err := json.Marshal(t)
	// Check for errors
	if err != nil {
		fmt.Println("Error while trying to marshal track:", err)
	}

	// Write the content-type, statuscode and content with the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", tj)
}

func (tc TrackController) CreateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
}

func (tc TrackController) DeleteTrack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO: there is no database to remove from yet
	w.WriteHeader(http.StatusNotImplemented)
}