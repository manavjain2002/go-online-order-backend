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

var productCollection = Db.Collection("Product")

func getOneProduct(id primitive.ObjectID) (models.Product, error) {
	data := productCollection.FindOne(context.Background(), bson.M{"_id": id})

	var product models.Product
	err := data.Decode(&product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func getAllProducts() ([]models.Product, error) {
	cursor, err := productCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return []models.Product{}, err
	}

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func insertOneProduct(product models.Product) (*mongo.InsertOneResult, error) {
	inserted, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return inserted, nil
}

func updateOneProduct(id primitive.ObjectID, updatedValues primitive.M) (models.Product, error) {

	filter := bson.M{"_id": id}

	update := bson.M{"$set": updatedValues}

	updated, err := productCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.Product{}, err
	}

	if updated.ModifiedCount > 0 {
		fmt.Println("Product successfully updated with id: ", id)
		updatedProduct, err := getOneProduct(id)
		if err != nil {
			return updatedProduct, err
		}

		return updatedProduct, nil
	} else {
		return models.Product{}, errors.New("no product modified")
	}

}

func deleteOneProduct(id primitive.ObjectID) (models.Product, error) {

	deletedProduct, err := getOneProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	filter := bson.M{"_id": id}

	deleted, err := productCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return models.Product{}, err
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Product successfully updated with id: ", id)
		return deletedProduct, nil
	} else {
		return models.Product{}, errors.New("no product deleted")
	}
}

func deleteAllProducts() ([]models.Product, error) {

	deletedProducts, err := getAllProducts()
	if err != nil {
		return []models.Product{}, err
	}
	deleted, err := productCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return []models.Product{}, err
	}

	if deleted.DeletedCount > 0 {
		return deletedProducts, nil
	} else {
		return []models.Product{}, errors.New("no deleted products")
	}

}

func GetOneProduct(id string) (models.Product, error) {
	productId, _ := primitive.ObjectIDFromHex(id)

	product, err := getOneProduct(productId)

	if err != nil {
		return models.Product{}, err
	}

	if product == (models.Product{}) {
		return models.Product{}, errors.New("no such id")
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := getAllProducts()

	if err != nil {
		return []models.Product{}, err
	}

	if len(products) == 0 {
		return []models.Product{}, errors.New("no products")
	}

	return products, nil
}

func CreateOneProduct(product models.Product) error {
	id, err := insertOneProduct(product)
	if err != nil {
		return err
	}

	if id == (&mongo.InsertOneResult{}){
		return errors.New("no product created")
	}
	return nil
}

func UpdateOneProduct(id string, updatedValues primitive.M) (models.Product, error) {
	productId, _ := primitive.ObjectIDFromHex(id)

	product, err := getOneProduct(productId)

	if err != nil {
		return models.Product{}, err
	}

	if product != (models.Product{}) {
		updatedProduct, err := updateOneProduct(productId, updatedValues)

		if err != nil {
			return models.Product{}, err
		}

		return updatedProduct, nil
	}

	return models.Product{}, errors.New("no updated product")
}

func DeleteOneProduct(id string) (models.Product, error) {
	productId, _ := primitive.ObjectIDFromHex(id)

	product, err := getOneProduct(productId)

	if err != nil {
		return models.Product{}, err
	}

	if product != (models.Product{}) {
		deletedProduct, err := deleteOneProduct(productId)
		if err != nil {
			return models.Product{}, err
		}
		return deletedProduct, nil
	}

	return models.Product{}, errors.New("no product deleted")
}

func DeleteAllProducts() ([]models.Product, error) {
	deletedProducts, err := deleteAllProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return deletedProducts, nil
}
