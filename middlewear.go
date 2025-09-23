package main

import (
	"fmt"
	"net/http"
	"webscrapper/internal/auth"
	"webscrapper/internal/database"
)

type apiauthhandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewearauth(handler apiauthhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apikey, err := auth.Getapikey(r.Header)
		if err != nil {
			handlingerrorwithjson(w, 403, fmt.Sprintf("couldnt read header %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			handlingerrorwithjson(w, 403, fmt.Sprintf("couldnt get user %s", err))
			return
		}
		handler(w, r, user)
	}

}
