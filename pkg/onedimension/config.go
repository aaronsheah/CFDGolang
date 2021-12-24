package onedimension

import "math"

type Config struct {
	gridPoints    int
	timesteps     int
	wavespeed     float64
	courantNumber float64
	viscosity     float64
	sigma         float64
}

func NewConfig(
	gridPoints int,
	timesteps int,
	wavespeed float64,
	courantNumber float64,
	viscosity float64,
	sigma float64,
) *Config {
	return &Config{
		gridPoints:    gridPoints,
		timesteps:     timesteps,
		wavespeed:     wavespeed,
		courantNumber: courantNumber,
		viscosity:     viscosity,
		sigma:         sigma,
	}
}

func (c *Config) GridPoints() int {
	return c.gridPoints
}

func (c *Config) Timesteps() int {
	return c.timesteps
}

func (c *Config) Wavespeed() float64 {
	return c.wavespeed
}

func (c *Config) CourantNumber() float64 {
	return c.courantNumber
}

func (c *Config) DistanceUnit() float64 {
	return float64(2) * math.Pi / float64(c.GridPoints()-1)
}

func (c *Config) Viscosity() float64 {
	return c.viscosity
}

func (c *Config) Sigma() float64 {
	return c.sigma
}
