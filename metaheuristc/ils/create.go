package ils

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/constants"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

func CreateDefaultILS(env definition.Environment, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *ILS {
	configuration := &Configuration{
		MaxIterations:       constants.DefaultMaxIterations,
		TimeLimitSeconds:    constants.DefaultTimeLimitSeconds,
		ShakeMin:            constants.DefaultShakeMin,
		ShakeMax:            constants.DefaultShakeMax,
		MetropolisCriterion: false,
	}

	local := search.CreateDefault(env, rg)

	return &ILS{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}

func CreateILS(env definition.Environment, configuration *Configuration, searchType search.Type, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *ILS {
	local := search.Create(searchType, env, rg)
	return &ILS{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(name),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}
