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
const AuthURL string = config.AuthURL

//token
const AuthToken string = config.AuthPrefix + config.AuthToken

//Service Data Type
const AuthFormat string = "JSON"

//API Interface
var AuthRoutes = arbor.RouteCollection{
	arbor.Route{
		"NewSession",
		"POST",
		"/session",
		NewSession,
	},
	arbor.Route{
		"EndUsersSessions",
		"DELETE",
		"/session?username={username}",
		EndUsersSessions,
	},
	arbor.Route{
		"GetAuthenticatedUser",
		"GET",
		"/session/{token}",
		GetAuthenticatedUser,
	},
	arbor.Route{
		"ValidateSession",
		"POST",
		"/session/{token}",
		ValidateSession,
	},
	arbor.Route{
		"EndSession",
		"DELETE",
		"/session/{token}",
		EndSession,
	},
}

// arbor.Route handler
// w = writer, r = reader
func NewSession(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func EndUsersSessions(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func EndSession(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}
