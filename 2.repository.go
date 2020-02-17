package main

import (
	"context"
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
)

//rpc FindAvailable(StoreSpecification) returns (StoreResponse) {}
//rpc Create(StoreVessel) returns (StoreResponse) {}
type StoreSpecification struct {
	Capacity int32 `json: "Capacity"`
	MaxWeight int32 `json: "MaxWeight"`
}

type StoreResponse struct {
	Vessel *StoreVessel
	Vessels []*StoreVessel
}
type StoreVessel struct {
	Id                   string   `json: "Id"`
	Capacity             int32    `json: "Capacity"`
	MaxWeight            int32    `json: "MaxWeight"`
	Name                 string   `json: "Name"`
	Available            bool     `json: "Available"`
	OwnerId              string   `json: "OwnerId"`
}

func MarshalSpecification(specification *pb.Specification) *StoreSpecification {
	return &StoreSpecification{
		Capacity:  specification.Capacity,
		MaxWeight: specification.MaxWeight,
	}
}
func MarshalResponse(response *pb.Response) *StoreResponse {
	vessels := make([]*StoreVessel, 0, len(response.Vessels))
	for _, vessel := range response.Vessels {
		vessels = append(vessels, MarshalVessel(vessel))
	}
	return &StoreResponse{
		Vessel:  MarshalVessel(response.Vessel),
		Vessels: vessels,
	}
}
func MarshalVessel(vessel *pb.Vessel) *StoreVessel {
	return &StoreVessel{
		Id:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerId,
	}
}

func UnmarshalVessel(vessel *StoreVessel) *pb.Vessel {
	return &pb.Vessel{
		Id:                   vessel.Id,
		Capacity:             vessel.Capacity,
		MaxWeight:            vessel.MaxWeight,
		Name:                 vessel.Name,
		Available:            vessel.Available,
		OwnerId:              vessel.OwnerId,
	}
}

type repository interface {
	FindAvailable(ctx context.Context, spec *StoreSpecification) (*StoreVessel, error)
	Create(ctx context.Context, vessel *StoreVessel) error
}