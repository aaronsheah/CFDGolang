package main

import (
	"net/http"

	"CFDGolang/pkg/onedimension"
	"CFDGolang/pkg/steps"

	"github.com/go-echarts/go-echarts/v2/components"
)

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

	charts := []components.Charter{
		steps.Step1Chart(oneDimensionConfig),
		steps.Step2Chart(oneDimensionConfig),
		steps.Step3Chart(oneDimensionConfig),
		steps.Step4Chart(oneDimensionConfig),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		page := components.NewPage()
		page.AddCharts(charts...)
		page.Render(w)
	})
	http.ListenAndServe(":1234", nil)
}
