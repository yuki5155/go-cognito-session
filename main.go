package main

import (
	"fmt"
	//"news.com/events/sample"
	//"lambdago.com/events/session"
	"session/session"
	"encoding/json"
	"os"
)

func main() {
	fmt.Println("Hello World")


	var s session.SessionStruct

	s.Credential.AccessToken = "accesstoken"
	fmt.Println(s)
	bytes2, _ := json.Marshal(s)
    fmt.Println(string(bytes2))
	
	host := os.Getenv("REDIS_ENDPOINT")
	fmt.Println(host)

	session.Redis_Save("Key2", "Value2")
	session.Redis_Get("Key2")



}