package main

import (
	"fmt"
	"net/http"

	"github.com/manavjain2002/go-amazon-clone-backend/routers"
)

func main() {
	
	fmt.Println("Welcome to Online Order App.")
	r := routers.Router()

	fmt.Println("Server starting listening.......")
	fmt.Println(http.ListenAndServe(":4000", r))
	fmt.Println("Server listening on port 4000......")

}
