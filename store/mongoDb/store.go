package mongoDb

import (
	"gopkg.in/mgo.v2"
	"github.com/brixel/MusicRestService/models"
	"gopkg.in/mgo.v2/bson"
)

// TrackStore represents a store to get, create and delete elements in a mongo database
type TrackStore struct {
	mgoSession       *mgo.Session
	connectionString string
	db               string
	collection       string
}

// NewTrackStore creates a pointer to a new value of the TrackStore type, using the connection string, db and collection.
func NewTrackStore(connectionString string, db string, collection string) (*TrackStore, error) {
	s, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return &TrackStore{
		mgoSession:       s,
		connectionString: connectionString,
		db:               db,
		collection:       collection,
	}, nil
}

// Create stores a new track in the database
func (ts *TrackStore) Create(t *models.Track) error {
	t.Id = bson.NewObjectId()
	return ts.mgoSession.DB(ts.db).C(ts.collection).Insert(t)
}

// GetAll retrieves all tracks from the database
func (ts *TrackStore) GetAll() ([]models.Track, error) {
	var ti []models.Track

	if err := ts.mgoSession.DB(ts.db).C(ts.collection).Find(nil).All(&ti); err != nil {
		return nil, err
	}

	return ti, nil
}

// Get retrieves a single track from the database, that with the given id
func (ts *TrackStore) Get(id bson.ObjectId) (*models.Track, error) {
	t := models.Track{}

	if err := ts.mgoSession.DB(ts.db).C(ts.collection).FindId(id).One(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

// Delete deletes a track from database, that with the given id
func (ts *TrackStore) Delete(id bson.ObjectId) error {
	return ts.mgoSession.DB(ts.db).C(ts.collection).RemoveId(id)
}
