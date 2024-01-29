package spotify

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
)

var Tracks []models.Song

func GetTracks(playlist models.PlaylistTracks) {
	for _, items := range playlist.Items {
		track := TrackToSong(items.Track)

		Tracks = append(Tracks, track)
	}
}

func TrackToSong(track models.Track) models.Song {
	song := models.Song{
		Name:     utils.CleanString(track.Name),
		Duration: track.Duration / 1000,
	}

	for _, artist := range track.Artists {
		song.Author = append(song.Author, utils.CleanAuthorString(artist.Name))
	}

	return song
}
