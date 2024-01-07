package routers

import (
	"github.com/gorilla/mux"
	"github.com/manavjain2002/go-amazon-clone-backend/api/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.GreeterForHome).Methods("GET")
	router.HandleFunc("/api", controllers.GreetForApi).Methods("GET")

	router.HandleFunc("/api/user", controllers.GreetForUserPage).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetSingleUser).Methods("GET")
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", controllers.DeleteAllUsers).Methods("DELETE")

	router.HandleFunc("/api/product", controllers.GreetForProductPage).Methods("GET")
	router.HandleFunc("/api/product/{id}", controllers.GetSingleProduct).Methods("GET")
	router.HandleFunc("/api/products", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/products", controllers.DeleteAllProducts).Methods("DELETE")

	router.HandleFunc("/api/order", controllers.GreetForOrderPage).Methods("GET")
	router.HandleFunc("/api/order/{id}", controllers.GetSingleOrder).Methods("GET")
	router.HandleFunc("/api/orders", controllers.GetAllOrders).Methods("GET")
	router.HandleFunc("/api/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/order/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/order/{id}", controllers.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/api/orders", controllers.DeleteAllOrders).Methods("DELETE")


	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/home", controllers.Home)
	router.HandleFunc("/refresh", controllers.Refresh)

	
	return router
}
