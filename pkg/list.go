package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetEvents() {
	client, err := GetClient()
	if err != nil {
		return
	}
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}

	startTime := time.Now()
	endTime := time.Now().Add(time.Hour * 24)

	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(startTime.Format(time.RFC3339)).TimeMax(endTime.Format(time.RFC3339)).Do()
	if err !=nil{
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		fmt.Println("Upcoming events:")
		for _, event := range events.Items {
			start, err := time.Parse(time.RFC3339, event.Start.DateTime)
			if err != nil {
				log.Fatalf("Unable to parse start time: %v", err)
			}
			
    	
			fmt.Printf("%v - %s\n", event.Summary, start.Format("15:04"))
		}
	}
}