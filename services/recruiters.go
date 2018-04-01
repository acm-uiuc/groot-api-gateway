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
const RecruiterURL string = config.RecruiterURL

//Service Data Type
const RecruiterFormat string = "JSON"

//API Interface
var RecruitersRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetJobs",
		"GET",
		"/jobs",
		GetJobs,
	},
	arbor.Route{
		"CreateJob",
		"POST",
		"/jobs",
		CreateJob,
	},
	arbor.Route{
		"ApproveJob",
		"PUT",
		"/jobs/{job_id}/approve",
		ApproveJob,
	},
	arbor.Route{
		"DeleteJob",
		"DELETE",
		"/jobs/{job_id}",
		DeleteJob,
	},
	arbor.Route{
		"GetRecruiters",
		"GET",
		"/recruiters",
		GetRecruiters,
	},
	arbor.Route{
		"RecruiterLogin",
		"POST",
		"/recruiters/login",
		RecruiterLogin,
	},
	arbor.Route{
		"CreateRecruiter",
		"POST",
		"/recruiters",
		CreateRecruiter,
	},
	arbor.Route{
		"GetRecruiter",
		"GET",
		"/recruiters/{recruiter_id}",
		GetRecruiter,
	},
	arbor.Route{
		"RenewRecruiter",
		"PUT",
		"/recruiters/{recruiter_id}/renew",
		RenewRecruiter,
	},
	arbor.Route{
		"ResetAllRecruiters",
		"POST",
		"/recruiters/reset",
		ResetAllRecruiters,
	},
	arbor.Route{
		"UpdateRecruiter",
		"PUT",
		"/recruiters/{recruiter_id}",
		UpdateRecruiter,
	},
	arbor.Route{
		"ResetRecruiter",
		"POST",
		"/recruiters/reset_password",
		ResetRecruiter,
	},
	arbor.Route{
		"GetRecruiterInvite",
		"GET",
		"/recruiters/{recruiter_id}/invite",
		GetRecruiterInvite,
	},
	arbor.Route{
		"SendRecruiterInvite",
		"POST",
		"/recruiters/{recruiter_id}/invite",
		SendRecruiterInvite,
	},
	arbor.Route{
		"ResetRecruiterInvite",
		"POST",
		"/recruiters/invite",
		ResetRecruiterInvite,
	},
	arbor.Route{
		"DeleteRecruiter",
		"DELETE",
		"/recruiters/{recruiter_id}",
		DeleteRecruiter,
	},
	arbor.Route{
		"GetStudents",
		"GET",
		"/students",
		GetStudents,
	},
	arbor.Route{
		"CreateStudent",
		"POST",
		"/students",
		CreateStudent,
	},
	arbor.Route{
		"ApproveStudent",
		"PUT",
		"/students/{netid}/approve",
		ApproveStudent,
	},
	arbor.Route{
		"GetStudent",
		"GET",
		"/students/{netid}",
		GetStudent,
	},
	arbor.Route{
		"DeleteStudent",
		"DELETE",
		"/students/{netid}",
		DeleteStudent,
	},
}

// arbor.Route handler
func GetJobs(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func ApproveJob(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func GetRecruiters(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func RecruiterLogin(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func GetRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func CreateRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func RenewRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func ResetAllRecruiters(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func UpdateRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func ResetRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func GetRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func SendRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func ResetRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func DeleteRecruiter(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func ApproveStudent(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, "", r)
}
