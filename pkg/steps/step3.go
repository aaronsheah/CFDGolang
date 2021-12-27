package steps

import (
	"CFDGolang/pkg/onedimension/diffusion"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
)

type step3Config interface {
	diffusion.DiffusionConfig
}

func Step3Chart(c step3Config) components.Charter {
	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())

	diffusion := diffusion.NewDiffusion(c)
	velocitiesDiffusion := diffusion.Calculate(initialVelocities)

	xAxisLabels := make([]string, len(initialVelocities))
	for index := range initialVelocities {
		xAxisLabels[index] = strconv.Itoa(index)
	}

	return createLineChart(
		"Step 3: 1D Diffusion",
		map[string][]float64{
			"Initial Velocities": initialVelocities,
			"Diffusion":          velocitiesDiffusion,
		},
		xAxisLabels,
	)
}
