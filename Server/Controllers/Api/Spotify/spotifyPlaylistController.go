package spotify

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	PlaylistInfo         models.PlaylistS
	CurrentUserPlaylists models.CurrentUserPlaylists
)

// Get a list of the playlists owned or followed by the current Spotify user.
func GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) {

	body, err := utils.GETRequest(utils.SpotifyCurrentUserPlaylists, Token.AccessToken, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	err = json.Unmarshal(body, &CurrentUserPlaylists)
	utils.ErrorManager(utils.UnmarshalJSONError, err)
}

// Get a list of the tracks of the specified Playlist in the request.
func GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request) {

	playlist_id := r.URL.Query().Get("playlist")
	endpoint := fmt.Sprintf(utils.SpotifyPlaylistItemsInfo+"%s/tracks", playlist_id)

	body, err := utils.GETRequest(endpoint, Token.AccessToken, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var tracks models.PlaylistTracks
	err = json.Unmarshal(body, &tracks)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	GetTracks(tracks)

	http.Redirect(w, r, utils.Host+"/convert_yt_to_s?destination="+PlaylistInfo.ID+"&exists=true", http.StatusFound)
}

// Creates a new playlist with the Name specified in the request.
// If already created, will return the items of said playlist.
func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	query := fmt.Sprint(r.URL.Query().Get("user_id") + "/playlists")
	name := r.URL.Query().Get("name")
	data := map[string]string{
		"name": name,
	}
	content_type := "application/json"

	for _, playlist := range CurrentUserPlaylists.Items {
		if playlist.Name == name {
			PlaylistInfo = playlist

			http.Redirect(w, r, utils.Host+"/playlist/spotify?playlist="+PlaylistInfo.ID, http.StatusFound)
			return
		}
	}

	json_data, err := json.Marshal(data)
	utils.ErrorManager(utils.MarshalJSONError, err)

	body, err := utils.POSTRequest(utils.SpotifyCreatePlaylist+query, Token.AccessToken, bytes.NewBuffer(json_data), content_type)
	utils.ErrorManager(utils.ReadResponseError, err)

	err = json.Unmarshal(body, &PlaylistInfo)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	http.Redirect(w, r, utils.Host+"/convert_yt_to_s?destination="+PlaylistInfo.ID+"&exists=false", http.StatusFound)
}

// Post the items specified on the URIs
func PostPlaylistItem(w http.ResponseWriter, r *http.Request, playlist_id string, uris []string) {
	query := fmt.Sprint("/tracks?uris=" + utils.SearchURIString(uris))

	utils.POSTRequest(utils.SpotifyPlaylistItemsInfo+playlist_id+query, Token.AccessToken, nil, "")
}
