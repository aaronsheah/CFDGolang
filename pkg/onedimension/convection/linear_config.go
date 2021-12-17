package convection

type linearConvectionConfig interface {
	convectionConfig
	GridPoints() int
	Timesteps() int
	Wavespeed() float64
}
