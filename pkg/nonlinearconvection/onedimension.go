package nonlinearconvection

type OneDimensionNonLinearConvection struct {
	Config OneDimensionNonLinearConvectionConfig
}

func (odlc *OneDimensionNonLinearConvection) singleTimestep(velocities []float64) []float64 {
	nextVelocities := make([]float64, len(velocities))
	nextVelocities[0] = velocities[0]
	for i := 1; i < odlc.Config.GridPoints(); i++ {
		nextVelocities[i] = velocities[i] - velocities[i]*odlc.Config.TimeUnit()/odlc.Config.DistanceUnit()*(velocities[i]-velocities[i-1])
	}
	return nextVelocities
}

func (odlc *OneDimensionNonLinearConvection) Calculate(velocities []float64) []float64 {
	for t := 0; t < odlc.Config.Timesteps(); t++ {
		velocities = odlc.singleTimestep(velocities)
	}
	return velocities
}
