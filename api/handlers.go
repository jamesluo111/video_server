package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取request中所有的数据
	res, _ := ioutil.ReadAll(r.Body)
	//将request数据json
	ubody := &defs.UserCredential{}
	err := json.Unmarshal(res, ubody)
	//检查转化是否正确
	if err != nil {
		sendErrorResponse(w, defs.ErrorResponseBody)
		return
	}
	//创建用户
	if err = dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	//将用户信息存到缓存中
	sessionId, err := session.GenerateNewSession(ubody.UserName)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	//返回请求体
	su := &defs.SignedUp{Success: true, SessionId: sessionId}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternal)
		return
	} else {
		sendNormalResponse(w, 201, string(resp))
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//从form表单中获取用户名和密码
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	uname := p.ByName("user_name")
	log.Printf("url username:%s", uname)
	log.Printf("body username: %s", ubody.UserName)
	//对比url及form表单内的数据
	if uname != ubody.UserName {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//验证完成，从数据库中获取密码
	pwd, err := dbops.GetUserCredential(uname)
	upwd := ubody.Pwd
	log.Printf("url pwd:%s", pwd)
	log.Printf("body pwd: %s", upwd)
	//验证密码一致性
	if err != nil && len(pwd) == 0 && pwd != upwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	sessionId, err := session.GenerateNewSession(ubody.UserName)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	si := &defs.SignedUp{Success: true, SessionId: sessionId}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	} else {
		sendNormalResponse(w, 200, string(resp))
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	//通过username获取userid
	username := p.ByName("user_name")
	id, err := dbops.GetUser(username)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	ui := &defs.UserInfo{Id: id}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(w, defs.ErrorInternal)
		return
	} else {
		sendNormalResponse(w, 200, string(resp))
	}
}

func AddNewVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	nvBody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvBody); err != nil {
		sendErrorResponse(w, defs.ErrorResponseBody)
		return
	}

	vi, err := dbops.AddNewVideo(nvBody.UserId, nvBody.UserName)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(w, defs.ErrorInternal)
		return
	} else {
		sendNormalResponse(w, 201, string(resp))
	}

}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	log.Println("进来啦!==============")
	userName := p.ByName("user_name")
	currentTime, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixMilli(), 10))
	//查询用户全部视频
	vis, err := dbops.GetVideoList(userName, 0, currentTime)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	v := &defs.VideosInfo{Videos: vis}
	if resp, err := json.Marshal(v); err != nil {
		sendErrorResponse(w, defs.ErrorInternal)
		return
	} else {
		sendNormalResponse(w, 200, string(resp))
	}
}

func DeleteVideos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	videoId := ps.ByName("vid-id")
	err := dbops.DeleteVideoInfo(videoId)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	sendNormalResponse(w, 200, "success")

}

func PostComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	req, _ := ioutil.ReadAll(r.Body)
	addComment := &defs.AddComments{}
	if err := json.Unmarshal(req, addComment); err != nil {
		sendErrorResponse(w, defs.ErrorResponseBody)
		return
	}
	//数据库新增数据
	err := dbops.AddNewComment(addComment.VideoId, addComment.AuthorId, addComment.Content)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	sendNormalResponse(w, 201, "success")

}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//验证请求头的session里带的username
	if !validateUser(w, r) {
		log.Printf("Unthorized user\n")
		return
	}
	req := p.ByName("vid-id")
	currentTime, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixMilli(), 10))
	comments, err := dbops.GetCommentList(req, 0, currentTime)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDbOps)
		return
	}
	commentsInfo := &defs.CommentsInfo{Comments: comments}
	if resp, err := json.Marshal(commentsInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternal)
		return
	} else {
		sendNormalResponse(w, 200, string(resp))
	}
}
