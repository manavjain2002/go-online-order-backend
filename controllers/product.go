package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manavjain2002/go-amazon-clone-backend/db"
	"github.com/manavjain2002/go-amazon-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GreetForProductPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Product Api Section")
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	user, err := db.GetOneProduct(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, _ := db.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	result, err := db.CreateOneProduct(product)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	
	var updateValues primitive.M
	json.NewDecoder(r.Body).Decode(&updateValues)
	result := db.UpdateOneProduct(params["id"], updateValues)
	json.NewEncoder(w).Encode(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := db.DeleteOneProduct(params["id"])	
	json.NewEncoder(w).Encode(result)
}

func DeleteAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result := db.DeleteAllProducts()	
	json.NewEncoder(w).Encode(result)
}