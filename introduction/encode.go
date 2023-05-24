package introduction

import (
	"encoding/json"
	"log"
	"net/http"
)

func Foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encodes bad data", err)
	}
}
