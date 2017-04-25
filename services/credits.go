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
const CreditsURL string = "http://localhost:8765"

//Service Data Type
const CreditsFormat string = "JSON"

//API Interface
var CreditsRoutes = arbor.RouteCollection{
	arbor.Route{
		"NewPayment",
		"POST",
		"/payment",
		NewPayment,
	},
	arbor.Route{
		"GetCreditsUser",
		"GET",
		"/credits/users/{netid}",
		GetCreditsUser,
	},
	arbor.Route{
		"GetCreditsUserMultiple",
		"GET",
		"/credits/users",
		GetCreditsUserMultiple,
	},
	arbor.Route{
		"GetTransactions",
		"GET",
		"/credits/transactions",
		GetTransactions,
	},
	arbor.Route{
		"CreateTransaction",
		"POST",
		"/credits/transactions",
		CreateTransaction,
	},
	arbor.Route{
		"DeleteTransaction",
		"DELETE",
		"/credits/transactions/{id}",
		DeleteTransaction,
	},
}

// services.Route handler
func NewPayment(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUser(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUserMultiple(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}
