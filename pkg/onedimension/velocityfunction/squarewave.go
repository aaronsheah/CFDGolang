package velocityfunction

func SquareWave(gridPoints int, distanceUnit float64) []float64 {
	velocities := make([]float64, gridPoints)

	minIndexToSetToTwo := int(0.5 / distanceUnit)
	maxIndexToSetToTwo := int(1 / distanceUnit)

	for i := 0; i < gridPoints; i++ {
		if i >= minIndexToSetToTwo && i <= maxIndexToSetToTwo {
			velocities[i] = 2.0
		} else {
			velocities[i] = 1.0
		}
	}

	return velocities
}
