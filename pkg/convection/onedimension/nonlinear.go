package onedimension

type nonLinearConvection struct {
	gridPoints int
	timesteps  int
	dtByDx     float64
}

func NewNonLinearConvection(config nonLinearConvectionConfig) *nonLinearConvection {
	return &nonLinearConvection{
		gridPoints: config.GridPoints(),
		timesteps:  config.Timesteps(),
		dtByDx:     config.TimeUnit() / config.DistanceUnit(),
	}
}

func (nlc *nonLinearConvection) singleTimestep(velocities []float64) []float64 {
	nextVelocities := make([]float64, len(velocities))
	nextVelocities[0] = velocities[0]
	for i := 1; i < nlc.gridPoints; i++ {
		nextVelocities[i] = velocities[i] - velocities[i]*nlc.dtByDx*(velocities[i]-velocities[i-1])
	}
	return nextVelocities
}

func (nlc *nonLinearConvection) Calculate(velocities []float64) []float64 {
	for t := 0; t < nlc.timesteps; t++ {
		velocities = nlc.singleTimestep(velocities)
	}
	return velocities
}
