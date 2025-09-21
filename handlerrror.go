package main

import (
	"net/http"
)

func errorhandler(w http.ResponseWriter, r *http.Request) {
	handlingerrorwithjson(w, 400, "something went horrible")

}
