package models

type Song struct {
	Name     string
	Author   []string
	Duration int //in secs
}

type Playlist struct {
	Name  string
	Songs []Song
}

type UserPlaylists struct {
	Playlists []Playlist
}

type User struct {
	ID   string
	Name string
}
