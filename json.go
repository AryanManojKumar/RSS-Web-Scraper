package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlingerrorwithjson(w http.ResponseWriter, code int, msg string) {

	//because i think code below 499 are client side and doesnt affect us
	if code > 499 {
		log.Println("responding with 500< error", msg)

	}
	type errmsg struct {
		Errmsg string `json:"error"`
	}
	responsewithjson(w, code, errmsg{Errmsg: msg})

}

func responsewithjson(w http.ResponseWriter, code int, playload interface{}) {

	data, err := json.Marshal(playload)
	if err != nil {
		log.Printf("problem with the payload %v", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
