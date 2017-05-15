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
const UsersURL string = config.UsersURL

//Service Data Type
const UsersFormat string = "JSON"

//API Interface
var UsersRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetUsers",
		"GET",
		"/users",
		GetUsers,
	},
	arbor.Route{
		"IsMember",
		"GET",
		"/users/{netid}/is_member",
		IsMember,
	},
	arbor.Route{
		"NewUser",
		"POST",
		"/users",
		NewUser,
	},
	arbor.Route{
		"ConfirmUser",
		"PUT",
		"/users/{user_id}/paid",
		ConfirmUser,
	},
	arbor.Route{
		"DeleteUser",
		"DELETE",
		"/users/{user_id}",
		DeleteUser,
	},
	arbor.Route{
		"UsersLogin",
		"POST",
		"/users/login",
		UsersLogin,
	},
	arbor.Route{
		"UsersLogout",
		"POST",
		"/users/logout",
		UsersLogout,
	},
}

// arbor.Route handler
func GetUsers(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func IsMember(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func ConfirmUser(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func UsersLogin(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func UsersLogout(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}
