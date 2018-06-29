package session

import "github.com/astaxie/beego/session"

var GlobalSessions *session.Manager

func init()  {
	sessionConfig := &session.ManagerConfig{
		CookieName: "IOSTSSID",
		EnableSetCookie: true,
		Gclifetime: 3600,
		Maxlifetime: 3600,
		Secure: true,
		CookieLifeTime: 0,
		ProviderConfig: "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}