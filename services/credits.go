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

	"github.com/acm-uiuc/arbor/proxy"
	"github.com/acm-uiuc/arbor/services"
	"github.com/acm-uiuc/groot-api-gateway/config"
)

//Location
const CreditsURL string = config.CreditsURL

//Service Data Type
const CreditsFormat string = "JSON"

//API Interface
var CreditsRoutes = services.RouteCollection{
	services.Route{
		"NewPayment",
		"POST",
		"/payment",
		NewPayment,
	},
	services.Route{
		"GetCreditsUser",
		"GET",
		"/credits/users/{netid}",
		GetCreditsUser,
	},
	services.Route{
		"GetCreditsUserMultiple",
		"GET",
		"/credits/users",
		GetCreditsUserMultiple,
	},
	services.Route{
		"GetTransactions",
		"GET",
		"/credits/transactions",
		GetTransactions,
	},
	services.Route{
		"CreateTransaction",
		"POST",
		"/credits/transactions",
		CreateTransaction,
	},
	services.Route{
		"DeleteTransaction",
		"DELETE",
		"/credits/transactions/{id}",
		DeleteTransaction,
	},
}

// services.Route handler
func NewPayment(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUser(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUserMultiple(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}
