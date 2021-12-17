package convection

type nonLinearConvectionConfig interface {
	convectionConfig
	GridPoints() int
	Timesteps() int
}
