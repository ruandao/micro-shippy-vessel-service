package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *StoreSpecification) (*StoreVessel, error) {
	vessel := &StoreVessel{}
	query := bson.M{"MaxWeight": bson.M{"$lte": spec.MaxWeight}, "Capacity": bson.M{"$lte": spec.Capacity}}
	if err := repository.Collection.FindOne(ctx, query).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repository *MongoRepository) Create(ctx context.Context, vessel *StoreVessel) error {
	_, err := repository.Collection.InsertOne(ctx, vessel)
	if err != nil {
		return err
	}
	return nil
}
