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
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		fmt.Println("Upcoming events (today):")
		for _, event := range events.Items {
			start, err := time.Parse(time.RFC3339, event.Start.DateTime)
			if err != nil {
				log.Fatalf("Unable to parse start time: %v", err)
			}

			fmt.Printf("%v - %s\n", event.Summary, start.Format("15:04"))
		}
	}
}

func GetOneEvent(list string) {
	client, err := GetClient()
	if err != nil {
		return
	}
	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}

	eventTitle := list

	events, err := calendarService.Events.List("primary").Q(eventTitle).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, event := range events.Items {
		fmt.Println("Title: ", event.Summary)
		if event.Description != "" {

			fmt.Println("- Description: ", event.Description)
		}
		start, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		fmt.Println("- Start time: ", start.Format("15:04"))
		if event.HangoutLink != "" {
			fmt.Println("- MeetLink: ", event.HangoutLink)
		}
		if event.Location != "" {

			fmt.Println("- Location: ", event.Location)
		}

	}
}
