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
const GroupsURL string = "http://localhost:9001"

//Service Data Type
const GroupsFormat string = "JSON"

//API Interface
var GroupsRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetGroupTypes",
		"GET",
		"/groups",
		GetGroupTypes,
	},
	arbor.Route{
		"GetGroups",
		"GET",
		"/groups/{groupType}",
		GetGroups,
	},
	arbor.Route{
		"GetGroup",
		"GET",
		"/groups/{groupType}/{groupName}",
		GetGroup,
	},
	arbor.Route{
		"IsGroupMember",
		"GET",
		"/groups/{groupType}/{groupName}?isMember={netid}",
		IsGroupMember,
	},
}

// arbor.Route handler
func GetGroupTypes(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GroupsURL+r.URL.String(), GroupsFormat, "", r)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GroupsURL+r.URL.String(), GroupsFormat, "", r)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GroupsURL+r.URL.String(), GroupsFormat, "", r)
}

func IsGroupMember(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, GroupsURL+r.URL.String(), GroupsFormat, "", r)
}
