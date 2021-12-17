package linearconvection

type OneDimensionLinearConvectionConfig struct {
	gridPoints   int
	timesteps    int
	distanceUnit float64
	timeUnit     float64
	waveSpeed    int
}

func NewOneDimensionLinearConvectionConfig(gridPoints int, timesteps int, wavespeed int, courantNumber float64) *OneDimensionLinearConvectionConfig {
	distanceUnit := float64(2) / float64(gridPoints-1)
	return &OneDimensionLinearConvectionConfig{
		gridPoints:   gridPoints,
		timesteps:    timesteps,
		distanceUnit: distanceUnit,
		timeUnit:     courantNumber * distanceUnit,
		waveSpeed:    wavespeed,
	}
}

func (config *OneDimensionLinearConvectionConfig) GridPoints() int {
	return config.gridPoints
}

func (config *OneDimensionLinearConvectionConfig) DistanceUnit() float64 {
	return config.distanceUnit
}

func (config *OneDimensionLinearConvectionConfig) Timesteps() int {
	return config.timesteps
}

func (config *OneDimensionLinearConvectionConfig) TimeUnit() float64 {
	return config.timeUnit
}

func (config *OneDimensionLinearConvectionConfig) WaveSpeed() int {
	return config.waveSpeed
}

func (config *OneDimensionLinearConvectionConfig) WavespeedDtDx() float64 {
	return float64(config.WaveSpeed()) * config.TimeUnit() / config.DistanceUnit()
}
