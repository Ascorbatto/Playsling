package models

type CurrentUserPlaylistsS struct {
	Items []PlaylistInfoS `json:"items"`
}

type PlaylistInfoS struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Owner       struct {
		ID string `json:"id"`
	} `json:"owner"`
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
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

type UserS struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
}
