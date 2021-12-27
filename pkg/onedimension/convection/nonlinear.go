package convection

type nonLinearConvection struct {
	gridPoints int
	timesteps  int
	dtByDx     float64
}

func NewNonLinearConvection(config NonLinearConvectionConfig) *nonLinearConvection {
	return &nonLinearConvection{
		gridPoints: config.GridPoints(),
		timesteps:  config.Timesteps(),
		dtByDx:     timeUnit(config) / config.DistanceUnit(),
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
	output := make([]float64, len(velocities))
	copy(output, velocities)

	for t := 0; t < nlc.timesteps; t++ {
		output = nlc.singleTimestep(output)
	}
	return output
}
