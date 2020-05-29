package factoryPattern

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

type componentRepo struct {
	sync.RWMutex
	collection *mongo.Collection
}

func NewComponentRepo(c *mongo.Collection) ComponentRepository {
	return &componentRepo{collection: c}
}

func (c *componentRepo) AddComponent(ctx context.Context, component Component) (*Component, error) {
	c.Lock()
	defer c.Unlock()
	result, err := c.collection.InsertOne(ctx, component)
	if err != nil {
		return nil, err
	}
	log.Println("inserted.ID : ", result.InsertedID)
	return &component, nil
}

func (c *componentRepo) GetComponent(ctx context.Context, code string) (*Component, error) {
	c.RLock()
	defer c.RUnlock()

	filter := bson.D{
		primitive.E{Key: "code", Value: code},
	}

	var comp Component
	if err := c.collection.FindOne(ctx, filter).Decode(&comp); err != nil {
		return nil, err
	}
	return &comp, nil
}

func convertHexToObjectId(code string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}
