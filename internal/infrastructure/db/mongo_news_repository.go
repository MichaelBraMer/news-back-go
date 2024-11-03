package db

import (
	"context"
	"errors"
	"news-back-go/internal/app/core"
	"news-back-go/internal/app/ports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoNewsRepository struct {
	collection *mongo.Collection
}

func NewMongoNewsRepository(client *mongo.Client, dbName, colName string) ports.NewsRepository {
	return &MongoNewsRepository{collection: client.Database(dbName).Collection(colName)}
}

func (repo *MongoNewsRepository) Create(news *core.News) error {
	_, err := repo.collection.InsertOne(context.TODO(), news)
	return err
}

func (repo *MongoNewsRepository) GetAll() ([]*core.News, error) {
	var newsList []*core.News

	cursor, err := repo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var news core.News
		err := cursor.Decode(&news)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, &news)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return newsList, nil
}

func (repo *MongoNewsRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	filter := bson.M{"_id": objectID}
	_, err = repo.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (repo *MongoNewsRepository) Update(news *core.News) error {
	if news.ID == "" {
		return errors.New("ID de noticia es requerido para actualizar")
	}

	objectID, err := primitive.ObjectIDFromHex(news.ID)
	if err != nil {
		return errors.New("ID inválido")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title":     news.Title,
			"paragraph": news.Paragraph,
		},
	}

	_, err = repo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoNewsRepository) GetById(id string) (*core.News, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var news core.News
	filter := bson.M{"_id": objectID}
	err = repo.collection.FindOne(context.TODO(), filter).Decode(&news)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("noticia no encontrada")
	}
	return &news, err
}
