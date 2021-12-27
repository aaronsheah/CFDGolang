package burgers

type BurgersConfig interface {
	Timesteps() int
	DistanceUnit() float64
	Viscosity() float64
}

func TimeUnit(c BurgersConfig) float64 {
	return c.DistanceUnit() * c.Viscosity()
}
