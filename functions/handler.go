package handler

import (
	"azure-ai-image-processing/functions/cosmos"
	"azure-ai-image-processing/functions/vision"
	"encoding/json"
	"net/http"
)

type Event struct {
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	var events []Event
	_ = json.NewDecoder(r.Body).Decode((&events))

	imageUrl := events[0].Data.URL
	tags, err := vision.ClassifyImage(imageUrl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_ = cosmos.Save(imageUrl, tags)
	w.Write([]byte("Image processed"))

}
