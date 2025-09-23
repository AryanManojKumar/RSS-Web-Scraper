package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webscrapper/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleuserfeedfollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decod := json.NewDecoder(r.Body)
	para := parameter{}

	err := decod.Decode(&para)
	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("prasing Error %s", err))
		return
	}

	Feedfollow, err := apiCfg.DB.Createfeedfollow(r.Context(), database.CreatefeedfollowParams{

		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    para.FeedID,
	})

	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt create feed follow %s", err))
		return
	}

	responsewithjson(w, 201, databasefeedfollowtofeedfollow(Feedfollow))

}

func (apiCfg *apiConfig) handleuserfeedfollows(w http.ResponseWriter, r *http.Request, user database.User) {

	Feedfollows, err := apiCfg.DB.GetAllFeedfollows(r.Context(), user.ID)
	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt get feed follow %s", err))
		return
	}

	responsewithjson(w, 201, databasefftoff(Feedfollows))

}

func (apiCfg *apiConfig) handledeletefeedfollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollowidstrg := chi.URLParam(r, "feedfollowid")

	feedfollowid, err := uuid.Parse(feedfollowidstrg)
	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt get feed id %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})

	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt delete feed follow %s", err))
		return
	}
	responsewithjson(w, 200, struct{}{})
}
