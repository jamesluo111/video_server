package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func steamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video, err := os.Open(vl)
	if err != nil {
		sentErrorResponse(w, 500, "bad")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取最大范围内的视频数据
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//判断上传内容是否超过最大值
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sentErrorResponse(w, http.StatusBadRequest, "file to big!")
		return
	}
	//获取文件流并验证传输是否为文件类型
	file, _, err := r.FormFile("file")
	if err != nil {
		sentErrorResponse(w, http.StatusInternalServerError, "error internal")
		return
	}
	//读取文件
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file")
		sentErrorResponse(w, http.StatusInternalServerError, "Error read file")
		return
	}
	//文件名
	fn := p.ByName("vid-id")
	//文件生成
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Println("file upload error")
		sentErrorResponse(w, http.StatusInternalServerError, "file upload error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload successfully!")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
