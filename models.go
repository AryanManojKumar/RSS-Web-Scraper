package main

import (
	"time"
	"webscrapper/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

func databaseUsertoUser(dbUser database.User) User {

	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}

}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"feedurl"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbfeed database.Feed) Feed {

	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}

}

func databaseFeedstoFeeds(dbfeeds []database.Feed) []Feed {

	feeds := []Feed{}

	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbfeed))
	}
	return feeds

}
