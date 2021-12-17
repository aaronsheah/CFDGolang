package onedimension

type linearConvection struct {
	gridPoints      int
	timesteps       int
	wavespeedDtByDx float64
}

func NewLinearConvection(config linearConvectionConfig) *linearConvection {
	return &linearConvection{
		gridPoints:      config.GridPoints(),
		timesteps:       config.Timesteps(),
		wavespeedDtByDx: config.Wavespeed() * config.TimeUnit() / config.DistanceUnit(),
	}
}

func (lc *linearConvection) singleTimestep(velocities []float64) []float64 {
	nextVelocities := make([]float64, len(velocities))
	nextVelocities[0] = velocities[0]
	for i := 1; i < lc.gridPoints; i++ {
		nextVelocities[i] = velocities[i] - lc.wavespeedDtByDx*(velocities[i]-velocities[i-1])
	}
	return nextVelocities
}

func (lc *linearConvection) Calculate(velocities []float64) []float64 {
	for t := 0; t < lc.timesteps; t++ {
		velocities = lc.singleTimestep(velocities)
	}
	return velocities
}
