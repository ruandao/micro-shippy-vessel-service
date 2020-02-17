package main

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *StoreSpecification) (*StoreResponse, error) {
	query := bson.M{"MaxWeight": bson.M{"$gte": spec.MaxWeight}, "Capacity": bson.M{"$gte": spec.Capacity}}
	cur, err := repository.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	if cur.Next(ctx) {
		var vessel StoreVessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		return &StoreResponse{
			Vessel:  &vessel,
			Vessels: nil,
		}, nil
	}
	return nil, errors.New("No vessel found by that spec")
}

func (repository *MongoRepository) Create(ctx context.Context, vessel *StoreVessel) (*StoreResponse, error) {
	_, err := repository.collection.InsertOne(ctx, vessel)
	if err != nil {
		return nil, err
	}
	cur, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var vessels []*StoreVessel
	if cur.Next(ctx) {
		var vessel StoreVessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		vessels = append(vessels, &vessel)
	}
	return &StoreResponse{
		Vessel:  vessel,
		Vessels: vessels,
	}, nil
}
