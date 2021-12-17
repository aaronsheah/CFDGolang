package diffusion

import (
	"math"
)

type diffusion struct {
	gridPoints     int
	timesteps      int
	vDtByDxSquared float64
}

func NewDiffusion(c diffusionConfig) *diffusion {
	return &diffusion{
		gridPoints:     c.GridPoints(),
		timesteps:      c.Timesteps(),
		vDtByDxSquared: c.Viscosity() * c.TimeUnit() / math.Pow(c.DistanceUnit(), 2),
	}
}

func (d *diffusion) singleTimestep(velocities []float64) []float64 {
	nextVelocities := make([]float64, len(velocities))
	nextVelocities[0] = velocities[0]
	for i := 1; i < d.gridPoints-1; i++ {
		nextVelocities[i] = velocities[i] + d.vDtByDxSquared*(velocities[i+1]-2*velocities[i]+velocities[i-1])
	}
	nextVelocities[d.gridPoints-1] = velocities[d.gridPoints-1]
	return nextVelocities
}

func (d *diffusion) Calculate(velocities []float64) []float64 {
	for t := 0; t < d.timesteps; t++ {
		velocities = d.singleTimestep(velocities)
	}
	return velocities
}
