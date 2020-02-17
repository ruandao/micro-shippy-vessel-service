package main

import "context"
import (
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err := h.repository.FindAvailable(ctx, MarshalSpecification(spec))
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, vessel *pb.Vessel, resp *pb.Response) error {
	err := h.repository.Create(ctx, MarshalVessel(vessel))
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

