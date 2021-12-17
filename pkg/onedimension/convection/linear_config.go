package convection

type linearConvectionConfig interface {
	GridPoints() int
	Timesteps() int
	Wavespeed() float64
	TimeUnit() float64
	DistanceUnit() float64
}
