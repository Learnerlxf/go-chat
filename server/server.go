package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	libHttp "go-chat/lib/http"
)

var clientMap map[int]string
/////////////////////////http/////////////////////////
func Regist(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		UserId int `json:"user_id"`
		Ip string `json:"ip"`
	}
	reqByte,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Printf("ioutil.ReadAll err %s \n",err.Error())
		http.Error(w,err.Error(),404)
		return
	}
	reqData := &Req{}
	err = json.Unmarshal(reqByte,reqData)
	if err!=nil{
		log.Printf("json unmarshal err %s \n",err.Error())
		http.Error(w,err.Error(),404)
		return
	}

	clientMap[reqData.UserId] = reqData.Ip
	w.Write([]byte("success"))
}

func UnRegist(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		UserId int `json:"user_id"`
	}
	reqByte,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Printf("ioutil.ReadAll err %s \n",err.Error())
		http.Error(w,err.Error(),404)
		return
	}
	reqData := &Req{}
	err = json.Unmarshal(reqByte,reqData)
	if err!=nil{
		log.Printf("json unmarshal err %s \n",err.Error())
		http.Error(w,err.Error(),404)
		return
	}
	delete(clientMap, reqData.UserId)
	w.Write([]byte("success"))
}

func Say(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		UserId int `json:"user_id"`
	}
	reqByte,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Printf("ioutil.ReadAll err %s \n",err.Error())
		http.Error(w,err.Error(),404)
		return
	}
	reqStr := string(reqByte)

	type SayContent struct {
		ToUserId int `json:"to_user_id"`
		Content string `json:"content"`
	}
	for k,v :=range clientMap  {
		sayContent:= SayContent{ToUserId:k,Content:reqStr}
		sayJson,_ := json.Marshal(sayContent)
		libHttp.HttpSend(v+"/say", string(sayJson))
	}
	w.Write([]byte("success"))
}

func HttpServer() {
	clientMap = make(map[int]string,1000)
	http.HandleFunc("/regist", Regist)
	http.HandleFunc("/unregist", UnRegist)
	http.HandleFunc("/Say", Say)
	http.ListenAndServe(":8080", nil)
	fmt.Println("service time end")
}
