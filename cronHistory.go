package main

import (
	"context"
	"fmt"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type (
	DbCollectionInfo struct {
		CollectionName string
		DBName string
	}

	CronjobHistory struct {
		name        string `bson:"name"`
		executeTime time.Time `bson:"executeTime"`
		dbClient *mongo.Client
		dbCtx *context.Context
		collectionInfo DbCollectionInfo
	}
)

const (
	CollectionName = "cronjob"
	DBName = "weather"
)

func (c CronjobHistory) GetLastTime(name string) time.Time {
	collection := c.getCollection()
	queryOptions := options.FindOneOptions{Sort:bson.D{{"executeTime", -1}}}

	var result CronjobHistory
	if err := collection.FindOne(*c.dbCtx, bson.D{{Key: "name", Value: c.name}}, &queryOptions).Decode(result); err != nil {
		log.Fatal("Could not search a last CronjobHistory; ", err)
	}
	return result.executeTime
}

func (c CronjobHistory) LogExecute(name string, lastExecuteTime time.Time) bool {
	collection := c.getCollection()
	insertOneResult, err := collection.InsertOne(*c.dbCtx, bson.D{
		{Key: "name", Value: c.name},
		{Key: "executeTime", Value: time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[func LogExecute] Create ", insertOneResult)
	return true
}

func (c CronjobHistory) getCollection() *mongo.Collection {
	return c.dbClient.Database(c.collectionInfo.DBName).Collection(c.collectionInfo.CollectionName)
}

func NewCronjobHistory(db *mongo.Client, ctx *context.Context) _interface.CronjobHistory {
	return &CronjobHistory{
		dbClient: db,
		dbCtx: ctx,
		collectionInfo: DbCollectionInfo{
			DBName: DBName,
			CollectionName: CollectionName,
		},
	}
}
