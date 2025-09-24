package main

import (
	"context"
	"log"
	"sync"
	"time"
	"webscrapper/internal/database"
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
		log.Println("found post", iteams.Title, "one feed", feed.Name)

	}
	log.Printf("feed %s collected , %v post found ", feed.Name, len(rssfeed.Channel.Items))

	defer wg.Done()
}
