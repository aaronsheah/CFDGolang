package steps

import (
	"CFDGolang/pkg/onedimension/diffusion"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"math"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type step3Config interface {
	diffusion.DiffusionConfig
}

func Step3Chart(c step3Config) components.Charter {
	line := charts.NewLine()
	line.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "insideRight",
		}),
		charts.WithLineChartOpts(opts.LineChart{
			Smooth: false,
			Step:   true,
		}),
	)

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Step 3: 1D Diffusion",
		}),
	)

	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())
	initialVelocitiesLineData := make([]opts.LineData, len(initialVelocities))

	diffusion := diffusion.NewDiffusion(c)
	velocitiesDiffusion := diffusion.Calculate(initialVelocities)
	velocitiesDiffusionLineData := make([]opts.LineData, len(velocitiesDiffusion))

	xAxisLabels := make([]string, int(math.Max(
		float64(len(initialVelocities)),
		float64(len(velocitiesDiffusion)),
	)))

	for index, velocity := range initialVelocities {
		initialVelocitiesLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}
	for index, velocity := range velocitiesDiffusion {
		velocitiesDiffusionLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}

	line.AddSeries("Initial Velocities", initialVelocitiesLineData)
	line.AddSeries("Non-Linear Convection", velocitiesDiffusionLineData)
	line.SetXAxis(xAxisLabels)

	return line
}
