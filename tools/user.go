package tools

import (
	model "apirest/internal/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertUser ...
func InsertUser(u model.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}

// CheckUserAlreadyExist ...
func CheckUserAlreadyExist(email string) (model.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result model.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID

	}

	return result, true, ID

}

// GetAllUsers ...
func GetAllUsers() ([]*model.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	var result []*model.User

	findOptions := options.Find()

	findOptions.SetLimit(20)

	cur, err := col.Find(ctx, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	for cur.Next(ctx) {
		var s model.User
		err := cur.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}

		result = append(result, &s)

	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	cur.Close(ctx)

	return result, true
}

// GetUser ...
func GetUser(ID string) (model.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	var user model.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&user)

	if err != nil {
		fmt.Println("user not found" + err.Error())
		return user, err
	}

	return user, nil

}

// UpdateUser ...
func UpdateUser(u model.User, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	user := make(map[string]interface{})

	if len(u.Name) > 0 {
		user["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		user["last_name"] = u.LastName
	}

	if len(u.Email) > 0 {
		user["email"] = u.Email
	}

	updtString := bson.M{
		"$set": user,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}
	return true, nil

}

// DeleteUser ...
func DeleteUser(ID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("users")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err

}
