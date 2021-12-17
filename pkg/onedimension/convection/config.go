package convection

type convectionConfig interface {
	DistanceUnit() float64
	CourantNumber() float64
}

func timeUnit(config convectionConfig) float64 {
	return config.DistanceUnit() * config.CourantNumber()
}
