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

var userCollection = Db.Collection("User")

func getOneUser(id string) (models.User, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	data := userCollection.FindOne(context.Background(), bson.M{"_id": userId})

	var user models.User
	err := data.Decode(&user)
	if err != nil {
		return models.User{}, errors.New("Unable to find User with id" + id + ".Error: " + err.Error())
	}

	return user, nil
}

func getAllUsers() ([]models.User, error) {
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Unable to decode users. Error: ", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func findUserByEmail(email string) primitive.ObjectID {
	data := userCollection.FindOne(context.Background(), bson.M{"email": email})
	var user models.User
	if err := data.Decode(&user); err != nil && user != (models.User{}) {
		return primitive.NilObjectID
	}
	return user.UserID
}

func insertOneUser(user models.User) (*mongo.InsertOneResult, error) {
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return inserted, nil
}

func updateOneUser(id string, updatedValues primitive.M) (models.User, error) {

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": userId}

	update := bson.M{"$set": updatedValues}

	updated, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.User{}, err
	}

	if updated.ModifiedCount > 0 {
		updatedUser, err := getOneUser(id)
		if err != nil {
			return models.User{}, err
		}

		return updatedUser, nil
	} else {
		return models.User{}, errors.New("no updated user")
	}

}

func deleteOneUser(id string) (models.User, error) {
	userId, _ := primitive.ObjectIDFromHex(id)

	deletedUser, err := getOneUser(id)
	if err != nil {
		return models.User{}, err
	}
	filter := bson.M{"_id": userId}

	deleted, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return models.User{}, err
	}

	if deleted.DeletedCount > 0 {
		return deletedUser, nil
	} else {
		return models.User{}, errors.New("no deleted user")
	}

}

func deleteAllUsers() ([]models.User, error) {

	deletedUsers, err := getAllUsers()
	if err != nil {
		return []models.User{}, err
	}

	deleted, err := userCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return []models.User{}, err
	}

	if deleted.DeletedCount > 0 {
		return deletedUsers, nil
	} else {
		return []models.User{}, errors.New("no deleted users")
	}

}

func GetOneUser(id string) (models.User, error) {
	user, err := getOneUser(id)

	if err != nil {
		return models.User{}, err
	}

	if user == (models.User{}) {
		return models.User{}, errors.New("no such id")
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := getAllUsers()

	if err != nil {
		return []models.User{}, err
	}

	if len(users) <= 0 {
		return []models.User{}, errors.New("no users")
	}
	return users, nil
}

func CreateOneUser(user models.User) primitive.ObjectID {
	id := findUserByEmail(user.Email)
	if id.IsZero() {
		_, err := insertOneUser(user)
		if err != nil {
			return primitive.NilObjectID
		}

	} else {
		return primitive.NilObjectID
	}
	userId := findUserByEmail(user.Email)
	return userId
}

func UpdateOneUser(id string, updatedValues primitive.M) (models.User, error) {
	user, err := getOneUser(id)
	if err != nil {
		return models.User{}, err
	}
	if user != (models.User{}) {
		updatedUser, err := updateOneUser(id, updatedValues)
		if err != nil {
			return models.User{}, err
		}
		return updatedUser, nil
	} else {
		return models.User{}, errors.New("no users updated")
	}

}

func DeleteOneUser(id string) (models.User, error) {
	user, err := getOneUser(id)
	if err != nil {
		return models.User{}, err
	}
	if user != (models.User{}) {
		deletedUser, err := deleteOneUser(id)
		if err != nil {
			return models.User{}, err
		}
		return deletedUser, nil
	}

	return models.User{}, errors.New("No deleted user")
}

func DeleteAllUsers() ([]models.User, error) {
	deletedUsers, err := deleteAllUsers()
	if err != nil {
		return []models.User{}, err
	}
	if len(deletedUsers) == 0 {
		return []models.User{}, errors.New("no deleted user")
	}
	return deletedUsers, nil
}
