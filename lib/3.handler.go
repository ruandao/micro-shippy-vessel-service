package lib

import "context"
import (
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
)

type Handler struct {
	Repository
}

func (h *Handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err := h.Repository.FindAvailable(ctx, MarshalSpecification(spec))
	if err != nil {
		return err
	}

	resp.Vessel = UnmarshalVessel(vessel)
	return nil
}

func (h *Handler) Create(ctx context.Context, vessel *pb.Vessel, resp *pb.Response) error {
	err := h.Repository.Create(ctx, MarshalVessel(vessel))
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

