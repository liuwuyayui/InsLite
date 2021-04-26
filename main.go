package main

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started-service")
	p := properties.MustLoadFile("credentials.properties", properties.UTF8)
	mySigningKey, _ := p.Get("mySigningKey")

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r := mux.NewRouter()
	r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
	r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", r))

}
