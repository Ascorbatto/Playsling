package youtube

import (
	spotify "Conversify/Server/Controllers/Api/Spotify"
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	"log"
	"net/http"
)

var (
	PlaylistItems models.PlaylistItems
	Playlist      models.PlaylistInfoYT
	PlaylistId    string
)

// Retrieves the playlist details (name, author, description...).
func GetPlaylistInfo(w http.ResponseWriter, r *http.Request) {
	/*
		endpoint := fmt.Sprint(utils.YoutubePlaylist + PlaylistId)
		body, err := utils.GETRequest(endpoint, Token.AccessToken, nil)
		utils.ErrorManager(utils.ReadResponseError, err)

		err = json.Unmarshal(body, &Playlist)
		utils.ErrorManager(utils.UnmarshalJSONError, err)

		log.Println("Playlist info retrieved without issues.")*/
}

// Retrieves the playlist items.
func GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request) {
	/*PlaylistId = r.URL.Query().Get("playlist")

	endpoint := fmt.Sprintf(utils.YoutubePlaylistItems + PlaylistId)
	body, err := utils.GETRequest(endpoint, Token.AccessToken, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	err = json.Unmarshal(body, &PlaylistItems)
	utils.ErrorManager(utils.UnmarshalJSONError, err)
	//GetPlaylistInfo(w, r)
	//GetSongs()

	//log.Println("Playlist items info retrieved without issues.")

	//http.Redirect(w, r, utils.Host+"/create/spotify?user_id=12164928183&name="+Playlist.Items[0].Snippet.Title, http.StatusFound)
	//
	//fmt.Println(songs)
	*/
}

func TestPostPlaylistItem(w http.ResponseWriter, r *http.Request) {
	playlist_id := r.URL.Query().Get("playlist")
	video_id := r.URL.Query().Get("video")
	PostPlaylistItem(w, r, playlist_id, video_id)
}

func PostPlaylistItem(w http.ResponseWriter, r *http.Request, playlist_id, video_id string) {
	/*
		data := map[string]interface{}{
			"snippet": map[string]interface{}{
				"playlistId": playlist_id,
				"resourceId": map[string]string{
					"kind":    "youtube#video",
					"videoId": video_id,
				},
			},
		}

		content_type := "application/json"

		json_data, err := json.Marshal(data)
		utils.ErrorManager(utils.MarshalJSONError, err)

		body, err := utils.POSTRequest(utils.YoutubePostPlaylistItem, Token.AccessToken, bytes.NewBuffer(json_data), content_type)
		utils.ErrorManager(utils.ReadResponseError, err)

		fmt.Println(string(body))*/
}

func ToSpotify(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.Query().Get("destination")
	exist := r.URL.Query().Get("exists")
	var search_Uris []string
	songs_to_tracks := Songs

	if exist == "true" {
		songs_to_tracks, _ = utils.ComparePlaylists(Songs, spotify.Tracks)
	}
	for _, song := range songs_to_tracks {
		search := spotify.GetSearch(w, r, song)
		search_Uris = append(search_Uris, search)

	}

	spotify.PostPlaylistItem(w, r, destination, search_Uris)
	log.Println("Items posted into the playlist correctly.")
}
