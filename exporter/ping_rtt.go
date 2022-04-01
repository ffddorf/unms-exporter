package exporter

import (
	"math"
	"sort"
	"time"
)

// PingMetrics is a dumb data point computed from a list of PingResults.
type PingMetrics struct {
	PacketsSent int           // number of packets sent
	PacketsLost int           // number of packets lost
	Best        time.Duration // best RTT
	Worst       time.Duration // worst RTT
	Median      time.Duration // median RTT
	Mean        time.Duration // mean RTT
	StdDev      time.Duration // RTT std deviation
}

// PingResult stores the information about a single ping, in particular
// the round-trip time or whether the packet was lost.
type PingResult struct {
	RTT  time.Duration
	Lost bool
}

// PingHistory represents the ping history for a single node/device.
type PingHistory []PingResult

// NewHistory creates a new History object with a specific capacity.
func NewHistory(capacity int) PingHistory {
	return make(PingHistory, 0, capacity)
}

// AddResult saves a ping result into the internal history.
func (h *PingHistory) Add(rtt time.Duration, lost bool) {
	*h = append(*h, PingResult{RTT: rtt, Lost: lost})
}

// Compute aggregates the result history into a single data point.
func (h PingHistory) Compute() *PingMetrics {
	numFailure := 0
	numTotal := len(h)

	if numTotal == 0 {
		return nil
	}

	data := make([]float64, 0, numTotal)
	var best, worst, mean, stddev, total, sumSquares float64

	for _, curr := range h {
		if curr.Lost {
			numFailure++
			continue
		}

		rtt := curr.RTT.Seconds()
		if rtt < best || len(data) == 0 {
			best = rtt
		}
		if rtt > worst || len(data) == 0 {
			worst = rtt
		}
		data = append(data, rtt)
		total += rtt
	}

	size := float64(numTotal - numFailure)
	mean = total / size
	for _, rtt := range data {
		sumSquares += math.Pow(rtt-mean, 2)
	}
	stddev = math.Sqrt(sumSquares / size)

	median := math.NaN()
	if l := len(data); l > 0 {
		sort.Float64Slice(data).Sort()
		if l%2 == 0 {
			median = (data[l/2-1] + data[l/2]) / 2
		} else {
			median = data[l/2]
		}
	}

	return &PingMetrics{
		PacketsSent: numTotal,
		PacketsLost: numFailure,
		Best:        time.Duration(best * float64(time.Second)),
		Worst:       time.Duration(worst * float64(time.Second)),
		Median:      time.Duration(median * float64(time.Second)),
		Mean:        time.Duration(mean * float64(time.Second)),
		StdDev:      time.Duration(stddev * float64(time.Second)),
	}
}
