package steps

import (
	"CFDGolang/pkg/onedimension/burgers"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
)

type step4Config interface {
	burgers.BurgersConfig
	GridPoints() int
}

func Step4Chart(c step4Config) components.Charter {
	burgersEquation := burgers.NewBurgers(c)

	initialVelocities := velocityfunction.SawTooth(c.GridPoints(), 0, c.Viscosity())
	velocitiesBurgers := burgersEquation.Calculate(initialVelocities)
	velocitiesAnalytical := velocityfunction.SawTooth(
		c.GridPoints(),
		float64(c.Timesteps())*burgers.TimeUnit(c),
		c.Viscosity(),
	)

	xAxisLabels := make([]string, len(initialVelocities))
	for index := range initialVelocities {
		xAxisLabels[index] = strconv.FormatFloat(float64(index)*2/float64(len(initialVelocities)), 'f', 2, 64) + "π"
	}

	return createLineChart(
		"Step 4: Burgers' Equation",
		map[string][]float64{
			"Initial Velocities":  initialVelocities,
			"Burgers":             velocitiesBurgers,
			"Analytical Solution": velocitiesAnalytical,
		},
		xAxisLabels,
	)
}
