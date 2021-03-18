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

// InsertPost ...
func InsertPost(u model.Post) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}

// CheckPostAlreadyExist ...
func CheckPostAlreadyExist(slug string) (model.Post, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	condition := bson.M{"slug": slug}

	var result model.Post

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID

	}

	return result, true, ID

}

// GetAllPosts ...
func GetAllPosts() ([]*model.Post, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	var result []*model.Post

	findOptions := options.Find()

	findOptions.SetLimit(20)

	cur, err := col.Find(ctx, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	for cur.Next(ctx) {
		var s model.Post
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

// GetPost ...
func GetPost(ID string) (model.Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	var post model.Post

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&post)

	if err != nil {
		fmt.Println("post not found" + err.Error())
		return post, err
	}

	return post, nil

}

// UpdatePost ...
func UpdatePost(u model.Post, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	post := make(map[string]interface{})

	if len(u.Body) > 0 {
		post["body"] = u.Body
	}

	if len(u.Slug) > 0 {
		post["slug"] = u.Slug
	}

	updtString := bson.M{
		"$set": post,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}
	return true, nil

}

// DeletePost ...
func DeletePost(ID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConn.Database("apirest")
	col := db.Collection("posts")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err

}
