package utils

import (
	models "Conversify/Server/Models"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/forPelevin/gomoji"
)

func GETRequest(endpoint, access_token string, body url.Values) ([]byte, error) {

	if body != nil {
		req, err := http.NewRequest("GET", endpoint, bytes.NewBufferString(body.Encode()))
		ErrorManager(CreateRequestError, err)
		return Request(req, access_token, "")
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	ErrorManager(CreateRequestError, err)

	return Request(req, access_token, "")
}

func POSTRequest(endpoint, access_token string, body io.Reader, content_type string) ([]byte, error) {

	if body != nil {
		req, err := http.NewRequest("POST", endpoint, body)
		ErrorManager(CreateRequestError, err)
		return Request(req, access_token, content_type)
	}

	req, err := http.NewRequest("POST", endpoint, nil)
	ErrorManager(CreateRequestError, err)
	return Request(req, access_token, "")
}

func Request(req *http.Request, access_token string, content_type string) ([]byte, error) {

	req.Header.Add("Authorization", "Bearer "+access_token)
	if len(content_type) != 0 {
		req.Header.Add("Content-Type", content_type)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	ErrorManager(SendRequestError, err)
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

//duration utils:

func ISO8601ToSec(duration string) int {
	duration = strings.TrimPrefix(duration, "PT")
	duration = strings.ToLower(duration)
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatal(err)
	}

	return int(d.Seconds())
}

func SecToISO8601(duration int) string {
	d := time.Duration(duration) * time.Second
	return fmt.Sprint("PT" + strings.ToUpper(d.String()))
}

func CompareDuration(duration_1 int, duration_2 int, delta int) bool {
	delta_duration := math.Abs(float64(duration_1 - duration_2))
	return delta_duration < float64(delta)
}

//String utils:

func CleanString(s string) string {
	result := gomoji.RemoveEmojis(s)
	result = strings.ToLower(result)
	if strings.Contains(result, " - ") {
		result = strings.ReplaceAll(result, " - ", " ")
	}
	return result
}

func CleanAuthorString(s string) string {
	s = strings.ToLower(s)
	if strings.Contains(s, " - topic") {
		s = strings.TrimSuffix(s, " - topic")
		return s
	}

	return ""
}

func EncodeString(s string) string {
	result := strings.ReplaceAll(s, "+", "%20")
	return result
}

func SearchURIString(sa []string) string {
	if len(sa) == 1 {
		return sa[0]
	}
	s := sa[0]
	for i := 1; i < len(sa); i++ {
		if len(s) == 0 {
			s = sa[i]
		}
		if len(sa[i]) != 0 {
			s = fmt.Sprint(s + "," + sa[i])
		}
	}
	return s
}

func ComparePlaylists(playlist_1 []models.Song, playlist_2 []models.Song) ([]models.Song, []models.Song) {
	var missing_songs_p1, not_found_p1 []models.Song
	found_songs_p2 := 0

	for i := 0; i < len(playlist_1); i++ {
		found := false
		for j := 0; j < len(playlist_2); j++ {
			if LcsComparation(playlist_1[i].Name, playlist_2[j].Name) && CompareDuration(playlist_1[i].Duration, playlist_2[j].Duration, 10) {
				playlist_2 = moveToLastPlace(playlist_2, j)
				found_songs_p2++
				found = true
				break
			}
			/* if playlist_1[i].Name == playlist_2[j].Name {
				if CompareDuration(playlist_1[i].Duration, playlist_2[j].Duration, 10) {
					playlist_2 = moveToLastPlace(playlist_2, j)
					found_songs_p2++
					found = true
					break
				}
			} */
		}
		if !found {
			missing_songs_p1 = append(missing_songs_p1, playlist_1[i])
		}
	}
	playlist_2 = playlist_2[:len(playlist_2)-found_songs_p2]
	found_songs_p2 = 0

	for i := 0; i < len(missing_songs_p1); i++ {
		found := false
		for j := 0; j < len(playlist_2); j++ {
			if LcsComparation(missing_songs_p1[i].Name, playlist_2[j].Name) && CompareDuration(missing_songs_p1[i].Duration, playlist_2[j].Duration, 10) {
				playlist_2 = moveToLastPlace(playlist_2, j)
				found_songs_p2++
				found = true
				break
			}
			/* if LcsComparation(missing_songs_p1[i].Name, playlist_2[j].Name) {
				if CompareDuration(missing_songs_p1[i].Duration, playlist_2[j].Duration, 10) {
					playlist_2 = moveToLastPlace(playlist_2, j)
					found_songs_p2++
					found = true
					break
				}
			} */
		}
		if !found {
			not_found_p1 = append(not_found_p1, missing_songs_p1[i])
		}
	}
	playlist_2 = playlist_2[:len(playlist_2)-found_songs_p2]

	return not_found_p1, playlist_2

}

func LcsComparation(s1, s2 string) bool {
	lcs_result := lcs(s1, s2)
	s1_len := len(s1)
	s2_len := len(s2)

	s_mean := (s1_len + s2_len) / 2

	return lcs_result > s_mean/2
}

// Función para calcular la longitud de la secuencia más larga común
func lcs(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Crear una matriz para almacenar las subsecuencias comunes más largas
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Calcular la matriz dp utilizando el algoritmo de LCS
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return dp[m][n]
}

func moveToLastPlace(slice []models.Song, index int) []models.Song {
	if index >= len(slice) || index < 0 {
		return []models.Song{}
	}

	if index == 0 {
		temp := slice[1:]
		return append(temp, slice[0])
	}

	value_to_move := slice[index]
	slice = append(slice[:index], slice[index+1:]...)
	return append(slice, value_to_move)

}

func ContainsIgnoreCase(slice []string, str string) bool {
	for _, item := range slice {
		log.Println(item + " " + str)
		if strings.EqualFold(item, str) {
			return true
		}
	}
	return false
}
