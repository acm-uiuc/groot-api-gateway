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
	"fmt"
	"net/http"

	"github.com/arbor-dev/arbor"
)

var Routes = arbor.RouteCollection{
	arbor.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I AM GROOT!\n")
}

func RegisterAPIs() arbor.RouteCollection {
	Routes = append(Routes, ADRoutes...)
	Routes = append(Routes, AuthRoutes...)
	Routes = append(Routes, CreditsRoutes...)
	Routes = append(Routes, EventsRoutes...)
	Routes = append(Routes, GigsRoutes...)
	Routes = append(Routes, GroupsRoutes...)
	Routes = append(Routes, HardwareRoutes...)
	Routes = append(Routes, MemeRoutes...)
	Routes = append(Routes, MerchRoutes...)
	Routes = append(Routes, NotificationRoutes...)
	Routes = append(Routes, QuotesRoutes...)
	Routes = append(Routes, RecruitersRoutes...)
	Routes = append(Routes, UsersRoutes...)
	return Routes
}
