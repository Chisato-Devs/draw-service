package v1

import (
	"chisato-draw-service/server/controllers"
	"encoding/json"
	"net/http"
)

var initLogger = controllers.Logger()

type Connected struct {
	Connected bool `json:"connected"`
}

func Init() {
	initLogger.Println("Starting...")

	http.HandleFunc("/v1/draw", DrawGetRequest)
	http.HandleFunc("/v1/stats", StatsHandler)
	http.HandleFunc("/v1/status", func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, err := json.Marshal(Connected{Connected: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonBytes)
		if err != nil {
			return
		}
	})

	initLogger.Println("Server started.")
	defer initLogger.Println("Server stopped.")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
