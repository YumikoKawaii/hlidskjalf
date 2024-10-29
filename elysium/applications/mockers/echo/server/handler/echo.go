package handler

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/mockers/echo"
	"net/http"
)

func (c *Handler) Echo(ctx context.Context, request *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.EchoResponse_Data{
			Stats: c.generateStats(utils.FakeInt(c.cfg.StatsRange.Lower, c.cfg.StatsRange.Upper)),
		},
	}, nil
}

func (c *Handler) generateStats(quantity int) []*pb.Stat {

	stats := make([]*pb.Stat, 0)
	for _ = range quantity {
		stats = append(stats, &pb.Stat{
			First:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Second:  utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Third:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Fourth:  utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Fifth:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Sixth:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Seventh: utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Eighth:  utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Ninth:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
			Tenth:   utils.FakeString(utils.FakeInt(c.cfg.StringRange.Lower, c.cfg.StringRange.Upper)),
		})
	}
	return stats
}
