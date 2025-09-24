package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webscrapper/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleuserdb(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decod := json.NewDecoder(r.Body)
	para := parameter{}

	err := decod.Decode(&para)
	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("prasing Error %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{

		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      para.Name,
	})

	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("COuldnt make user %s", err))
		return
	}

	responsewithjson(w, 201, databaseUsertoUser(user))

}

func (apiCfg *apiConfig) handlerGetUserFromApi(w http.ResponseWriter, r *http.Request, user database.User) {

	responsewithjson(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerGetUserpostforuser(w http.ResponseWriter, r *http.Request, user database.User) {

	userpost, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10})

	if err != nil {
		handlingerrorwithjson(w, 400, fmt.Sprintf("couldnt get post %v", err))
		return
	}
	responsewithjson(w, 200, databsePostsToPost(userpost))
}
