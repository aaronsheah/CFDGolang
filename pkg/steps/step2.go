package steps

import (
	"CFDGolang/pkg/onedimension/convection"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
)

type step2Config interface {
	convection.NonLinearConvectionConfig
}

func Step2Chart(c step2Config) components.Charter {
	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())
	nonLinearConvection := convection.NewNonLinearConvection(c)
	velocitiesNonLinearConvection := nonLinearConvection.Calculate(initialVelocities)

	xAxisLabels := make([]string, len(initialVelocities))

	for index := range initialVelocities {
		xAxisLabels[index] = strconv.Itoa(index)
	}

	return createLineChart(
		"Step 2: 1D Non-Linear Convection",
		map[string][]float64{
			"Initial Velocities":    initialVelocities,
			"Non-Linear Convection": velocitiesNonLinearConvection,
		},
		xAxisLabels,
	)
}
