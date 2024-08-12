package repository

import (
	"backend/mod/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	MongoCollection *mongo.Collection
}

func (r *CategoryRepo) InsertCategory(category *model.Category) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), category)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *CategoryRepo) FindCategoryByID(categoryId string) (*model.Category, error) {
	var category model.Category
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "category_id", Value: categoryId}}).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) FindByEmailPass(email string) (*model.Category, error) {
	var category model.Category
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) FindAllCategorys() ([]*model.Category, error) {
	var categorys []*model.Category
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	// defer cursor.Close(context.Background())

	// for cursor.Next(context.Background()) {
	// 	var category model.Category
	// 	err := cursor.Decode(&category)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	categorys = append(categorys, &category)
	// }
	err = cursor.All(context.Background(), &categorys)
	if err != nil {
		return nil, err
	}

	return categorys, nil
}

func (r *CategoryRepo) UpdateCategory(categoryId string, category *model.Category) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "category_id", Value: categoryId}}, bson.D{{Key: "$set", Value: category}})
	if err != nil {
		return 0, err
	}
	if result.MatchedCount == 0 {
		return 0, fmt.Errorf("category not found")
	}
	return result.MatchedCount, nil
}

func (r *CategoryRepo) DeleteCategory(categoryId string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "category_id", Value: categoryId}})
	if err != nil {
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("category not found")
	}
	return result.DeletedCount, nil
}

func (r *CategoryRepo) DeleteAll() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
