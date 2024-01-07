package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manavjain2002/go-amazon-clone-backend/api/db"
	"github.com/manavjain2002/go-amazon-clone-backend/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GreetForUserPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to User Api Section")
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	user, err := db.GetOneUser(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if user == (models.User{}) {
		json.NewEncoder(w).Encode("No user found")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Print(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if len(users) == 0 {
		json.NewEncoder(w).Encode("No users found")
		return
	}
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result := db.CreateOneUser(user)
	fmt.Println(result)
	if (result == primitive.ObjectID{}) {
		json.NewEncoder(w).Encode("User already exists")
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	var updateValues primitive.M
	json.NewDecoder(r.Body).Decode(&updateValues)
	result, err := db.UpdateOneUser(params["id"], updateValues)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if result == (models.User{}) {
		json.NewEncoder(w).Encode("No user found")
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result, err := db.DeleteOneUser(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if result == (models.User{}) {
		json.NewEncoder(w).Encode("No such user found")
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result, err := db.DeleteAllUsers()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if len(result) == 0 {
		json.NewEncoder(w).Encode("No users found")
		return
	}
	json.NewEncoder(w).Encode(result)
}
