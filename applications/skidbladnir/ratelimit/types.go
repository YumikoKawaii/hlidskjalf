package ratelimit

type TokenGrant struct {
	Granted int
	RPS     float64
}

type Status struct {
	IsLeader  bool
	LeaderIP  string
	PodIP     string
	GlobalRPS float64
}
