package ga

import (
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc/constants"
	"rko/metaheuristc/search"
	"rko/metaheuristc/solution"
	"rko/random"
)

func CreateDefaultBRKGA(env definition.Environment, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *BRKGA {
	configuration := &ConfigurationBRKGA{
		TimeLimitSeconds:           constants.DefaultTimeLimitSeconds,
		PopulationSize:             constants.DefaultPopulationSize,
		EliteRatio:                 constants.DefaultEliteRatio,
		MutantRation:               constants.DefaultMutantRation,
		CrossoverAlpha:             constants.DefaultCrossoverAlpha,
		MutationAlpha:              constants.DefaultMutationAlpha,
		MaxGenerations:             constants.DefaultMaxGenerations,
		MaxGenerationNoImprovement: constants.DefaultMaxGenerationNoImprovement,
	}

	local := search.CreateMirrorLocalSearch(env)

	return &BRKGA{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(nameGA),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}

func CreateBRKGA(env definition.Environment, configuration *ConfigurationBRKGA, searchType search.Type, rg *random.Generator, solutionPool *solution.Pool, logger *logger.Log) *BRKGA {
	local := search.Create(searchType, env, rg)
	return &BRKGA{
		env:           env,
		configuration: configuration,
		logger:        logger.GetLogger(nameGA),
		search:        local,
		solutionPool:  solutionPool,
		RG:            rg,
	}
}
