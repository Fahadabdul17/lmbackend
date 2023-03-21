package fahad

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDataAkreditas(db string, dataakreditas DataAkreditas) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dataakreditas").InsertOne(context.TODO(), dataakreditas)
	if err != nil {
		fmt.Printf("InsertDataAkreditas: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataProgramStudi(db string, dataprogramstudi DataProgramStudi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dataprogramstudi").InsertOne(context.TODO(), dataprogramstudi)
	if err != nil {
		fmt.Printf("InsertDataProgramStudi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProfile(db string, profile Profile) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("profile").InsertOne(context.TODO(), profile)
	if err != nil {
		fmt.Printf("InsertProfile: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDataCompFromStatus(status string, db *mongo.Database, col string) (data DataAkreditas) {
	user := db.Collection(col)
	filter :=bson.M{"status": status}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataAkreditas: %v\n", err)
	}
	return data
}

func GetDataAllbyStats(stats string, db *mongo.Database, col string) (data []DataAkreditas) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}