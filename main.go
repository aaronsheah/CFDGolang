package main

import (
	"fmt"
	"net/http"
	"strconv"

	"CFDGolang/pkg/linearconvection"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func setupInitialVelocities(c *linearconvection.OneDimensionLinearConvectionConfig) []float64 {
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
	lineDatum []opts.LineData,
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
		line.AddSeries("Test", lineDatum)
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
	gridPoints, timesteps, timeUnit, wavespeed := 41, 25, float64(0.025), 1
	oneDimensionalLinearConvection := linearconvection.OneDimensionLinearConvection{
		Config: *linearconvection.NewOneDimensionLinearConvectionConfig(gridPoints, timesteps, timeUnit, wavespeed),
	}
	velocities := setupInitialVelocities(&oneDimensionalLinearConvection.Config)
	velocities = oneDimensionalLinearConvection.Calculate(velocities)
	fmt.Println(velocities)

	points := make([]opts.LineData, len(velocities))
	xAxisLabels := make([]string, len(velocities))
	for i := 0; i < len(velocities); i++ {
		points[i] = opts.LineData{
			Value: velocities[i],
		}
		xAxisLabels[i] = strconv.Itoa(i)
	}

	http.HandleFunc("/", httpServer(points, xAxisLabels))
	http.ListenAndServe(":1234", nil)
}
