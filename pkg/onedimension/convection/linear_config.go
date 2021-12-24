package convection

type LinearConvectionConfig interface {
	convectionConfig
	GridPoints() int
	Timesteps() int
	Wavespeed() float64
}
