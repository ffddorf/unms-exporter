package exporter

import (
	"github.com/ffddorf/unms-exporter/models"
)

// weightedMean condenses a list of models.CoordinatesXY down to a single
// value. It computes a weighted arith. mean of all data points between
// tmin and tmax, giving values <= tmin a weight of 0, values >= tmax a
// weight of 1, and using a linear interpolation between these points.
//
//	weight  ↑
//	     1 -+              ··
//	        |            ·
//	        |          ·
//	        |        ·
//	     0 -+ ······
//	        |------+-------+-→ time
//	               tmin    tmax
//
// The typical differenece between tmin and tmax represents a time span
// of 10 minutes, and tmax represents the current timestamp.
func weightedMean(tmin, tmax float64, list models.ListOfCoordinates) (avg float64) {
	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	// f(x) = mx + b, with b = 0 and slope m = 1/Δx
	slope := 1.0 / (tmax - tmin)

	// value and weight accumulator: avg = Σᵢ wᵢ·vᵢ / Σᵢ wᵢ
	var numerator, denominator float64

	for _, xy := range list {
		t, val := xy.X, xy.Y
		if t < tmin {
			continue
		}
		weight := 1.0
		if t < tmax {
			weight = slope * (t - tmin)
		}
		numerator += weight * val
		denominator += weight
	}

	if denominator <= 0 {
		return 0
	}
	return numerator / denominator
}
