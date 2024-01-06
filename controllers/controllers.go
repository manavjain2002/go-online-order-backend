package controllers

import (
	"encoding/json"
	"net/http"
)

func GreeterForHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Online Order Backend")
}

func GreetForApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Api Section")
}