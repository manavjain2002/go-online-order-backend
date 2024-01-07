package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/manavjain2002/go-amazon-clone-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productCollection = Db.Collection("Product")

func getOneProduct(id primitive.ObjectID) models.Product {
	data := productCollection.FindOne(context.Background(), bson.M{"_id": id})

	var product models.Product
	err := data.Decode(&product)
	if err != nil {
		fmt.Println("Unable to find Product with id", id, ".Error: ", err)
		return models.Product{}
	}

	return product
}

func getAllProducts() []models.Product {
	cursor, err := productCollection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to find all products. Error: ", err)
	}

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			fmt.Println("Unable to decode products. Error: ", err)
		}
		products = append(products, product)
	}

	return products
}

func insertOneProduct(product models.Product) interface{} {
	inserted, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		fmt.Println("Unable to insert one products. Error: ", err)
		return primitive.NilObjectID
	}

	fmt.Println("Product created with id : ", inserted.InsertedID)
	return inserted.InsertedID
}

func updateOneProduct(id primitive.ObjectID, updatedValues primitive.M) models.Product {

	filter := bson.M{"_id": id}

	update := bson.M{"$set": updatedValues}

	updated, err := productCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Unable to update one product. Error: ", err)
	}

	if updated.ModifiedCount > 0 {
		fmt.Println("Product successfully updated with id: ", id)
	}

	updatedProduct := getOneProduct(id)

	return updatedProduct
}

func deleteOneProduct(id primitive.ObjectID) models.Product {

	deletedProduct := getOneProduct(id)
	filter := bson.M{"_id": id}

	deleted, err := productCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Unable to delete one product. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Product successfully updated with id: ", id)
	}

	return deletedProduct
}

func deleteAllProducts() []models.Product {

	deletedProducts := getAllProducts()

	deleted, err := productCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to delete all products. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Products successfully deleted")
	}

	return deletedProducts
}

func GetOneProduct(id string) (models.Product, error) {
	productId, _ := primitive.ObjectIDFromHex(id)

	product := getOneProduct(productId)

	if product == (models.Product{}) {
		return models.Product{}, errors.New("no such id")
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	products := getAllProducts()

	if len(products) == 0 {
		return []models.Product{}, errors.New("no products")
	}
	return products, nil
}

func CreateOneProduct(product models.Product) (interface{}, error) {
	id := insertOneProduct(product)
	if id == primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("no product created")
	}
	return id, nil
}

func UpdateOneProduct(id string, updatedValues primitive.M) models.Product {
	productId, _ := primitive.ObjectIDFromHex(id)

	product := getOneProduct(productId)

	if product != (models.Product{}) {
		updatedProduct := updateOneProduct(productId, updatedValues)
		return updatedProduct
	}

	return models.Product{}
}

func DeleteOneProduct(id string) models.Product {
	productId, _ := primitive.ObjectIDFromHex(id)

	product := getOneProduct(productId)

	if product != (models.Product{}) {
		deletedProduct := deleteOneProduct(productId)
		return deletedProduct
	}

	return models.Product{}
}

func DeleteAllProducts() []models.Product {
	deletedProducts := deleteAllProducts()
	return deletedProducts
}
