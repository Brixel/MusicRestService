package models

// Music represents the structure of our music data
type Music struct {
	// Id is the unique ide by which we can search our music
	Id     string
	// Title is the title of the music
	Title  string
	// Artist is the artist that created the music
	Artist string
	// Album is the album in which the music is found
	Album  string
	// Genre is the genre of the music
	Genre  string
}
