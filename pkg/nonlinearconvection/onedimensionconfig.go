package nonlinearconvection

type OneDimensionNonLinearConvectionConfig struct {
	gridPoints   int
	timesteps    int
	distanceUnit float64
	timeUnit     float64
}

func NewOneDimensionNonLinearConvectionConfig(gridPoints int, timesteps int, courantNumber float64) *OneDimensionNonLinearConvectionConfig {
	distanceUnit := float64(2) / float64(gridPoints-1)
	return &OneDimensionNonLinearConvectionConfig{
		gridPoints:   gridPoints,
		timesteps:    timesteps,
		distanceUnit: distanceUnit,
		timeUnit:     courantNumber * distanceUnit,
	}
}

func (config *OneDimensionNonLinearConvectionConfig) GridPoints() int {
	return config.gridPoints
}

func (config *OneDimensionNonLinearConvectionConfig) DistanceUnit() float64 {
	return config.distanceUnit
}

func (config *OneDimensionNonLinearConvectionConfig) Timesteps() int {
	return config.timesteps
}

func (config *OneDimensionNonLinearConvectionConfig) TimeUnit() float64 {
	return config.timeUnit
}
