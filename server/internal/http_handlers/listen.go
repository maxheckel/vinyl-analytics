package http_handlers

import (
	"app/internal/models"
	"app/internal/responses"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func (h *handlers) Listen(writer http.ResponseWriter, request *http.Request){
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
			listenResponse := &models.ListenResponse{}
			err = json.Unmarshal(out, listenResponse)
			if err != nil {
				responses.NewInternalError(err.Error(), writer)
				return
			}
			responses.WriteResponse(200, listenResponse, writer)
			go os.Remove(fileName)
		}
	}
}

//
//func GetListenInfo(filename string){
//	hmac.New(sha1.New(), )
//}

func randomString() string{
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 24)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
