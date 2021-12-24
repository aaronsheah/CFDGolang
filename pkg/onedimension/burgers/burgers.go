package burgers

type burgers struct {
	timesteps      int
	vDtByDxSquared float64
	dtByDx         float64
}

func NewBurgers(config BurgersConfig) *burgers {
	dtByDx := TimeUnit(config) / config.DistanceUnit()
	vDtByDxSquared := config.Viscosity() * dtByDx / config.DistanceUnit()
	return &burgers{
		timesteps:      config.Timesteps(),
		dtByDx:         dtByDx,
		vDtByDxSquared: vDtByDxSquared,
	}
}

func (b *burgers) singleTimestep(velocities []float64) []float64 {
	nextVelocities := make([]float64, len(velocities))
	for i := 1; i < len(velocities)-1; i++ {
		nextVelocities[i] = velocities[i] - b.convectionTerm(velocities, i, i-1) + b.diffusionTerm(velocities, i, i-1, i+1)
	}

	nextVelocities[0] = velocities[0] -
		b.convectionTerm(velocities, 0, len(nextVelocities)-2) +
		b.diffusionTerm(velocities, 0, len(nextVelocities)-2, 1)
	nextVelocities[len(nextVelocities)-1] = nextVelocities[0]
	return nextVelocities
}

func (b *burgers) convectionTerm(velocities []float64, i int, iMinus1 int) float64 {
	return velocities[i] * b.dtByDx * (velocities[i] - velocities[iMinus1])
}

func (b *burgers) diffusionTerm(velocities []float64, i int, iMinus1 int, iPlus1 int) float64 {
	return b.vDtByDxSquared * (velocities[iPlus1] - 2*velocities[i] + velocities[iMinus1])
}

func (b *burgers) Calculate(velocities []float64) []float64 {
	output := make([]float64, len(velocities))
	copy(output, velocities)

	for t := 0; t < b.timesteps; t++ {
		output = b.singleTimestep(output)
	}
	return output
}
