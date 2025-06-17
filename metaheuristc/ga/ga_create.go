package ga

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/constants"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

func CreateDefaultGA(env definition.Environment, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *GA {
	configuration := &ConfigurationGA{
		TimeLimitSeconds:           constants.DefaultTimeLimitSeconds,
		PopulationSize:             constants.DefaultPopulationSize,
		CrossoverAlpha:             constants.DefaultCrossoverAlpha,
		MutationAlpha:              constants.DefaultMutationAlpha,
		MaxGenerations:             constants.DefaultMaxGenerations,
		MaxGenerationNoImprovement: constants.DefaultMaxGenerationNoImprovement,
	}

	local := search.CreateMirrorLocalSearch(env)

	return &GA{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(nameGA),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}

func CreateGA(env definition.Environment, configuration *ConfigurationGA, searchType search.Type, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *GA {
	local := search.Create(searchType, env, rg)
	return &GA{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(nameGA),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}
