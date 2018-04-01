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
const MerchURL string = config.MerchURL

//Service Data Type
const MerchFormat string = "JSON"

//API Interface
var MerchRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetMerchLocations",
		"GET",
		"/merch/locations",
		GetMerchLocations,
	},
	arbor.Route{
		"GetMerchItemAtLocation",
		"GET",
		"/merch/locations/{location}",
		GetMerchItemAtLocation,
	},
	arbor.Route{
		"ClearMerchItemAtLocation",
		"PUT",
		"/merch/locations/{location}",
		ClearMerchItemAtLocation,
	},
	arbor.Route{
		"GetMerchUsers",
		"GET",
		"/merch/users",
		GetMerchUsers,
	},
	arbor.Route{
		"GetMerchUserByNetid",
		"GET",
		"/merch/users/{netid}",
		GetMerchUserByNetid,
	},
	arbor.Route{
		"GetMerchUserByPin",
		"POST",
		"/merch/users/pins",
		GetMerchUserByPin,
	},
	arbor.Route{
		"CreateMerchTransaction",
		"POST",
		"/merch/transactions",
		CreateMerchTransaction,
	},
	arbor.Route{
		"GetItems",
		"GET",
		"/merch/items",
		GetItems,
	},
	arbor.Route{
		"CreateMerch",
		"POST",
		"/merch/items",
		CreateMerch,
	},
	arbor.Route{
		"GetMerch",
		"GET",
		"/merch/items/{id}",
		GetMerch,
	},
	arbor.Route{
		"UpdateMerch",
		"PUT",
		"/merch/items/{id}",
		UpdateMerch,
	},
	arbor.Route{
		"DeleteMerch",
		"DELETE",
		"/merch/items/{id}",
		DeleteMerch,
	},
}

//Route handler
func GetMerchLocations(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetMerchItemAtLocation(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func ClearMerchItemAtLocation(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetMerchUsers(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetMerchUserByNetid(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetMerchUserByPin(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func CreateMerchTransaction(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func CreateMerch(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func GetMerch(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func UpdateMerch(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}

func DeleteMerch(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, MerchURL+r.URL.String(), MerchFormat, "", r)
}
