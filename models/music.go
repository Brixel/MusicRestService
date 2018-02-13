package models

// Track represents the structure of our music data
type Track struct {
	// Id is the unique ide by which we can search our music
	Id     string	`json:"id"`
	// Title is the title of the track
	Title  string	`json:"title"`
	// Artist is the artist that created the track
	Artist string	`json:"artist"`
	// Album is the album in which the track is found
	Album  string	`json:"album"`
	// Genre is the genre of the track
	Genre  string	`json:"genre"`
}
