package store

import "github.com/wim07101993/MusicRestService/models"

type Storer interface {
	GetAll() ([]models.Track, error)
	Get(id string) (*models.Track, error)
	Create(t *models.Track) (error)
	Delete(id string) (error)
}
