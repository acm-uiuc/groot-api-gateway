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
const RequestURL string = config.RequestURL

//Service Data Type
const RequestFormat string = "JSON"

//API Interface
var RequestRoutes = arbor.RouteCollection{
    arbor.Route{
        "PostRequest",
        "POST",
        "/request",
        PostRequest,
    },
    arbor.Route{
        "GetRequest",
        "GET",
        "/request",
        GetRequest,
    },
    arbor.Route{
        "GetRequestByID",
        "GET",
        "/request/{id}",
        GetRequestByID,
    },
    arbor.Route{
        "UpdateRequest",
        "PUT",
        "/request/{id}",
        UpdateRequest,
    },
}

// arbor.Route handler
func PostRequest(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RequestURL+r.URL.String(), RequestFormat, "", r)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RequestURL+r.URL.String(), RequestFormat, "", r)
}

func GetRequestByID(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RequestURL+r.URL.String(), RequestFormat, "", r)
}

func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RequestURL+r.URL.String(), RequestFormat, "", r)
}
