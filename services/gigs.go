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
const GigsURL string = config.GigsURL

//Service Data Type
const GigFormat string = "JSON"

//API Interface
var GigsRoutes = services.RouteCollection{
    services.Route{
        "ListGigs",
        "GET",
        "/gigs",
        ListGigs,
    },
   services.Route{
        "NewGig",
        "POST",
        "/gigs",
        NewGig,
    },
   services.Route{
        "GigInfo",
        "GET",
        "/gigs/{gig_id}",
        GigInfo,
    },
   services.Route{
        "EditGig",
        "PUT",
        "/gigs/{gig_id}",
        EditGig,
    },
   services.Route{
        "DeleteGig",
        "DELETE",
        "/gigs/{gig_id}",
        DeleteGig,
    },
    services.Route{
        "ListClaims",
        "GET",
        "/gigs/claims",
        ListClaims,
    },
   services.Route{
        "CreateClaim",
        "POST",
        "/gigs/claims",
        CreateClaim,
    },
    services.Route{
        "ListSingleClaim",
        "GET",
        "/gigs/claims/{claim_id}",
        ListSingleClaim,
    },
    services.Route{
        "AcceptClaim",
        "PUT",
        "/gigs/claims/{claim_id}",
        AcceptClaim,
    },
    services.Route{
        "DeleteClaim",
        "DELETE",
        "/gigs/claims/{claim_id}",
        DeleteClaim,
    },
}

//Route handler
func ListGigs(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func NewGig(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func GigInfo(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func CreateClaim(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func EditGig(w http.ResponseWriter, r *http.Request) {
    proxy.PUT(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func DeleteGig(w http.ResponseWriter, r *http.Request) {
    proxy.DELETE(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func ListClaims(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func ListSingleClaim(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func AcceptClaim(w http.ResponseWriter, r *http.Request) {
    proxy.PUT(w, GigsURL+r.URL.String(), GigFormat, "", r)
}

func DeleteClaim(w http.ResponseWriter, r *http.Request) {
    proxy.DELETE(w, GigsURL+r.URL.String(), GigFormat, "", r)
}