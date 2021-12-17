package convection

type nonLinearConvectionConfig interface {
	GridPoints() int
	Timesteps() int
	TimeUnit() float64
	DistanceUnit() float64
}
