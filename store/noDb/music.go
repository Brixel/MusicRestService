package noDb

import (
	"github.com/wim07101993/MusicRestService/models"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// TrackStore represents a store to get, create and delete elements in a temporarily stored collection
type TrackStore struct {
	values []models.Track
}

// NewTrackStore creates a pointer to a new value of the TrackStore type, prepopulated with data.
func NewTrackStore() (*TrackStore) {
	// create a pointer to a new value of the type
	return &TrackStore{
		values: []models.Track{
			{
				Id:     bson.NewObjectId(),
				Title:  "Rock Or Bust",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Play Ball",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Rock The Blues Away",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Miss Adventure",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Dogs Of War",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Got Some Rock & Roll Thunder",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Hard Times",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Baptism By Fire",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Rock The House",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Sweet Candy",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Emission Control",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},

			{
				Id:     bson.NewObjectId(),
				Title:  "Reconnect",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Resistance",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Boss Mode",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "EDM Trend Machine",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "404",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Begin Again",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Give It Up",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "D.I.M.H",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Micropenis",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Superstar",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Red Dawn",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			{
				Id:     bson.NewObjectId(),
				Title:  "Kaleidoscope",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
		},
	}
}

// Create stores a new track in memory
func (ts *TrackStore) Create(t *models.Track) error {
	// set the max id
	t.Id = bson.NewObjectId()
	// add the new track to the slice of tracks
	ts.values = append(ts.values, *t)
	return nil
}

// GetAll retrieves all tracks in memory
func (ts *TrackStore) GetAll() ([]models.Track, error) {
	// return all the tracks
	return ts.values, nil
}

// Get retrieves a single track from memory, that with the given id
func (ts *TrackStore) Get(id bson.ObjectId) (*models.Track, error) {
	// iterate over all the tracks and return the track with the right id
	for _, t := range ts.values {
		if t.Id == id {
			return &t, nil
		}
	}

	// if no id matches, return an error
	return nil, errors.New("not found")
}

// Delete deletes a track from memory, that with the given id
func (ts *TrackStore) Delete(id bson.ObjectId) error {
	// iterate over all the tracks and delete the track with the right id
	for i, v := range ts.values {
		if v.Id == id {
			ts.values = append(ts.values[:i], ts.values[i+1:]...)
			return nil
		}
	}

	// if no id matches, return an error
	return errors.New("not found")
}
