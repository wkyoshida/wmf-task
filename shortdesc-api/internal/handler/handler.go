package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"shortdesc-api/internal/config"
	"shortdesc-api/internal/mediawiki"
)

// Response is a struct to hold the response of the shortdesc-api.
type Response struct {
	Pages []mediawiki.PageShortDesc `json:"pages"`
}

// HandleRequests handles the HTTP requests for the shortdesc-api.
func HandleRequests() {
	http.HandleFunc("/shortdesc", ServeHTTP)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Env.ApiPort), nil))
}

// ServeHTTP serves the HTTP requests for the shortdesc-api.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		titles, queryPresent := query["titles"]
		if !queryPresent || len(titles) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Bad request. Query parameter 'titles' not present."}`))
			return
		}

		pages, err := mediawiki.GetShortDescs(titles)
		if err != nil {
			internalServerError(w, r)
			return
		}

		response := Response{
			Pages: pages,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			internalServerError(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return
	default:
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Can't find method requested."}`))
		return
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "Internal server error."}`))
}
