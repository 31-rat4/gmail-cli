package main

import (
	"context"
	"fmt"
	"github.com/31-rat4/gmail-cli/cmd/ascess"
	"log"
)

func main() {
	ctx := context.Background()
	AscessHandler := ascess.NewAscessHandler(ctx)
	srv := AscessHandler.GenGmailClient()
	fmt.Println(srv)
	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}
