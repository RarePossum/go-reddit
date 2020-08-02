package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vartanbeno/go-reddit"
)

var ctx = context.Background()

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	withCredentials := reddit.WithCredentials("id", "secret", "username", "password")

	client, err := reddit.NewClient(nil, withCredentials)
	if err != nil {
		return
	}

	post, _, err := client.Post.SubmitText(ctx, reddit.SubmitTextOptions{
		Subreddit: "test",
		Title:     "This is a title",
		Text:      "This is some text",
	})
	if err != nil {
		return
	}

	fmt.Printf("The text post is available at: %s\n", post.URL)

	post, _, err = client.Post.SubmitLink(ctx, reddit.SubmitLinkOptions{
		Subreddit: "test",
		Title:     "This is a title",
		URL:       "http://example.com",
		Resubmit:  true,
	})
	if err != nil {
		return
	}

	fmt.Printf("The link post is available at: %s\n", post.URL)

	return
}
