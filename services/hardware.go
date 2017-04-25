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
const HardwareURL string = "http://localhost:4523/api/v1.0"

//Service Data Type
const HardwareFormat string = "JSON"

//API Interface
var HardwareRoutes = arbor.RouteCollection{
	arbor.Route{
		"NewItem",
		"POST",
		"/items/{id}",
		NewSession,
	},
	arbor.Route{
		"UpdateItem",
		"PUT",
		"/items/{id}",
		EndUsersSessions,
	},
	arbor.Route{
		"DeleteItem",
		"DELETE",
		"/item/{id}",
		GetAuthenticatedUser,
	},
}

// services.Route handler
func NewItem(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}
