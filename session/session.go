package session

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Session Session
type Session struct {
	Content  string
	Identify int
}

var allSessions map[string]*Session

// InitSessionStorage a
func InitSessionStorage() {
	allSessions = make(map[string]*Session)
}

// getRandomCode getRandomCode
func getRandomCode() string {
	u1 := uuid.NewV4()
	return u1.String()
}

// getSessionCookie getSessionCookie
func getSessionCookie(day int) (*http.Cookie, string) {
	// expiration := time.Now()
	// expiration = expiration.AddDate(0, 0, day)
	value := getRandomCode()
	// cookie := http.Cookie{Name: "GO_SESSION", Value: value, Expires: expiration}
	cookie := http.Cookie{Name: "GO_SESSION", Value: value}
	return &cookie, value
}

// SetSession SetSession
func SetSession(w http.ResponseWriter, params *Session) {
	session, code := getSessionCookie(1)
	for k, v := range allSessions { // 清除已经登录当前账号的会话信息
		if params.Identify == v.Identify {
			delete(allSessions, k)
		}
	}
	allSessions[code] = params
	http.SetCookie(w, session)
}

// GetSession GetSession
func GetSession(r *http.Request) *Session {
	cookie, err := r.Cookie("GO_SESSION")
	if err != nil {
		return nil
	}
	return allSessions[cookie.Value]
}

// DelSession DelSession
func DelSession(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("GO_SESSION")
	delete(allSessions, c.Value)
	cookie := http.Cookie{Name: "GO_SESSION", MaxAge: -1}
	http.SetCookie(w, &cookie)
}
