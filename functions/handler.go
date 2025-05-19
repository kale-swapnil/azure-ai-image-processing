package functions

import (
	"encoding/json"
	"net/http"
)

type Event struct {
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var events []Event
	_ = json.NewDecoder(r.Body).Decode((&events))

	imageUrl := events[0].Data.URL
	tags, err := ClassifyImage(imageUrl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_ = Save(imageUrl, tags)
	w.Write([]byte("Image processed"))

}
