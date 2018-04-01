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
const NotificationURL string = config.NotificationURL

//Service Data Type
const NotificationFormat string = "JSON"

//API Interface
var NotificationRoutes = arbor.RouteCollection{
	arbor.Route{
		"PostNotification",
		"POST",
		"/notification",
		PostNotification,
	},
}

// arbor.Route handler
func PostNotification(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, NotificationURL+r.URL.String(), NotificationFormat, "", r)
}
