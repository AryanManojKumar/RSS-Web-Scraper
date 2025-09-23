package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webscrapper/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleuserfeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decod := json.NewDecoder(r.Body)
	para := parameter{}

	err := decod.Decode(&para)
	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("prasing Error %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{

		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      para.Name,
		Url:       para.Url,
		UserID:    user.ID,
	})

	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt create feed %s", err))
		return
	}

	responsewithjson(w, 201, databaseFeedtoFeed(feed))

}
