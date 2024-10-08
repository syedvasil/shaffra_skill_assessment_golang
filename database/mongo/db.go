package dbmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB = "shaffra"

//func InitDB(uri string) (*mongo.Database, *error) {
//	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
//	if err != nil {
//		return nil, &err
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	err = client.Connect(ctx)
//	if err != nil {
//		return nil, &err
//	}
//
//	return client.Database(DB), nil
//}

func InitDBConnect(uri string) (*mongo.Database, *error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, &err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, &err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database(DB), nil
}
