package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)
func setupRoutes() error{
	log.Println("setup router")

	r := mux.NewRouter()
	r.HandleFunc("/register",registerUser)
	err:=http.ListenAndServe(":8080",r)
	if err!=nil{
		return err
	}
	return nil
}