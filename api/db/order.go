package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/manavjain2002/go-amazon-clone-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection = Db.Collection("Order")

func getOneOrder(id primitive.ObjectID) models.Order {
	data := orderCollection.FindOne(context.Background(), bson.M{"_id": id})

	var order models.Order
	err := data.Decode(&order)
	if err != nil {
		fmt.Println("Unable to find Order with id", id, ".Error: ", err)
	}

	return order
}

func getAllOrders() []models.Order {
	cursor, err := orderCollection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to find all orders. Error: ", err)
	}

	var orders []models.Order
	for cursor.Next(context.Background()) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			fmt.Println("Unable to decode orders. Error: ", err)
		}
		orders = append(orders, order)
	}

	return orders
}

func insertOneOrder(order models.Order) (*mongo.InsertOneResult, error) {
	inserted, err := orderCollection.InsertOne(context.Background(), order)
	if err != nil {
		fmt.Println("Unable to insert one orders. Error: ", err)
		return &mongo.InsertOneResult{}, err
	}

	fmt.Println("Order created with id : ", inserted.InsertedID)
	fmt.Println("Created Order : ", order)
	return inserted, nil
}

func updateOneOrder(id primitive.ObjectID, updatedValues primitive.M) models.Order {

	filter := bson.M{"_id": id}

	update := bson.M{"$set": updatedValues}

	updated, err := orderCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Unable to update one order. Error: ", err)
	}

	if updated.ModifiedCount > 0 {
		fmt.Println("Order successfully updated with id: ", id)
	}

	updatedOrder := getOneOrder(id)

	return updatedOrder
}

func deleteOneOrder(id primitive.ObjectID) models.Order {

	deletedOrder := getOneOrder(id)
	filter := bson.M{"_id": id}

	deleted, err := orderCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Unable to delete one order. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Order successfully updated with id: ", id)
	}

	return deletedOrder
}

func deleteAllOrders() []models.Order {

	deletedOrders := getAllOrders()

	deleted, err := orderCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to delete all orders. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Orders successfully deleted")
	}

	return deletedOrders
}

func GetOneOrder(id string) (models.Order, error) {
	orderId, _ := primitive.ObjectIDFromHex(id)
	order := getOneOrder(orderId)

	if order == (models.Order{}) {
		return models.Order{}, errors.New("no such order")
	}

	return order, nil
}

func GetAllOrders() ([]models.Order, error) {
	orders := getAllOrders()

	if len(orders) > 0 {
		return []models.Order{}, errors.New("no orders")
	}
	return orders, nil
}

func CreateOneOrder(order models.Order) error {
	product, _ := GetOneProduct(order.ProductID.Hex())

	order.Price = product.Price * order.Quantity

	inserted, err := insertOneOrder(order)
	if err != nil {
		return err
	}
	if inserted != (&mongo.InsertOneResult{}){
	
		var updatedValue = primitive.M{
			"quantity": product.Quantity-order.Quantity,
		}

		var updatedProduct = UpdateOneProduct(order.ProductID.Hex(), updatedValue)
		if updatedProduct != (models.Product{}){
			return errors.New("product quantity not yet updated")
		}
		return nil
	} else {
		return errors.New("unable to insert new order")
	}
}

func UpdateOneOrder(id string, updatedValues primitive.M) models.Order {
	orderId, _ := primitive.ObjectIDFromHex(id)
	order := getOneOrder(orderId)

	if order != (models.Order{}) {
		updatedOrder := updateOneOrder(orderId, updatedValues)
		return updatedOrder
	}

	return models.Order{}
}

func DeleteOneOrder(id string) models.Order {
	orderId, _ := primitive.ObjectIDFromHex(id)
	order := getOneOrder(orderId)

	if order != (models.Order{}) {
		deletedOrder := deleteOneOrder(orderId)
		return deletedOrder
	}

	return models.Order{}
}

func DeleteAllOrders() []models.Order {
	deletedOrders := deleteAllOrders()
	return deletedOrders
}
