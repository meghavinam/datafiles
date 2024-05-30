package main

import (
	"encoding/json"
	"lil/emojiserv/search"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", SearchData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SearchData(w http.ResponseWriter, r *http.Request) {
	var params search.Params
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	result := search.ByDescription(params)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
