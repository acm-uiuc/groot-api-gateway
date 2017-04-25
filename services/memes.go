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
const MemeURL string = "http://localhost:42069"

//Service Data Type
const MemeFormat string = "JSON"

//API Interface
var MemeRoutes = arbor.RouteCollection{
	arbor.Route{
		"ListMemes",
		"GET",
		"/memes",
		ListMemes,
	},
	arbor.Route{
		"NewMeme",
		"POST",
		"/memes",
		NewMeme,
	},
	arbor.Route{
		"MemeInfo",
		"GET",
		"/memes/{meme_id}",
		MemeInfo,
	},
	arbor.Route{
		"DeleteMeme",
		"DELETE",
		"/memes/{meme_id}",
		DeleteMeme,
	},
	arbor.Route{
		"ApproveMeme",
		"PUT",
		"/memes/{meme_id}/approve",
		ApproveMeme,
	},
	arbor.Route{
		"CastMemeVote",
		"PUT",
		"/memes/{meme_id}/vote",
		CastMemeVote,
	},
	arbor.Route{
		"DeleteMemeVote",
		"DELETE",
		"/memes/{meme_id}/vote",
		DeleteMemeVote,
	},
	arbor.Route{
		"GetRandomMeme",
		"GET",
		"/memes/random",
		GetRandomMeme,
	},
}

// services.Route handler
func ListMemes(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func NewMeme(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func MemeInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func DeleteMeme(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func ApproveMeme(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func CastMemeVote(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func DeleteMemeVote(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func GetRandomMeme(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}
