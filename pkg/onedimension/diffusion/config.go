package diffusion

import "math"

type DiffusionConfig interface {
	GridPoints() int
	Timesteps() int
	Viscosity() float64
	DistanceUnit() float64
	Sigma() float64
}

func timeUnit(c DiffusionConfig) float64 {
	return c.Sigma() * math.Pow(c.DistanceUnit(), 2.0) / c.Viscosity()
}
