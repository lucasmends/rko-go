package vns

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/constants"
	"rko/metaheuristc/solution"
	"rko/random"
)

func CreateDefaultVNS(env definition.Environment, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *VNS {
	configuration := &Configuration{
		MaxIterations:    constants.DefaultMaxIterations,
		TimeLimitSeconds: constants.DefaultTimeLimitSeconds,
		Rate:             constants.DefaultRate,
	}

	return &VNS{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		solutionPool:  solutionPool,
		RG:            rg,
	}
}

func CreateVNS(env definition.Environment, configuration *Configuration, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *VNS {
	return &VNS{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		solutionPool:  solutionPool,
		RG:            rg,
	}
}
