package main

import (
	"log"
	"net/http"
	"strings"
)

func HttpSend(url string, content string)  {
	_,err:=http.Post(url,"application/json", strings.NewReader(content))
	if err!=nil{
		log.Printf("http.post err %s ,url=%s \n",err.Error(),url)
	}
}
