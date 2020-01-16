package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		profile := Profile{"SuperLee"}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	})
	http.ListenAndServe(":9000", nil)

}
