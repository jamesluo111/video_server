package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

type HomePage struct {
	Name string
}

type UserHomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName, err1 := r.Cookie("username")
	session, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "jamesluo"}
		t, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Printf("Parsing template home.html error:")
			return
		}
		t.Execute(w, p)
		return
	}

	if len(userName.Value) != 0 && len(session.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//查询cookie中是否有此字段
	userName, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	uName := r.FormValue("username")
	var p *UserHomePage
	if len(userName.Value) != 0 {
		p = &UserHomePage{Name: userName.Value}
	} else if len(uName) != 0 {
		p = &UserHomePage{Name: uName}
	}
	t, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		fmt.Println("出错啦")
		return
	}
	t.Execute(w, p)
	return
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("rbody:%s", string(res))
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		w.WriteHeader(400)
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	request(apiBody, w, r)
	defer r.Body.Close()
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url, _ := url2.Parse("http://127.0.0.1:8002/")
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}
