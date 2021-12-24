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

type step2Config interface {
	convection.NonLinearConvectionConfig
}

func Step2Chart(c step2Config) components.Charter {
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
			Title: "Step 2: 1D Non-Linear Convection",
		}),
	)

	initialVelocities := velocityfunction.SquareWave(c.GridPoints(), c.DistanceUnit())
	nonLinearConvection := convection.NewNonLinearConvection(c)
	velocitiesNonLinearConvection := nonLinearConvection.Calculate(initialVelocities)

	initialVelocitiesLineData := make([]opts.LineData, len(initialVelocities))
	velocitiesNonLinearConvectionLineData := make([]opts.LineData, len(velocitiesNonLinearConvection))
	xAxisLabels := make([]string, int(math.Max(
		float64(len(initialVelocities)),
		float64(len(velocitiesNonLinearConvection)),
	)))

	for index, velocity := range initialVelocities {
		initialVelocitiesLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}
	for index, velocity := range velocitiesNonLinearConvection {
		velocitiesNonLinearConvectionLineData[index] = opts.LineData{Value: velocity}
		xAxisLabels[index] = strconv.Itoa(index)
	}

	line.AddSeries("Initial Velocities", initialVelocitiesLineData)
	line.AddSeries("Non-Linear Convection", velocitiesNonLinearConvectionLineData)
	line.SetXAxis(xAxisLabels)

	return line
}
