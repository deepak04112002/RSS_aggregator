package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString==""{
		log.Fatal("PORT is not found in environment")
	}
    
	router :=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
       AllowedOrigins: []string{"https://*","http://*"},
	   AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
	   AllowedHeaders: []string{"*"},
	   ExposedHeaders: []string{"Link"},
	   AllowCredentials: false,
	   MaxAge: 300,
	}))
    
	v1Router:=chi.NewRouter()
	v1Router.Get("/healthz",handleReadiness)
	v1Router.Get("/err",handleErr)

	router.Mount("/v1",v1Router)

	srv:=&http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	log.Printf("Server started at %v",portString)
	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}