/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package services

import (
	"net/http"

	"github.com/acm-uiuc/arbor"
)

//Location
const EventsURL string = "http://localhost:8002"

//Service Data Type
const EventsFormat string = "JSON"

//API Interface
var EventsRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetEvents",
		"GET",
		"/events",
		GetEvents,
	},
	arbor.Route{
		"GetUpcomingEvents",
		"GET",
		"/events/upcoming",
		GetUpcomingEvents,
	},
}

// services.Route handler
func GetEvents(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, EventsURL+r.URL.String(), EventsFormat, "", r)
}

func GetUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, EventsURL+r.URL.String(), EventsFormat, "", r)
}
