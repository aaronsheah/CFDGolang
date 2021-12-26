package steps

import (
	"CFDGolang/pkg/onedimension/convection"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
)

type step1Config interface {
	convection.LinearConvectionConfig
}

func Step1Chart(c step1Config) components.Charter {
	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())

	linearConvection := convection.NewLinearConvection(c)
	velocitiesLinearConvection := linearConvection.Calculate(initialVelocities)

	xAxisLabels := make([]string, len(initialVelocities))
	for index := range initialVelocities {
		xAxisLabels[index] = strconv.Itoa(index)
	}

	return createLineChart(
		"Step 1: 1D Linear Convection",
		map[string][]float64{
			"Initial Velocities": initialVelocities,
			"Linear Convection":  velocitiesLinearConvection,
		},
		xAxisLabels,
	)
}
