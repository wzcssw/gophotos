package controllers

import (
	"net/http"
	"photos/dao"
	"photos/session"
)

// SessionHandler login page
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(r)
	if sess != nil { // 如果已登录跳转的首页
		http.Redirect(w, r, "/", 302)
		return
	}

	if r.Method == "GET" {
		Templates["./pages/session/login.html"].Execute(w, nil)
	} else {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		user, hasResult := dao.GetUser(username, password)
		if hasResult {
			s := &session.Session{Content: user.Realname, Identify: user.ID}
			session.SetSession(w, s)
			http.Redirect(w, r, "/", 302)
		} else {
			Templates["./pages/session/login.html"].Execute(w, "用户名或密码错误")
		}
	}
}

// SessionDelHandler logout
func SessionDelHandler(w http.ResponseWriter, r *http.Request) {
	session.DelSession(w, r)
	http.Redirect(w, r, "/session", 302)
}

// Authenticate Authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(r)
	if session == nil {
		http.Redirect(w, r, "/session", 302)
		return
	}
}
