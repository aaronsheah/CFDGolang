package linearconvection

type OneDimensionLinearConvectionConfig struct {
	gridPoints int
	timesteps  int
	timeUnit   float64
	waveSpeed  int
}

func NewOneDimensionLinearConvectionConfig(gridPoints int, timesteps int, timeUnit float64, wavespeed int) *OneDimensionLinearConvectionConfig {
	return &OneDimensionLinearConvectionConfig{
		gridPoints: gridPoints,
		timesteps:  timesteps,
		timeUnit:   timeUnit,
		waveSpeed:  wavespeed,
	}
}

func (config *OneDimensionLinearConvectionConfig) GridPoints() int {
	return config.gridPoints
}

func (config *OneDimensionLinearConvectionConfig) DistanceUnit() float64 {
	return float64(2) / float64(config.gridPoints-1)
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
