package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		timeZoneString := r.URL.Query().Get("tz")
		timeZones := strings.Split(timeZoneString, ",")
		var response map[string]string = make(map[string]string)
		if timeZones[0] != "" {
			for _, tz := range timeZones {
				loc, err := time.LoadLocation(string(tz))
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				response[tz] = time.Now().In(loc).String()
			}
		} else {
			response["current_time"] = time.Now().UTC().String()
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
