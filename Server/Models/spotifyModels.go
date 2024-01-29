package models

type CurrentUserPlaylists struct {
	Items []PlaylistS `json:"items"`
}

type PlaylistS struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		ID string `json:"id"`
	} `json:"owner"`
}

type PlaylistTracks struct {
	Items []struct {
		Track Track `json:"track"`
	} `json:"items"`
	Total int `json:"total"`
}

type Track struct {
	Name     string `json:"name"`
	Duration int    `json:"duration_ms"`
	Artists  []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Album struct {
		Name string `json:"name"`
	} `json:"album"`
	Uri string `json:"uri"`
}

type SearchItems struct {
	Tracks struct {
		Items []Track `json:"items"`
	} `json:"tracks"`
}
