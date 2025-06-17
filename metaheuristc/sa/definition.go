package sa

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

const name = "SA"

type Configuration struct {
	MaxIterations      int
	TimeLimitSeconds   float64
	Alpha              float64
	ChangeImpact       float64
	TemperatureInitial float64
	TemperatureGoal    float64
	TemperatureReheat  float64
	ShakeMin           float64
	ShakeMax           float64
	QtdReheat          uint8
	Iterations         int
}

type SimulatedAnnealing struct {
	env           definition.Environment
	configuration *Configuration
	logger        *logger.Log
	search        search.Local
	RG            *random.Generator
	solutionPool  *solution.Pool
}
