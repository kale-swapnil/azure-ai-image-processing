package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/EventGridTrigger", handler.handleEvent)
	http.ListenAndServe(":8000", nil)
}
