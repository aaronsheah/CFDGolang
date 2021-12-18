package burgers

type burgersConfig interface {
	Timesteps() int
	DistanceUnit() float64
	Viscosity() float64
}

func TimeUnit(c burgersConfig) float64 {
	return c.DistanceUnit() * c.Viscosity()
}
