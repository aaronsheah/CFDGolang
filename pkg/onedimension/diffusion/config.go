package diffusion

import "math"

type Config struct {
	gridPoints int
	timesteps  int
	viscosity  float64
	sigma      float64
}

type diffusionConfig interface {
	GridPoints() int
	Timesteps() int
	Viscosity() float64
	TimeUnit() float64
	DistanceUnit() float64
}

func NewConfig(
	gridPoints int, timesteps int, viscosity float64, sigma float64,
) diffusionConfig {
	return &Config{
		gridPoints: gridPoints,
		timesteps:  timesteps,
		viscosity:  viscosity,
		sigma:      sigma,
	}
}

func (c *Config) GridPoints() int {
	return c.gridPoints
}

func (c *Config) Timesteps() int {
	return c.timesteps
}

func (c *Config) Viscosity() float64 {
	return c.viscosity
}

func (c *Config) Sigma() float64 {
	return c.sigma
}

func (c *Config) DistanceUnit() float64 {
	return float64(2) / float64(c.GridPoints()-1)
}

func (c *Config) TimeUnit() float64 {
	return c.Sigma() * math.Pow(c.DistanceUnit(), 2.0) / c.Viscosity()
}
