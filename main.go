package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"


	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main(){

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("port is not found in environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https:/*","https:/*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))


	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerErr)


	router.Mount("/v1",v1Router)
	srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}

	log.Printf("Server starting on port %v",portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}