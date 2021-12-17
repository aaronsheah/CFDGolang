package nonlinearconvection

type OneDimensionNonLinearConvectionConfig struct {
	gridPoints int
	timesteps  int
	timeUnit   float64
}

func NewOneDimensionNonLinearConvectionConfig(gridPoints int, timesteps int, timeUnit float64) *OneDimensionNonLinearConvectionConfig {
	return &OneDimensionNonLinearConvectionConfig{
		gridPoints: gridPoints,
		timesteps:  timesteps,
		timeUnit:   timeUnit,
	}
}

func (config *OneDimensionNonLinearConvectionConfig) GridPoints() int {
	return config.gridPoints
}

func (config *OneDimensionNonLinearConvectionConfig) DistanceUnit() float64 {
	return float64(2) / float64(config.gridPoints-1)
}

func (config *OneDimensionNonLinearConvectionConfig) Timesteps() int {
	return config.timesteps
}

func (config *OneDimensionNonLinearConvectionConfig) TimeUnit() float64 {
	return config.timeUnit
}
