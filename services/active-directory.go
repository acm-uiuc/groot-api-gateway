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
	"github.com/acm-uiuc/groot-api-gateway/config"
)

//Location
const ADServiceURL string = config.ADServerURL

//ADToken
const ADToken string = config.ADToken

//Service Data Type
const ADServiceFormat string = "JSON"

//API Interface
var ADServiceRoutes = arbor.RouteCollection{
	arbor.Route{
		"UserPaid",
		"GET",
		"/activedirectory/add/{netid}",
		UserPaid,
	},
}

// arbor.Route handler

func UserPaid(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, ADServiceURL+r.URL.String(), ADServiceFormat, ADToken, r)
}
