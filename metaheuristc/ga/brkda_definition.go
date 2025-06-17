package ga

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

const nameBRKGA = "BRKGA"

type ConfigurationBRKGA struct {
	TimeLimitSeconds           float64
	PopulationSize             int
	EliteRatio                 float64
	MutantRation               float64
	CrossoverAlpha             float64
	MutationAlpha              float64
	MaxGenerations             int
	MaxGenerationNoImprovement int
}

type BRKGA struct {
	env           definition.Environment
	configuration *ConfigurationBRKGA
	logger        *logger.Log
	search        search.Local
	RG            *random.Generator
	solutionPool  *solution.Pool
}
