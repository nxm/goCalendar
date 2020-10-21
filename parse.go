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
 */
package goCalendar

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strconv"
)

func (cal VCalendar) ParseVCalendar() []string {
	var strs []string

	s := reflect.ValueOf(&cal).Elem()
	typeOfT := s.Type()

	strs = append(strs, "BEGIN:VCALENDAR\n")

	for i := 0; i < s.NumField(); i++ {
		var value string
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())

		//shitty fix, sorry for that
		if typeOfT.Field(i).Type.String() == "float32" {
			value = fmt.Sprintf("%f", f.Interface().(float32))
		} else if typeOfT.Field(i).Type.String() == "bool"{
			value = strconv.FormatBool(f.Interface().(bool))
		} else {
			value = f.Interface().(string)
		}

		strs = append(strs, typeOfT.Field(i).Name + ":" + value + "\n")

	}

	for _, evn := range cal.getEvents() {
		tempEvent := cal.ParseVEvent(evn)
		strs = append(strs, tempEvent...)
	}

	strs = append(strs, "END:VCALENDAR")

	return strs
}

func (VCalendar) ParseVEvent(ev VEvent) []string {
	var strs []string

	//i cant cast parameter from interface{} to VEvent cuz
	//reflect: call of reflect.Value.NumField on interface Value
	s := reflect.ValueOf(&ev).Elem()
	typeOfT := s.Type()

	//beginning event
	strs = append(strs, "BEGIN:VEVENT\n")

	for i := 0; i < s.NumField(); i++ {
		var value string
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())

		//shitty fix, rly sorry
		if typeOfT.Field(i).Type.String() == "uuid.UUID" {
			value = f.Interface().(uuid.UUID).String()
		} else {
			value = f.Interface().(string)
		}

		//panic: interface conversion: interface {} is uuid.UUID, not string
		//strs = append(strs, typeOfT.Field(i).Name + ":" + f.Interface().(string) + "\n")

		strs = append(strs, typeOfT.Field(i).Name + ":" + value + "\n")
	}

	//ending event
	strs = append(strs, "END:VEVENT\n")

	return strs
}
