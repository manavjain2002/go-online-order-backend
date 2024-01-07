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

func getOneOrder(id primitive.ObjectID) (models.Order, error) {
	data := orderCollection.FindOne(context.Background(), bson.M{"_id": id})

	var order models.Order
	err := data.Decode(&order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func getAllOrders() ([]models.Order, error) {
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

	if len(orders) == 0 {
		return []models.Order{}, errors.New("no order data is present in ddatabase")
	}

	return orders, nil
}

func insertOneOrder(order models.Order) (*mongo.InsertOneResult, error) {
	inserted, err := orderCollection.InsertOne(context.Background(), order)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return inserted, nil
}

func updateOneOrder(id primitive.ObjectID, updatedValues primitive.M) (models.Order, error) {

	filter := bson.M{"_id": id}

	update := bson.M{"$set": updatedValues}

	updated, err := orderCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.Order{}, err
	}

	if updated.ModifiedCount > 0 {
		updatedOrder, _ := getOneOrder(id)
		return updatedOrder, nil
	}
	return models.Order{}, errors.New("no matching order found")
}

func deleteOneOrder(id primitive.ObjectID) (models.Order, error) {

	deletedOrder, _ := getOneOrder(id)
	if deletedOrder == (models.Order{}) {
		return models.Order{}, errors.New("no matching order found")
	}

	filter := bson.M{"_id": id}

	deleted, err := orderCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return models.Order{}, errors.New("Unable to delete one order. Error : " + err.Error())
	}

	if deleted.DeletedCount == 0 {
		return models.Order{}, errors.New("no deleted order")
	}

	return deletedOrder, nil
}

func deleteAllOrders() ([]models.Order, error) {

	deletedOrders, _ := getAllOrders()

	deleted, err := orderCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return []models.Order{}, err
	}

	if deleted.DeletedCount == 0 {
		return []models.Order{}, errors.New("no deleted record")
	}

	return deletedOrders, nil
}

func GetOneOrder(id string) (models.Order, error) {
	orderId, _ := primitive.ObjectIDFromHex(id)

	order, err := getOneOrder(orderId)

	if err != nil {
		return models.Order{}, err
	}

	if order == (models.Order{}) {
		return models.Order{}, errors.New("no such order")
	}

	return order, nil
}

func GetAllOrders() ([]models.Order, error) {
	orders, err := getAllOrders()

	if err != nil {
		return []models.Order{}, err
	}

	if len(orders) > 0 {
		return []models.Order{}, errors.New("no orders")
	}
	return orders, nil
}

func CreateOneOrder(order models.Order) (models.Order, error) {
	product, _ := GetOneProduct(order.ProductID.Hex())

	order.Price = product.Price * order.Quantity

	inserted, err := insertOneOrder(order)
	if err != nil {
		return models.Order{}, err
	}
	if inserted != (&mongo.InsertOneResult{}) {

		id := inserted.InsertedID.(string)

		createdOrder, _ := GetOneOrder(id)
		var updatedValue = primitive.M{
			"quantity": product.Quantity - order.Quantity,
		}

		updatedProduct, err := UpdateOneProduct(order.ProductID.Hex(), updatedValue)
		if err != nil {
			return models.Order{}, errors.New("product quantity not updated")
		}
		if updatedProduct != (models.Product{}) {
			return createdOrder, errors.New("product quantity not yet updated")
		}
		return createdOrder, nil
	} else {
		return models.Order{}, errors.New("unable to insert new order")
	}
}

func UpdateOneOrder(id string, updatedValues primitive.M) (models.Order, error) {
	orderId, _ := primitive.ObjectIDFromHex(id)
	order, err := getOneOrder(orderId)

	if err != nil {
		return models.Order{}, err
	}

	if order != (models.Order{}) {
		updatedOrder, err := updateOneOrder(orderId, updatedValues)
		if err != nil {
			return models.Order{}, err
		}
		return updatedOrder, nil
	}

	return models.Order{}, errors.New("no updated order")
}

func DeleteOneOrder(id string) (models.Order, error) {
	orderId, _ := primitive.ObjectIDFromHex(id)
	order, err := getOneOrder(orderId)

	if err != nil {
		return models.Order{}, err
	}

	if order != (models.Order{}) {
		deletedOrder, err := deleteOneOrder(orderId)

		if err != nil {
			return models.Order{}, err
		}

		return deletedOrder, nil
	}

	return models.Order{}, errors.New("no deleted order")
}

func DeleteAllOrders() ([]models.Order, error) {
	deletedOrders, err := deleteAllOrders()
	if err != nil {
		return []models.Order{}, err
	}
	return deletedOrders, nil
}
