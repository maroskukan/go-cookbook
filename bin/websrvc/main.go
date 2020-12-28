package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Handle registration
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query Parameters e.g. localhost:3000/?name=Maros
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		//w.Write([]byte("Hello " + name))
		//Return output in json format
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
