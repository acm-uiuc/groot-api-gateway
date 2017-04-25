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
const GigsURL string = config.GigsURL

//Service Data Type
const GigFormat string = "JSON"

//API Interface
var GigsRoutes = arbor.RouteCollection{
	arbor.Route{
		"ListGigs",
		"GET",
		"/gigs",
		ListGigs,
	},
	arbor.Route{
		"NewGig",
		"POST",
		"/gigs",
		NewGig,
	},
	arbor.Route{
		"GigInfo",
		"GET",
		"/gigs/{gig_id}",
		GigInfo,
	},
	arbor.Route{
		"DeactivateGig",
		"PUT",
		"/gigs/{gig_id}",
		DeactivateGig,
	},
	arbor.Route{
		"DeleteGig",
		"DELETE",
		"/gigs/{gig_id}",
		DeleteGig,
	},
	arbor.Route{
		"ListClaims",
		"GET",
		"/gigs/claims",
		ListClaims,
	},
	arbor.Route{
		"CreateClaim",
		"POST",
		"/gigs/claims",
		CreateClaim,
	},
	arbor.Route{
		"ListSingleClaim",
		"GET",
		"/gigs/claims/{claim_id}",
		ListSingleClaim,
	},
	arbor.Route{
		"AcceptClaim",
		"PUT",
		"/gigs/claims/{claim_id}",
		AcceptClaim,
	},
	arbor.Route{
		"DeleteClaim",
		"DELETE",
		"/gigs/claims/{claim_id}",
		DeleteClaim,
	},
}

//Route handler
func ListGigs(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func NewGig(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func GigInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func CreateClaim(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func DeactivateGig(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func DeleteGig(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func ListClaims(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func ListSingleClaim(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func AcceptClaim(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func DeleteClaim(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, GigsURL+r.URL.String(), GigFormat, "", r)
}
