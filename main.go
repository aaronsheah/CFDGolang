package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"CFDGolang/pkg/onedimension"
	"CFDGolang/pkg/onedimension/burgers"
	"CFDGolang/pkg/onedimension/convection"
	"CFDGolang/pkg/onedimension/diffusion"
	"CFDGolang/pkg/onedimension/velocityfunction"
	"CFDGolang/pkg/step1"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func httpServer(
	config *onedimension.Config,
	titleToLineData map[string][]opts.LineData,
	xAxisLabels []string,
) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		page := components.NewPage()
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

		page.AddCharts(
			step1.Chart(config),
			line,
		)
		page.Render(w)
	}
}

func main() {
	const gridPoints, timesteps, wavespeed, courantNumber = 101, 100, 1.0, 0.5
	const viscosity, sigma = 0.07, 0.2
	oneDimensionConfig := onedimension.NewConfig(
		gridPoints,
		timesteps,
		wavespeed,
		courantNumber,
		viscosity,
		sigma,
	)

	linearConvection := convection.NewLinearConvection(oneDimensionConfig)
	nonLinearConvection := convection.NewNonLinearConvection(oneDimensionConfig)
	diffusion := diffusion.NewDiffusion(oneDimensionConfig)

	fmt.Println(gridPoints, timesteps, oneDimensionConfig.DistanceUnit(), viscosity, burgers.TimeUnit(oneDimensionConfig))
	burgersEquation := burgers.NewBurgers(oneDimensionConfig)

	velocities := velocityfunction.SawTooth(gridPoints, 0, viscosity)
	velocitiesLinearConvection := linearConvection.Calculate(velocities)
	velocitiesNonLinearConvection := nonLinearConvection.Calculate(velocities)
	velocitiesDiffusion := diffusion.Calculate(velocities)
	velocitiesBurgers := burgersEquation.Calculate(velocities)
	velocitiesAnalytical := velocityfunction.SawTooth(
		gridPoints,
		float64(timesteps)*burgers.TimeUnit(oneDimensionConfig),
		viscosity,
	)

	fmt.Println("velocities", velocities)
	fmt.Println("velocitiesLinearConvection", velocitiesLinearConvection)
	fmt.Println("velocitiesNonLinearConvection", velocitiesNonLinearConvection)
	fmt.Println("velocitiesDiffusion", velocitiesDiffusion)
	fmt.Println("velocitiesBurgers", velocitiesBurgers)
	fmt.Println("velocitiesAnalytical", velocitiesAnalytical)

	initialVelocitiesLineData := make([]opts.LineData, len(velocities))
	velocitiesLinearConvectionLineData := make([]opts.LineData, len(velocities))
	velocitiesNonLinearConvectionLineData := make([]opts.LineData, len(velocities))
	velocitiesDiffusionLineData := make([]opts.LineData, len(velocities))
	velocitiesBurgersLineData := make([]opts.LineData, len(velocities))
	velocitiesAnalyticalLineData := make([]opts.LineData, len(velocities))

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
		velocitiesDiffusionLineData[i] = opts.LineData{
			Value: velocitiesDiffusion[i],
		}
		velocitiesBurgersLineData[i] = opts.LineData{
			Value: velocitiesBurgers[i],
		}
		velocitiesAnalyticalLineData[i] = opts.LineData{
			Value: velocitiesAnalytical[i],
		}
		xAxisLabels[i] = strconv.FormatFloat(float64(i)/float64(2*math.Pi), 'f', 2, 64)
	}

	http.HandleFunc("/", httpServer(oneDimensionConfig, map[string][]opts.LineData{
		"Intitial Velocities":  initialVelocitiesLineData,
		"Analyical Velocities": velocitiesAnalyticalLineData,
		"1D Linear Convection": velocitiesLinearConvectionLineData,
		// "1D Non Linear Convection": velocitiesNonLinearConvectionLineData,
		"1D Diffusion": velocitiesDiffusionLineData,
		"Burgers":      velocitiesBurgersLineData,
	}, xAxisLabels))
	http.ListenAndServe(":1234", nil)
}
