package multistart

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/constants"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

func CreateDefaultMultiStart(env definition.Environment, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *MultiStart {
	configuration := &Configuration{
		MaxIterations:    constants.DefaultMaxIterations,
		TimeLimitSeconds: constants.DefaultTimeLimitSeconds,
	}

	local := search.CreateDefault(env, rg)

	return &MultiStart{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}

func CreateMultiStart(env definition.Environment, configuration *Configuration, searchType search.Type, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *MultiStart {
	local := search.Create(searchType, env, rg)
	return &MultiStart{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}
