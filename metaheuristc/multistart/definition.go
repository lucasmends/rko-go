package multistart

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

const name = "MultiStart"

type Configuration struct {
	MaxIterations    int
	TimeLimitSeconds float64
}

type MultiStart struct {
	env           definition.Environment
	configuration *Configuration
	logger        *logger.Log
	search        search.Local
	RG            *random.Generator
	solutionPool  *solution.Pool
}
