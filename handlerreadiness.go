package main

import (
	"net/http"
)

func readiness(w http.ResponseWriter, r *http.Request) {
	responsewithjson(w, 200, struct{}{})

}
