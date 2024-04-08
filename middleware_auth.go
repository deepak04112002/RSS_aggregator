package main

import (
	"fmt"
	"net/http"

	"github.com/deepak04112002/rssagg/internal/auth"
	"github.com/deepak04112002/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey,err:=auth.GetAPIKey(r.Header)
		if err!=nil{
			responseWithError(w, 403, fmt.Sprintf("Auth Error: %s", err))
			return
		}
		user,err:=apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err!=nil{
			responseWithError(w, 400, fmt.Sprintf("Couldnt get user: %s", err))
			return
		}
		handler(w,r,user)
	}
}