package simple_packet

import "time"

type RttStats struct {
	smoothedRtt time.Duration
	latestRtt   time.Duration
}

func NewRttStats() *RttStats {
	return &RttStats{
		smoothedRtt: 0,
		latestRtt:   0,
	}
}

func (stats *RttStats) NewRtt(rtt time.Duration) {
	if stats.smoothedRtt == 0 {
		stats.smoothedRtt = rtt
	} else {
		stats.smoothedRtt = stats.smoothedRtt*7/8 + rtt/8
	}
	stats.latestRtt = rtt
}
