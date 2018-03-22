package models

import "gopkg.in/mgo.v2/bson"

// Track represents the structure of our music data
type Track struct {
	// Id is the unique ide by which we can search our music
	Id bson.ObjectId `json:"id" bson:"_id"`
	// Title is the title of the track
	Title string `json:"title" bson:"title"`
	// Artist is the artist that created the track
	Artist string `json:"artist" bson:"artist"`
	// Album is the album in which the track is found
	Album string `json:"album" bson:"album"`
	// Genre is the genre of the track
	Genre string `json:"genre" bson:"genre"`
}
