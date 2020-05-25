package main


import (
	"context"
	"fmt"
	 "log"

	 "go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/mongo"
	  //"go.mongodb.org/mongo-driver/mongo/options"
)



type Trainer struct {
	Name string
	Age  int
	City string
}

// func (mytrainer *Trainer) getTrainer() (string,error)

func (mytrainer *Trainer) getTrainer(collection *mongo.Collection) (Trainer,error) {
	//var nameTrainer string = "Souleymane"
	//return nameTrainer

	// Find a single document
	var result Trainer
	var errmongo error

	// Update a document
	filter := bson.D{{"name", "Ash"}}

	errmongo  = collection.FindOne(context.TODO(), filter).Decode(&result)
	if errmongo != nil {
		log.Fatal(errmongo)
		return  result, errmongo
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result,nil
}