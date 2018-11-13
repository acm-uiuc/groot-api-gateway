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

	"github.com/acm-uiuc/groot-api-gateway/config"
	"github.com/arbor-dev/arbor"
)

//Location
const ADURL string = config.ADURL

//Service Data Type
const ADFormat string = "JSON"

//API Interface
var ADRoutes = arbor.RouteCollection{
	arbor.Route{
		"AddUser",
		"POST",
		"/activedirectory/add",
		AddUser,
	},
}

// arbor.Route handler
func AddUser(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, ADURL+r.URL.String(), ADFormat, "", r)
}
