package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
}

func getItems(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    items := []Item{
        {ID: "1", Name: "Coffee", Price: 3.50},
		{ID: "2", Name: "Tea", Price: 2.50},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func main() {
    http.HandleFunc("/api/items", getItems)
    http.ListenAndServe(":8000", nil)

    fmt.Println("Server running at http://localhost:8000/api/items")

    log.Println("Server running at http://localhost:8000/api/items")
    log.Fatal(http.ListenAndServe(":8000", nil))
}