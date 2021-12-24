package steps

import (
	"CFDGolang/pkg/onedimension/convection"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"math"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type step1Config interface {
	convection.LinearConvectionConfig
}

func Step1Chart(c step1Config) components.Charter {
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
			Title: "Step 1: 1D Linear Convection",
		}),
	)

	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())

	linearConvection := convection.NewLinearConvection(c)
	velocitiesLinearConvection := linearConvection.Calculate(initialVelocities)

	initialVelocitiesLineData := make([]opts.LineData, len(initialVelocities))
	velocitiesLinearConvectionLineData := make([]opts.LineData, len(velocitiesLinearConvection))
	xAxisLabels := make([]string, int(math.Max(
		float64(len(initialVelocities)),
		float64(len(velocitiesLinearConvection)),
	)))

	for index, velocity := range initialVelocities {
		initialVelocitiesLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}
	for index, velocity := range velocitiesLinearConvection {
		velocitiesLinearConvectionLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}

	line.AddSeries("Initial Velocities", initialVelocitiesLineData)
	line.AddSeries("Linear Convection", velocitiesLinearConvectionLineData)
	line.SetXAxis(xAxisLabels)

	return line
}
