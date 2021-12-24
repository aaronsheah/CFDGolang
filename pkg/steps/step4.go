package steps

import (
	"CFDGolang/pkg/onedimension/burgers"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"math"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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

	initialVelocitiesLineData := make([]opts.LineData, len(initialVelocities))
	velocitiesBurgersLineData := make([]opts.LineData, len(velocitiesBurgers))
	velocitiesAnalyticalLineData := make([]opts.LineData, len(velocitiesAnalytical))

	for index, velocity := range initialVelocities {
		initialVelocitiesLineData[index] = opts.LineData{Value: velocity}
	}
	for index, velocity := range velocitiesBurgers {
		velocitiesBurgersLineData[index] = opts.LineData{Value: velocity}
	}
	for index, velocity := range velocitiesAnalytical {
		velocitiesAnalyticalLineData[index] = opts.LineData{Value: velocity}
	}

	xAxisLabels := make([]string, int(math.Max(
		float64(len(initialVelocities)),
		float64(len(velocitiesBurgers)),
	)))

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Step 4: Burgers' Equation",
		}),
	)

	line.SetXAxis(xAxisLabels)
	line.AddSeries("Initial Velocities", initialVelocitiesLineData)
	line.AddSeries("Burgers", velocitiesBurgersLineData)
	line.AddSeries("Analytical Solution", velocitiesAnalyticalLineData)

	line.SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			Smooth: true,
			Step:   true,
		}),
	)

	return line
}
