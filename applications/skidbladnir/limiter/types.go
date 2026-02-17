package limiter

type Capacity struct {
	RPS   float64
	Burst int32
}

type Status struct {
	LeaderIP    string
	PodIP       string
	GlobalRps   float64
	GlobalBurst int32
}
