package vns

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/solution"
	"rko/random"
)

const name = "VNS"

type Configuration struct {
	MaxIterations    int
	TimeLimitSeconds float64
	Rate             float64
}

type VNS struct {
	env           definition.Environment
	configuration *Configuration
	logger        *logger.Log
	RG            *random.Generator
	solutionPool  *solution.Pool
}
