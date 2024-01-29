package models

type Song struct {
	Name     string
	Author   []string
	Duration int //in secs
}

type Playlist struct {
	Songs []Song
}

type AuthToken struct {
	AccessToken string
}
