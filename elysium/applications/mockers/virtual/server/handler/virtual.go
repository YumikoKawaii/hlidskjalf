package handler

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/mockers/virtual"
	"math"
	"net/http"
	"time"
)

func (c *Handler) Virtual(ctx context.Context, request *pb.VirtualRequest) (*pb.VirtualResponse, error) {

	if c.cfg.TrafficDelayConfig.Enable {
		time.Sleep(time.Duration(c.cfg.TrafficDelayConfig.ValueInMilliSecond) * time.Millisecond)
	}

	return &pb.VirtualResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.VirtualResponse_Data{
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
			First:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Second:      uint32(utils.FakeInt(0, math.MaxUint32)),
			Third:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Fourth:      uint32(utils.FakeInt(0, math.MaxUint32)),
			Fifth:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Sixth:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Seventh:     uint32(utils.FakeInt(0, math.MaxUint32)),
			Eighth:      uint32(utils.FakeInt(0, math.MaxUint32)),
			Ninth:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Tenth:       uint32(utils.FakeInt(0, math.MaxUint32)),
			Secondaries: c.generateSecondaries(utils.FakeInt(c.cfg.SecondaryRange.Lower, c.cfg.SecondaryRange.Upper)),
		})
	}
	return primaries
}

func (c *Handler) generateSecondaries(quantity int) []*pb.Secondary {
	secondaries := make([]*pb.Secondary, 0)
	for _ = range quantity {
		secondaries = append(secondaries, &pb.Secondary{
			First:      uint64(utils.FakeInt(0, math.MaxInt/2)),
			Second:     uint64(utils.FakeInt(0, math.MaxInt/2)),
			Third:      uint64(utils.FakeInt(0, math.MaxInt/2)),
			Fourth:     uint64(utils.FakeInt(0, math.MaxInt/2)),
			Fifth:      uint64(utils.FakeInt(0, math.MaxInt/2)),
			Tertiaries: c.generateTertiary(utils.FakeInt(c.cfg.TertiaryRange.Lower, c.cfg.TertiaryRange.Upper)),
		})
	}
	return secondaries
}

func (c *Handler) generateTertiary(quantity int) []*pb.Tertiary {
	tertiaries := make([]*pb.Tertiary, 0)
	for _ = range quantity {
		tertiaries = append(tertiaries, &pb.Tertiary{
			First:        utils.FakeDouble(0, math.MaxFloat64),
			Second:       utils.FakeDouble(0, math.MaxFloat64),
			Third:        utils.FakeDouble(0, math.MaxFloat64),
			Fourth:       utils.FakeDouble(0, math.MaxFloat64),
			Fifth:        utils.FakeDouble(0, math.MaxFloat64),
			Quaternaries: c.generateQuaternaries(utils.FakeInt(c.cfg.QuaternaryRange.Lower, c.cfg.QuaternaryRange.Upper)),
		})
	}
	return tertiaries
}

func (c *Handler) generateQuaternaries(quantity int) []*pb.Quaternary {
	quaternaries := make([]*pb.Quaternary, 0)
	for _ = range quantity {
		quaternaries = append(quaternaries, &pb.Quaternary{
			First:  utils.FakeBoolean(),
			Second: utils.FakeBoolean(),
			Third:  utils.FakeBoolean(),
			Fourth: utils.FakeBoolean(),
			Fifth:  utils.FakeBoolean(),
		})
	}
	return quaternaries
}
