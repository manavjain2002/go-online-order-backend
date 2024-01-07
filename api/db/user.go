package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/manavjain2002/go-amazon-clone-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = Db.Collection("User")

func getOneUser(id string) models.User {
	userId, _ := primitive.ObjectIDFromHex(id)
	data := userCollection.FindOne(context.Background(), bson.M{"_id": userId})

	var user models.User
	err := data.Decode(&user)
	if err != nil {
		fmt.Println("Unable to find User with id", id, ".Error: ", err)
	}

	return user
}

func getAllUsers() []models.User {
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to find all users. Error: ", err)
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
	fmt.Println(users)
	return users
}

func findUserByEmail(email string) primitive.ObjectID {
	data := userCollection.FindOne(context.Background(), bson.M{"email": email})
	var user models.User
	if err := data.Decode(&user); err != nil && user != (models.User{}) {
		return primitive.NilObjectID
	}
	return user.UserID
}

func insertOneUser(user models.User) error {
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	fmt.Println("User created with id : ", inserted.InsertedID)
	fmt.Println("Created User : ", user)
	return nil
}

func updateOneUser(id string, updatedValues primitive.M) models.User {

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": userId}

	update := bson.M{"$set": updatedValues}

	updated, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Unable to update one user. Error: ", err)
	}

	if updated.ModifiedCount > 0 {
		fmt.Println("User successfully updated with id: ", id)
	}

	updatedUser := getOneUser(id)

	return updatedUser
}

func deleteOneUser(id string) models.User {
	userId, _ := primitive.ObjectIDFromHex(id)

	deletedUser := getOneUser(id)
	filter := bson.M{"_id": userId}

	deleted, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Unable to delete one user. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("User successfully updated with id: ", id)
	}

	return deletedUser
}

func deleteAllUsers() []models.User {

	deletedUsers := getAllUsers()

	deleted, err := userCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Unable to delete all users. Error: ", err)
	}

	if deleted.DeletedCount > 0 {
		fmt.Println("Users successfully deleted")
	}

	return deletedUsers
}

func GetOneUser(id string) (models.User, error) {
	user := getOneUser(id)

	if user == (models.User{}) {
		return models.User{}, errors.New("no such id")
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users := getAllUsers()

	if len(users) <= 0 {
		return []models.User{}, errors.New("no users")
	}
	return users, nil
}

func CreateOneUser(user models.User) primitive.ObjectID {
	id := findUserByEmail(user.Email)
	if id.IsZero() {
		err := insertOneUser(user)
		if err != nil {
			return primitive.NilObjectID
		}
	} else {
		return primitive.NilObjectID
	}
	userId := findUserByEmail(user.Email)
	return userId
}

func UpdateOneUser(id string, updatedValues primitive.M) models.User {
	user := getOneUser(id)

	if user != (models.User{}) {
		updatedUser := updateOneUser(id, updatedValues)
		return updatedUser
	}

	return models.User{}
}

func DeleteOneUser(id string) models.User {
	user := getOneUser(id)

	if user != (models.User{}) {
		deletedUser := deleteOneUser(id)
		return deletedUser
	}

	return models.User{}
}

func DeleteAllUsers() []models.User {
	deletedUsers := deleteAllUsers()
	if len(deletedUsers) == 0 {
		return []models.User{}
	}
	return deletedUsers
}
