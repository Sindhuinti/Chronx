package pkg

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func DeleteEvent(title string) {
	client, err := GetClient()
	if err != nil {
		return
	}
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}


	events, err := srv.Events.List("primary").Q(title).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, event := range events.Items {

		err := srv.Events.Delete("primary", event.Id).Do()
		if err != nil {
			log.Fatalf("Unable to delete event: %v", err)
		}

		fmt.Printf("Event '%s' deleted.\n", event.Summary)
	}
}