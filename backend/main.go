package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"net/http"
)

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/items", itemsHandler)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	items := []Item{
		{
			ID:    1,
			Title: "Элемент 1",
		},
		{
			ID:    2,
			Title: "Элемент 2",
		},
		{
			ID:    3,
			Title: "Элемент 3",
		},
	}

	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
	}
}
