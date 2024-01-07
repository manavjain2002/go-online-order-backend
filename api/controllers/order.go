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
	orders, _ := db.GetAllOrders()
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	result := db.CreateOneOrder(order)
	json.NewEncoder(w).Encode(result)
}


func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	
	var updateValues primitive.M
	json.NewDecoder(r.Body).Decode(&updateValues)
	result := db.UpdateOneOrder(params["id"], updateValues)
	json.NewEncoder(w).Encode(result)
}


func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := db.DeleteOneOrder(params["id"])	
	json.NewEncoder(w).Encode(result)
}

func DeleteAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result := db.DeleteAllOrders()	
	if(len(result) > 0){
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode("No orders in db")
	}
}

