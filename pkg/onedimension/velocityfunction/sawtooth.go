package velocityfunction

import (
	"fmt"
	"math"
)

func SawTooth(gridPoints int, time float64, viscosity float64) []float64 {
	velocities := make([]float64, gridPoints)
	for i := 0; i < len(velocities); i++ {
		x := float64(i) / float64(gridPoints) * (2 * math.Pi)
		velocities[i] = 4 - 2*viscosity*dPhiByDxDividedByPhi(
			x,
			time,
			float64(viscosity),
		)
		fmt.Println("dPhiByDxDividedByPhi", dPhiByDxDividedByPhi(
			x,
			time,
			float64(viscosity),
		))
	}
	return velocities
}

func dPhiByDxDividedByPhi(x float64, time float64, viscosity float64) float64 {
	denominator := 2.0 * viscosity * (time + 1.0)

	first := x - 4*time
	second := first - 2*math.Pi

	firstExponent := 1 / math.Exp(first*first/(2*denominator))
	secondExponent := 1 / math.Exp(second*second/(2*denominator))

	phi := firstExponent + secondExponent

	return -(first*firstExponent + second*secondExponent) / denominator / phi
}
