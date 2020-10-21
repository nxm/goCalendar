/**
 * MIT License
 *
 * Copyright (c) 2020 nxm
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
package goCalendar

import (
	"github.com/google/uuid"
	"time"
)

const (
	ProdId = "-//github.com//nxm"
)

var (
	CalendarEvents []VEvent
)

type VCalendar struct {
	ProdId   string  `str:"PRODID"`
	Version  float32 `str:"VERSION"`
	CalScale bool    `str:"CALSCALE"`
}

type VEvent struct {
	DTStamp     string    `str:"DTSTAMP"`
	DTStart     string     `str:"DTSTART"`
	DTEnd       string     `str:"DTEND"`
	Summary     string    `str:"SUMMARY"`
	UID         uuid.UUID `str:"UID"`
	Location    string    `str:"LOCATION"`
	Description string    `str:"DESCRIPTION"`
}

func (VCalendar) addEvent(event VEvent) {
	CalendarEvents = append(CalendarEvents, event)
}

func (VCalendar) getEvents() []VEvent {
	return CalendarEvents
}

func createCalendar() VCalendar {
	return VCalendar{ProdId: ProdId, Version: 2.0, CalScale: false}
}

func createEvent(start, end time.Time) VEvent {
	event := VEvent{
		DTStamp: ParseTimeToCalendar(time.Now()),
		DTStart: ParseTimeToCalendar(start),
		DTEnd:   ParseTimeToCalendar(end),
		UID: GenerateUID(),
	}

	return event
}

func (e *VEvent) setTitle(title string) {
	e.Summary = title
}

func (e *VEvent) setDescription(description string) {
	e.Description = description
}

func (e *VEvent) setLocation(location string) {
	e.Location = location
}
