package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/brixel/MusicRestService/models"
	"github.com/brixel/MusicRestService/store"
	"gopkg.in/mgo.v2/bson"
	//"github.com/brixel/MusicRestService/store/noDb"
	"github.com/brixel/MusicRestService/store/mongoDb"
)

// TrackController represents the controller to modify the music collection
type TrackController struct {
	storer store.Storer
}

// NewTrackController creates a pointer to a new value of the TrackController type
func NewTrackController() *TrackController {
	ts, err := mongoDb.NewTrackStore("mongodb://localhost", "music", "tracks")

	if err != nil {
		panic(err)
	}

	return &TrackController{
		//storer: noDb.NewTrackStore(),
		storer: ts,
	}
}

func (tc *TrackController) GetAllTracks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get all values from the store
	ts, err := tc.storer.GetAll()
	if err != nil {
		fmt.Println("Error while trying to get all tracks:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Marshal the value to json
	tsj, err := json.Marshal(ts)
	if err != nil {
		fmt.Println("Error while trying to marshal tracks:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Write the content-type, statuscode and content with the response writer
	tc.respondWithJSON(w, tsj, http.StatusOK)
}

// GetTrack retrieves a single track from the music collection
func (tc *TrackController) GetTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Get the track to serve
	t, err := tc.storer.Get(oid)
	if err != nil {
		fmt.Println("Error while trying to get the track:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if t.Id == "" {
		// Write the content-type, statuscode and content with the response writer
		tc.respondWithJSON(w, nil, http.StatusOK)
		return
	}

	// Marshal the value to json
	tj, err := json.Marshal(t)
	// Check for errors
	if err != nil {
		fmt.Println("Error while trying to marshal track:", err)
	}

	// Write the content-type, statuscode and content with the response writer
	tc.respondWithJSON(w, tj, http.StatusOK)
}

func (tc *TrackController) CreateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Create empty value to fill with body of request
	t := models.Track{}

	// Fill the track's fields
	json.NewDecoder(r.Body).Decode(&t)

	err := tc.storer.Create(&t)
	if err != nil {
		fmt.Println("Error while trying to create track:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Marshal the value to json
	tj, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Error while trying to marshal track:", err)
	}

	// Write the content-type, statuscode and content with the response writer
	tc.respondWithJSON(w, tj, http.StatusCreated)
}

func (tc *TrackController) DeleteTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	err := tc.storer.Delete(oid)

	if err != nil {
		fmt.Println("Error while trying to delete track:", err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (tc *TrackController) respondWithJSON(w http.ResponseWriter, json []byte, statusCode int) {
	// Write the content-type, statuscode and content with the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s", json)
}
