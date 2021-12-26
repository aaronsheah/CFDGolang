package steps

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func initLineChart(title string) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
	)
	return line
}

func createLineChart(title string, titleToData map[string][]float64, xAxisLabels []string) *charts.Line {
	line := initLineChart(title)
	line.SetXAxis(xAxisLabels)

	for title, data := range titleToData {
		lineData := make([]opts.LineData, len(data))
		for index, velocity := range data {
			lineData[index] = opts.LineData{Value: velocity}
		}
		line.AddSeries(title, lineData)
	}

	line.SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			Smooth: true,
			Step:   true,
		}),
	)

	return line
}
