package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type User struct{
	UserID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"username" bson:"username"`
	Password string `json:"password"`
	Number string `json:"number"`
	Email string `json:"email"`
	NumberOfOrders int `json:"noOfOrders"`
}

type Product struct{
	ProductID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	ColorOfClothes string `json:"color"`
	CategoryOfClothes string `json:"category"`
	Size string `json:"size"`
	Brand string `json:"brand"`
}

type Order struct{
	OrderID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	ProductID primitive.ObjectID `json:"productId" bson:"productId"`
	OrderDate time.Time `json:"orderDate"`
	DeliveryDate time.Time `json:"deliveryDate"`
	Completed bool `json:"completed"`
}