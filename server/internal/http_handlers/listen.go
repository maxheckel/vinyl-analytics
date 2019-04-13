package http_handlers

import (
	"app/internal/models"
	"app/internal/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func (h *handlers) Listen(writer http.ResponseWriter, request *http.Request) {
	read_form, _ := request.MultipartReader()

	for {
		part, err_part := read_form.NextPart()
		if err_part == io.EOF {
			break
		}
		if part.FormName() == "data" {

			buf := new(bytes.Buffer)
			buf.ReadFrom(part)
			fileName := fmt.Sprintf("/tmp/%s.wav", randomString())
			err := ioutil.WriteFile(fileName, []byte(buf.String()), 7777)
			if err != nil {
				responses.NewInternalError("could not write file to server", writer)
				return
			}

			out, err := exec.Command("/usr/bin/php", "/Users/maxheckel/Sites/vinyl-analytics/server/server/test.php", fileName).Output()
			if err != nil {
				log.Fatal(err)
			}
			go os.Remove(fileName)
			//outString := `{"metadata":{"timestamp_utc":"2019-04-11 02:11:58","music":[{"label":"Roadrunner Records","play_offset_ms":20980,"external_ids":{},"result_from":1,"artists":[{"name":"Coheed and Cambria"}],"external_metadata":{},"title":"Old Flames","duration_ms":350300,"album":{"name":"Old Flames"},"score":100,"acrid":"3de981f9fc8adb792648cd11b5cd0488","release_date":"2018-09-27"},{"label":"Rogue Johnsen","play_offset_ms":16920,"external_ids":{},"result_from":1,"external_metadata":{},"acrid":"c86678614798c26d28bab96cff938935","title":"Tides of Time","duration_ms":272160,"album":{"name":"Way Back"},"score":82,"artists":[{"name":"Rogue Johnsen"}],"release_date":"2016-01-01"},{"label":"eOne Music International Classics","play_offset_ms":4500,"external_ids":{},"result_from":1,"external_metadata":{},"acrid":"e7170f840e8747c9540f6526aeb04795","title":"Twelve Preludes - 4","duration_ms":106600,"album":{"name":"Arden, David - Contemporary Piano Music: Piano Music By Part, Gorecki And Ustvolskaya"},"score":79,"artists":[{"name":"David Arden"}],"release_date":"1995-05-23"},{"play_offset_ms":15940,"external_ids":{},"artists":[{"name":"Jeanette Alexander"}],"result_from":1,"acrid":"4ac7a6491e5cd2c0dd1e6be7c07faa88","title":"Don't Ask Me Why","duration_ms":239060,"album":{"name":"Still Point"},"score":76,"external_metadata":{},"release_date":"1999-01-01"},{"label":"Jonathan Reed","play_offset_ms":116300,"external_ids":{},"result_from":1,"external_metadata":{},"acrid":"a93a960fc8fcac47f6af0b15efd5f1af","title":"Abide With Me \/ Jesus, I Am Resting, Resting","duration_ms":245613,"album":{"name":"Prayer Hymns"},"score":73,"artists":[{"name":"Jonathan Reed"}],"release_date":"2015-08-19"}]},"cost_time":2.1119999885559,"status":{"msg":"Success","version":"1.0","code":0},"result_type":0}`
			listenResponse := &models.ListenResponse{}
			err = json.Unmarshal(out, listenResponse)
			if err != nil {
				responses.NewInternalError(err.Error(), writer)
				return
			}
			if len(listenResponse.Metadata.Music) == 0 {
				responses.NewNotFoundError("Could not find album", writer)
				return
			}
			discogsAlbums, _, err := h.discogs.Search().AlbumWithArtist(listenResponse.Metadata.Music[0].AlbumName.Name, listenResponse.Metadata.Music[0].Artists[0].Name, nil)

			if err != nil || len(discogsAlbums) == 0 {
				responses.NewNotFoundError("could not find album in discogs", writer)
				return
			}

			discogsAlbums[0].ListenResponse = listenResponse
			existingAlbum, err := h.albumService.GetByDiscogsId(discogsAlbums[0].MasterId)
			var albumFound bool
			albumID := uint(0)
			if existingAlbum.ID > 0 {
				albumFound = true
				albumID = existingAlbum.ID
			}

			listen := &responses.Listen{
				AlbumFound:  albumFound,
				AlbumID:     albumID,
				AlbumTitle:  discogsAlbums[0].Title,
				Image:       discogsAlbums[0].CoverImage.String(),
				SongTitle:   listenResponse.Metadata.Music[0].Title,
				ArtistTitle: listenResponse.Metadata.Music[0].Artists[0].Name,
				DiscogsId:   discogsAlbums[0].MasterId,
			}
			responses.WriteResponse(200, listen, writer)
			return
		}
	}
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 24)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
