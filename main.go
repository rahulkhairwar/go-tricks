package main

import (
	"context"
	factoryPattern "factoryPattern/component"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	ctx := context.TODO()
	db := initMongoAndGetDB(ctx)
	componentRepo := factoryPattern.NewComponentRepo(db.Collection("components"))
	cs := factoryPattern.NewComponentServer(componentRepo)
	sgComp := factoryPattern.Component{
		Name: "Singapore Component",
		Code: "SG-COMP-001",
		CountryCode: "SG",
		Compliance: &factoryPattern.SingaporeCompliance{
			CountryCode: "SG",
			SomeInfo:    "Singapore compliance info...",
			ValidForCPF: true,
			SDL:         false,
			SHG:         true,
		},
	}
	inComp := factoryPattern.Component{
		Name: "India Component",
		Code: "IN-COMP-001",
		CountryCode: "IN",
		Compliance: &factoryPattern.IndiaCompliance{
			CountryCode:       "IN",
			IndiaSpecificInfo: "India specific info...",
		},
	}

	if err := cs.AddComponent(ctx, sgComp); err != nil {
		log.Println("[main] failed to add component due to : ", err)
	}
	if err := cs.AddComponent(ctx, inComp); err != nil {
		log.Println("[main] failed to add component due to : ", err)
	}

	sgc, err := cs.GetComponent(ctx, "SG-COMP-001")
	if err != nil {
		log.Println("[main] failed to get component due to : ", err)
	}
	inc, err := cs.GetComponent(ctx, "IN-COMP-001")
	if err != nil {
		log.Println("[main] failed to get component due to : ", err)
	}
	log.Println("SG Component : ", sgc)
	log.Println("IN Component : ", inc)
}

func initMongoAndGetDB(ctx context.Context) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/factory")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("error occurred while trying to connect to mongoDb : ", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("could not ping mongoDB due to error : ", err.Error())
	}
	return client.Database("factory_pattern_db")
}
