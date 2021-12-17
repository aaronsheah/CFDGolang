package onedimension

type Config struct {
	gridPoints    int
	timesteps     int
	wavespeed     float64
	courantNumber float64
}

func NewConfig(
	gridPoints int, timesteps int, wavespeed float64, courantNumber float64,
) *Config {
	return &Config{
		gridPoints:    gridPoints,
		timesteps:     timesteps,
		wavespeed:     wavespeed,
		courantNumber: courantNumber,
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
	return float64(2) / float64(c.GridPoints()-1)
}

func (c *Config) TimeUnit() float64 {
	return c.DistanceUnit() * c.CourantNumber()
}
