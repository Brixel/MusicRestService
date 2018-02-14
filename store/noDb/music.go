package noDb

import (
	"github.com/wim07101993/MusicRestService/models"
	"strconv"
)

var values = []models.Track{
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
}
var maxId = 22

type TrackStore struct {
}

func NewMusicStore() (*TrackStore) {
	return &TrackStore{}
}

func (ts TrackStore) GetAll() []models.Track {
	return values
}

func (ts TrackStore) Get(id string) *models.Track {
	for _, t := range values {
		if t.Id == id {
			return &t
		}
	}
	return nil
}

func (ts TrackStore) Create(t *models.Track) *models.Track {
	maxId++
	t.Id = strconv.Itoa(maxId)
	values = append(values, *t)
	return t
}

func (ts TrackStore) Delete(id string) bool {
	for i, v := range values {
		if v.Id == id{
			values = append(values[:i], values[i+1:]...)
			return true
		}
	}
	return false
}
