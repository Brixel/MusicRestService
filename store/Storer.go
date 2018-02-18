package store

import "github.com/wim07101993/MusicRestService/models"

// Storer is the interface used to get, create and delete elements from a database.
type Storer interface {
	// GetAll gets all the tracks from a database
	GetAll() ([]models.Track, error)
	// Get returns the track with the given id from the database
	Get(id string) (*models.Track, error)
	// Create creates a new track in the database using the given data and adds an id
	Create(t *models.Track) (error)
	// Delete deletes the track with the given id from the database
	Delete(id string) (error)
}
