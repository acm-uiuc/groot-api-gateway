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
const QuotesURL string = config.QuotesURL

//Service Data Type
const QuoteFormat string = "JSON"

//API Interface
var QuotesRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetAllQuotes",
		"GET",
		"/quotes",
		GetAllQuotes,
	},
	arbor.Route{
		"DeleteQuote",
		"DELETE",
		"/quotes/{id}",
		DeleteQuote,
	},
	arbor.Route{
		"GetQuote",
		"GET",
		"/quotes/{id}",
		GetQuote,
	},
	arbor.Route{
		"CastVote",
		"POST",
		"/quotes/{id}/vote",
		CastVote,
	},
	arbor.Route{
		"ApproveQuote",
		"PUT",
		"/quotes/{id}/approve",
		ApproveQuote,
	},
	arbor.Route{
		"DeleteVote",
		"DELETE",
		"/quotes/{id}/vote",
		DeleteVote,
	},
	arbor.Route{
		"CreateQuote",
		"POST",
		"/quotes",
		CreateQuote,
	},
	arbor.Route{
		"UpdateQuote",
		"PUT",
		"/quotes/{id}",
		UpdateQuote,
	},
}

// arbor.Route handler
func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func CastVote(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func ApproveQuote(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func DeleteVote(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}

func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, QuotesURL+r.URL.String(), QuoteFormat, "", r)
}
