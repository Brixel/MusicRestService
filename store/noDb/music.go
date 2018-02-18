package noDb

import (
	"github.com/wim07101993/MusicRestService/models"
	"strconv"
	"errors"
)

// TrackStore represents a store to get, create and delete elements in a temporarily stored collection
type TrackStore struct {
	values []models.Track
	maxId  int
}

// NewMusicStore creates a pointer to a new value of the TrackStore type, prepopulated with data.
func NewMusicStore() (*TrackStore) {
	// create a pointer to a new value of the type
	return &TrackStore{
		values: []models.Track{
			models.Track{
				Id:     "0",
				Title:  "Rock Or Bust",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "1",
				Title:  "Play Ball",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "2",
				Title:  "Rock The Blues Away",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "3",
				Title:  "Miss Adventure",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "4",
				Title:  "Dogs Of War",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "5",
				Title:  "Got Some Rock & Roll Thunder",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "6",
				Title:  "Hard Times",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "7",
				Title:  "Baptism By Fire",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "8",
				Title:  "Rock The House",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "9",
				Title:  "Sweet Candy",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},
			models.Track{
				Id:     "10",
				Title:  "Emission Control",
				Artist: "AC/DC",
				Album:  "Rock Or Burst",
				Genre:  "Rock",
			},

			models.Track{
				Id:     "11",
				Title:  "Reconnect",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "12",
				Title:  "Resistance",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "13",
				Title:  "Boss Mode",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "14",
				Title:  "EDM Trend Machine",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "15",
				Title:  "404",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "16",
				Title:  "Begin Again",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "17",
				Title:  "Give It Up",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "18",
				Title:  "D.I.M.H",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "19",
				Title:  "Micropenis",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "20",
				Title:  "Superstar",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "21",
				Title:  "Red Dawn",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
			models.Track{
				Id:     "22",
				Title:  "Kaleidoscope",
				Artist: "Knife Party",
				Album:  "Abandon Ship",
				Genre:  "Dance",
			},
		},
		maxId: 22,
	}
}

// GetAll retrieves all tracks in memory
func (ts *TrackStore) GetAll() ([]models.Track, error) {
	// return all the tracks
	return ts.values, nil
}

// Get retrieves a single track from memory, that with the given id
func (ts *TrackStore) Get(id string) (*models.Track, error) {
	// iterate over all the tracks and return the track with the right id
	for _, t := range ts.values {
		if t.Id == id {
			return &t, nil
		}
	}

	// if no id matches, return an error
	return nil, errors.New("not found")
}

// Create stores a new track in memory
func (ts *TrackStore) Create(t *models.Track) error {
	// increment the max id
	ts.maxId++
	// set the max id
	t.Id = strconv.Itoa(ts.maxId)
	// add the new track to the slice of tracks
	ts.values = append(ts.values, *t)
	return nil
}

// Delete deletes a track from memory, that with the given id
func (ts *TrackStore) Delete(id string) error {
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
