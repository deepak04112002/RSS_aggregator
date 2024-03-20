package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWIthError(w http.ResponseWriter,code int,msg string){
     if code>499{
		log.Println("Responding with 5XX error: ",msg)
	 }
	 type errResponse struct{
          Error string `json:"error"`
	 }
	 responseWithJSON(w,code,errResponse{
		Error: msg,
	 })
}

func responseWithJSON(w http.ResponseWriter,code int,payload interface{}){
   data,err:=json.Marshal(payload)
   if err!=nil{
	log.Printf("Failed to Marsel json response: %v ",payload)
	w.WriteHeader(500)
	return
   }
   w.Header().Add("Cotent-Type","application/json")
   w.WriteHeader(code)
   w.Write(data)
}
