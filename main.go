package main

import (
	"net/http"

	"github.com/kale-swapnil/azure-ai-image-processing/functions"
)

func main() {
	http.HandleFunc("/api/EventGridTrigger", functions.HandleEvent)
	http.ListenAndServe(":8000", nil)
}
