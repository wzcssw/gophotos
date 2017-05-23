package controllers

import (
	"io"
	"net/http"
	"os"
	"photos/config"
	"photos/models"
	"photos/session"
)

// Result asd
type phontoResult struct {
	Data    []models.Photo
	Session *session.Session
}

// IndexHandler 首页控制器
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Authenticate(w, r)
	sess := session.GetSession(r)
	Templates["./pages/photos/index.html"].Execute(w, &phontoResult{Data: models.GetFiles(), Session: sess})

	// Templates["./pages/photos/index.html"].Execute(w, models.GetFiles())
}

// ImgHandler 静态文件 控制器
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, config.Conf.FilePath+r.URL.Query()["filename"][0])
}

// DelHandler 删除文件控制器
func DelHandler(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(config.Conf.FilePath + r.URL.Query()["filename"][0])
	if err != nil {
		io.WriteString(w, "文件删除错误")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// UploadHandler 上传控制器
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("myPicture")
	if err != nil {
		return
	}
	defer file.Close()
	f, err := os.OpenFile(config.Conf.FilePath+handler.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer f.Close()
	io.Copy(f, file)
	http.Redirect(w, r, "/", 302)
}
