package spotify

import (
	models "Conversify/Server/Models"
	"net/http"
)

func GetSearch(w http.ResponseWriter, r *http.Request, song models.Song) string {
	/*
		search_name := song.Name
		for _, author := range song.Author {
			search_name = search_name + " " + author
		}

		query := utils.EncodeString(url.QueryEscape(search_name))

		body, err := utils.GETRequest(utils.SpotifySearch+query, Token.AccessToken, nil)
		utils.ErrorManager(utils.ReadResponseError, err)

		var results models.SearchItems
		err = json.Unmarshal(body, &results)
		utils.ErrorManager(utils.UnmarshalJSONError, err)

		for _, item := range results.Tracks.Items {
			track_result := TrackToSong(item)

			if utils.LcsComparation(song.Name, track_result.Name) {
				if utils.ContainsIgnoreCase(track_result.Author, song.Author[0]) && utils.CompareDuration(track_result.Duration, song.Duration, 10) {
					log.Printf("Song %s found.", search_name)
					return item.Uri
				}

				if utils.CompareDuration(track_result.Duration, song.Duration, 5) {
					log.Printf("Song %s found by Name similarity.", search_name)
					return item.Uri
				}
			}

			if utils.CompareDuration(track_result.Duration, song.Duration, 3) {
				log.Printf("Track %s found by duration: %s", song.Name, item.Name)
				return item.Uri
			}

		}

		log.Printf("Song %s not found.", search_name)*/

	return ""
}
