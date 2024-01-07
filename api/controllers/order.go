package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manavjain2002/go-amazon-clone-backend/api/db"
	"github.com/manavjain2002/go-amazon-clone-backend/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GreetForOrderPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Order Api Section")
}
func GetSingleOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	order, err := db.GetOneOrder(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(order)
	}
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := db.GetAllOrders()
	if len(orders) != 0 {
		json.NewEncoder(w).Encode(orders)
	} else {
		json.NewEncoder(w).Encode(err)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	result, err := db.CreateOneOrder(order)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	var updateValues primitive.M
	json.NewDecoder(r.Body).Decode(&updateValues)
	result, err := db.UpdateOneOrder(params["id"], updateValues)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result, err := db.DeleteOneOrder(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result, err := db.DeleteAllOrders()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		if len(result) > 0 {
			json.NewEncoder(w).Encode(result)
		} else {
			json.NewEncoder(w).Encode("No orders in db")
		}
	}
}
