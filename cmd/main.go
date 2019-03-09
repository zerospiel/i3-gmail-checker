package main

import (
	"fmt"
	"log"

	fetcher "github.com/zerospiel/i3-gmail-checker/pkg"
)

func main() {
	gs, err := fetcher.NewGmailService()
	if err != nil {
		log.Fatalf("cannot create gmail service: %s\n", err)
	}

	cnt, err := gs.FetchUnread()
	if err != nil {
		log.Fatalf("cannot fetch unread count with err: %s\n", err)
	}

	fmt.Println(cnt)
}
