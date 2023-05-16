package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func UpdateEvent() {
	client, err := GetClient()
	if err != nil {
		return
	}
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}


	newStartTime,err1 := time.Parse("2006-01-02 15:04:05", start)
	newEndTime,err2 := time.Parse("2006-01-02 15:04:05", end)

	if(err1 !=nil && err2 !=nil){
		fmt.Println("Please enter correct time format")
		return
	}

	events, err := srv.Events.List("primary").Q(title).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, event := range events.Items {

		event.Start = &calendar.EventDateTime{
			DateTime: newStartTime.Format(time.RFC3339),
		}
		event.End = &calendar.EventDateTime{
			DateTime: newEndTime.Format(time.RFC3339),
		}

		_, err := srv.Events.Update("primary", event.Id, event).Do()
		if err != nil {
			log.Fatalf("Unable to update event: %v", err)
		}

		fmt.Printf("Event '%s' updated.\n", event.Summary)
	}
}
