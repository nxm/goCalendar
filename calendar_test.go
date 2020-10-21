package goCalendar

import (
	"fmt"
	"testing"
	"time"
)

func TestApp(*testing.T) {
	fmt.Println("Starting app...")
	//creating new Calendar
	cal := createCalendar()

	//creating new Event
	start := time.Now().Add(time.Hour) //start of event
	end := start.Add(time.Minute*45) //end of event

	simpleEvent := createEvent(start, end)
	simpleEvent.setTitle("Music lesson") //title
	simpleEvent.setDescription("Learning musical notes") //description
	simpleEvent.setLocation("School") //location

	//adding Event to Calendar
	cal.addEvent(simpleEvent)

	//saving Calendar with Events to file
	err := cal.save("example.ics")

	if err != nil {
		fmt.Print(err)
	}
}