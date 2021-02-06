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
		//ID    primitive.ObjectID `bson:"_id"`
		Name string
		ExecuteTime time.Time

		dbClient *mongo.Client
		dbCtx *context.Context
		collectionInfo DbCollectionInfo
	}
)

const (
	CollectionName = "cronjob"
	DBName = "weather"
)

func (c CronjobHistory) GetLastTime(name string) (lastTime time.Time, isFirstTime bool) {
	col := c.getCollection()

	var result CronjobHistory
	opt := options.FindOneOptions{Sort:bson.D{{"ExecuteTime", -1}}}
	filter := bson.D{{"Name", "기상청"}}

	err := col.FindOne(*c.dbCtx, filter, &opt).Decode(&result);
	// Empty result, Run first time
	if err != nil && result.Name == "" {
		return time.Now(), true
	}
	return result.ExecuteTime, false
}

func (c CronjobHistory) LogExecute(name string, lastExecuteTime time.Time) bool {
	collection := c.getCollection()
	insertOneResult, err := collection.InsertOne(*c.dbCtx, bson.D{
		{Key: "Name", Value: name},
		{Key: "ExecuteTime", Value: lastExecuteTime},
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
