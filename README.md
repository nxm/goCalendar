# 📅 goCalendar - Open Source iCalendar Library

**goCalendar** is simply golang library for generating **iCalendar** format 

You can find more information in [specification](https://icalendar.org/RFC-Specifications/iCalendar-RFC-5545/)
### Usage
```go
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
```


Written under [MIT License](https://en.wikipedia.org/wiki/MIT_License)
