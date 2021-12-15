package main

import (
	"fmt"

	"CFDGolang/pkg/linearconvection"
)

func setupInitialVelocities(c *linearconvection.OneDimensionLinearConvectionConfig) []int {
	velocities := make([]int, c.GridPoints())

	minIndexToSetToTwo := int(0.5 / c.DistanceUnit())
	maxIndexToSetToTwo := int(1 / c.DistanceUnit())

	for i := 0; i < c.GridPoints(); i++ {
		if i >= minIndexToSetToTwo && i <= maxIndexToSetToTwo {
			velocities[i] = 2
		} else {
			velocities[i] = 1
		}
	}

	return velocities
}

func main() {
	gridPoints, timesteps, timeUnit, wavespeed := 41, 25, float64(0.025), 1
	config := linearconvection.NewOneDimensionLinearConvectionConfig(gridPoints, timesteps, timeUnit, wavespeed)

	velocities := setupInitialVelocities(config)

	fmt.Println(velocities)
}
