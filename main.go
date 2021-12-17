package main

import (
	"fmt"
	"net/http"
	"strconv"

	"CFDGolang/pkg/convection/onedimension"
	"CFDGolang/pkg/nonlinearconvection"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func setupInitialVelocities(c *onedimension.Config) []float64 {
	velocities := make([]float64, c.GridPoints())

	minIndexToSetToTwo := int(0.5 / c.DistanceUnit())
	maxIndexToSetToTwo := int(1 / c.DistanceUnit())

	for i := 0; i < c.GridPoints(); i++ {
		if i >= minIndexToSetToTwo && i <= maxIndexToSetToTwo {
			velocities[i] = 2.0
		} else {
			velocities[i] = 1.0
		}
	}

	return velocities
}

func httpServer(
	titleToLineData map[string][]opts.LineData,
	xAxisLabels []string,
) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		line := charts.NewLine()

		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
			charts.WithTitleOpts(opts.Title{
				Title:    "Line example in Westeros theme",
				Subtitle: "Line chart rendered by the http server this time",
			}))

		line.SetXAxis(xAxisLabels)
		for title, lineData := range titleToLineData {
			line.AddSeries(title, lineData)
		}
		line.SetSeriesOptions(
			charts.WithLineChartOpts((opts.LineChart{
				Smooth: true,
				Step:   true,
			})),
		)

		line.Render(w)
	}
}

func main() {
	const gridPoints, timesteps, wavespeed, courantNumber = 85, 25, 1.0, 0.5
	config := onedimension.NewOneDimensionConfig(
		gridPoints,
		timesteps,
		wavespeed,
		courantNumber,
	)

	oneDimensionalLinearConvection := onedimension.NewOneDimensionLinearConvection(config)
	oneDimensionalNonLinearConvection := nonlinearconvection.OneDimensionNonLinearConvection{
		Config: *nonlinearconvection.NewOneDimensionNonLinearConvectionConfig(gridPoints, timesteps, courantNumber),
	}

	velocities := setupInitialVelocities(config)
	velocitiesLinearConvection := oneDimensionalLinearConvection.Calculate(velocities)
	velocitiesNonLinearConvection := oneDimensionalNonLinearConvection.Calculate(velocities)
	fmt.Println("velocities", velocities)
	fmt.Println("velocitiesLinearConvection", velocitiesLinearConvection)
	fmt.Println("velocitiesNonLinearConvection", velocitiesNonLinearConvection)

	initialVelocitiesLineData := make([]opts.LineData, len(velocities))
	velocitiesLinearConvectionLineData := make([]opts.LineData, len(velocities))
	velocitiesNonLinearConvectionLineData := make([]opts.LineData, len(velocities))

	xAxisLabels := make([]string, len(velocities))
	for i := 0; i < len(velocities); i++ {
		initialVelocitiesLineData[i] = opts.LineData{
			Value: velocities[i],
		}
		velocitiesLinearConvectionLineData[i] = opts.LineData{
			Value: velocitiesLinearConvection[i],
		}
		velocitiesNonLinearConvectionLineData[i] = opts.LineData{
			Value: velocitiesNonLinearConvection[i],
		}
		xAxisLabels[i] = strconv.FormatFloat(float64(i)/float64(len(velocities)), 'f', 2, 64)
	}

	http.HandleFunc("/", httpServer(map[string][]opts.LineData{
		"Intitial Velocities":      initialVelocitiesLineData,
		"1D Linear Convection":     velocitiesLinearConvectionLineData,
		"1D Non Linear Convection": velocitiesNonLinearConvectionLineData,
	}, xAxisLabels))
	http.ListenAndServe(":1234", nil)
}
