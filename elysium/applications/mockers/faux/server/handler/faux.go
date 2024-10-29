package handler

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/mockers/faux"
	"math"
	"net/http"
	"time"
)

func (c *Handler) Faux(ctx context.Context, request *pb.FauxRequest) (*pb.FauxResponse, error) {

	time.Sleep(time.Duration(c.cfg.TrafficDelayInMilliSec) * time.Millisecond)

	return &pb.FauxResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.FauxResponse_Data{
			Stats: c.generateStats(utils.FakeInt(c.cfg.StatsRange.Lower, c.cfg.StatsRange.Upper)),
		},
	}, nil
}

func (c *Handler) generateStats(quantity int) []*pb.Stat {

	stats := make([]*pb.Stat, 0)
	for _ = range quantity {
		stats = append(stats, &pb.Stat{
			First:     utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Second:    utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Third:     utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Fourth:    utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Fifth:     utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Primaries: c.generatePrimaries(utils.FakeInt(c.cfg.PrimaryRange.Lower, c.cfg.PrimaryRange.Upper)),
		})
	}
	return stats
}

func (c *Handler) generatePrimaries(quantity int) []*pb.Primary {
	primaries := make([]*pb.Primary, 0)
	for _ = range quantity {
		primaries = append(primaries, &pb.Primary{
			First:   uint32(utils.FakeInt(0, math.MaxUint32)),
			Second:  uint32(utils.FakeInt(0, math.MaxUint32)),
			Third:   uint32(utils.FakeInt(0, math.MaxUint32)),
			Fourth:  uint32(utils.FakeInt(0, math.MaxUint32)),
			Fifth:   uint32(utils.FakeInt(0, math.MaxUint32)),
			Sixth:   uint32(utils.FakeInt(0, math.MaxUint32)),
			Seventh: uint32(utils.FakeInt(0, math.MaxUint32)),
			Eighth:  uint32(utils.FakeInt(0, math.MaxUint32)),
			Ninth:   uint32(utils.FakeInt(0, math.MaxUint32)),
			Tenth:   uint32(utils.FakeInt(0, math.MaxUint32)),
		})
	}
	return primaries
}
