package controllers

import (
	"net/http"
	"photos/dao"
	"photos/models"
	"photos/session"
	"strconv"
)

// Result asd
type result struct {
	Data    []models.Hospital
	Session *session.Session
}

// SQLHandler Sql控制器
func SQLHandler(w http.ResponseWriter, r *http.Request) {
	Authenticate(w, r)
	sess := session.GetSession(r)
	limit := "20"
	if isPresent(r, "limit") {
		limit = r.URL.Query()["limit"][0]
	}
	data := dao.GetAllHospital(limit)
	Templates["./pages/sql/sql.html"].Execute(w, &result{Data: data, Session: sess})
}

// SQLEditHandler Sql控制器
func SQLEditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := "1"
		params := r.URL.Query()["id"]
		if len(params) > 0 {
			id = r.URL.Query()["id"][0]
		}
		data := dao.GetHospitalByID(id)
		Templates["./pages/sql/edit.html"].Execute(w, data)
	} else {
		id, _ := strconv.ParseInt(r.PostFormValue("id"), 10, 0)
		dao.EditHospital(models.Hospital{ID: int(id), Name: r.PostFormValue("name"), Address: r.PostFormValue("address")})
		http.Redirect(w, r, "/sql", 303)
	}

}

// SQLAddHandler SQLAddHandler
func SQLAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Templates["./pages/sql/add.html"].Execute(w, nil)
	} else {
		dao.AddHospital(models.Hospital{Name: r.PostFormValue("name"), Address: r.PostFormValue("address")})
		http.Redirect(w, r, "/sql", 303)
	}
}

// SQLDelHandler SQLDelHandler
func SQLDelHandler(w http.ResponseWriter, r *http.Request) {
	id := "99999999"
	params := r.URL.Query()["id"]
	if len(params) > 0 {
		id = r.URL.Query()["id"][0]
	}
	dao.DelHospital(id)
	http.Redirect(w, r, "/sql", 303)
}
