package convection

type NonLinearConvectionConfig interface {
	convectionConfig
	GridPoints() int
	Timesteps() int
}
