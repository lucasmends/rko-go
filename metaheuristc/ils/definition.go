package ils

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

const name = "ILS"

type Configuration struct {
	MaxIterations       int
	TimeLimitSeconds    float64
	ShakeMin            float64
	ShakeMax            float64
	MetropolisCriterion bool
}

type ILS struct {
	env           definition.Environment
	configuration *Configuration
	logger        *logger.Log
	search        search.Local
	RG            *random.Generator
	solutionPool  *solution.Pool
}
