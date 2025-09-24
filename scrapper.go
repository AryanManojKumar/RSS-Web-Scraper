package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"
	"webscrapper/internal/database"

	"github.com/google/uuid"
)

func scrappig(dbquerire *database.Queries, concurenccy int, timeinterval time.Duration) {

	log.Printf("started scrapping  on  %v go routines every %s duration", concurenccy, timeinterval)
	ticker := time.NewTicker(timeinterval)

	for ; ; <-ticker.C {
		feeds, err := dbquerire.Getnextfeedtofetch(context.Background(), int32(concurenccy))
		if err != nil {
			log.Println("error feetching feeds")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go fetchingfeed(dbquerire, wg, feed)

		}
		wg.Wait()

	}

}

func fetchingfeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched ", err)
		return

	}

	rssfeed, err := Urltofeed(feed.Url)
	if err != nil {
		log.Println("couldnt feetch the feed", err)
	}

	for _, iteams := range rssfeed.Channel.Items {
		description := sql.NullString{}
		if iteams.Description != "" {
			description.String = iteams.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, iteams.PubDate)
		if err != nil {
			log.Printf("couldnt parse date %v with err %v", iteams.PubDate, err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       iteams.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         iteams.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post", err)
		}

	}
	log.Printf("feed %s collected , %v post found ", feed.Name, len(rssfeed.Channel.Items))

	defer wg.Done()
}
