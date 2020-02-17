package main

import "context"
import (
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	store, err := h.repository.FindAvailable(ctx, MarshalSpecification(spec))
	if err != nil {
		return err
	}
	vessels := make([]*pb.Vessel, 0, len(store.Vessels))
	for _, vessel := range store.Vessels {
		vessels = append(vessels, UnmarshalVessel(vessel))
	}
	resp.Vessel = UnmarshalVessel(store.Vessel)
	resp.Vessels = vessels
	return nil
}

func (h *handler) Create(ctx context.Context, vessel *pb.Vessel, resp *pb.Response) error {
	store, err := h.repository.Create(ctx, MarshalVessel(vessel))
	if err != nil {
		return err
	}
	vessels := make([]*pb.Vessel, 0, len(store.Vessels))
	for _, vessel := range store.Vessels {
		vessels = append(vessels, UnmarshalVessel(vessel))
	}
	resp.Vessel = UnmarshalVessel(store.Vessel)
	resp.Vessels = vessels
	return nil
}

