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

type Feedfollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	feedID    uuid.UUID `json:"feed_id"`
}

func databasefeedfollowtofeedfollow(dbfeedfollow database.FeedFollow) Feedfollow {
	return Feedfollow{
		ID:        dbfeedfollow.ID,
		CreatedAt: dbfeedfollow.CreatedAt,
		UpdatedAt: dbfeedfollow.UpdatedAt,
		UserID:    dbfeedfollow.UserID,
		feedID:    dbfeedfollow.FeedID,
	}

}

func databasefftoff(dbffs []database.FeedFollow) []Feedfollow {

	feedsfollows := []Feedfollow{}

	for _, dbsfeedfollow := range dbffs {
		feedsfollows = append(feedsfollows, databasefeedfollowtofeedfollow(dbsfeedfollow))
	}
	return feedsfollows

}

type post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databaseposttopost(ab database.Post) post {

	var description *string
	if ab.Description.Valid {
		description = &ab.Description.String
	}

	return post{
		ID:          ab.ID,
		CreatedAt:   ab.CreatedAt,
		UpdatedAt:   ab.UpdatedAt,
		Title:       ab.Title,
		Description: description,
		PublishedAt: ab.PublishedAt,
		Url:         ab.Url,
		FeedID:      ab.FeedID,
	}
}

func databsePostsToPost(dbpost []database.Post) []post {
	post := []post{}

	for _, dbdbpost := range dbpost {
		post = append(post, databaseposttopost(dbdbpost))
	}
	return post

}
