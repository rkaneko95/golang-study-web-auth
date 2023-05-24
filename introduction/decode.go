package introduction

import (
	"encoding/json"
	"log"
	"net/http"
)

func Bar(w http.ResponseWriter, r *http.Request) {
	var p1 Person

	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Encodes bad data", err)
	}

	log.Printf("Person: %v", p1)
}
