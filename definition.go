package rko

import (
	"fmt"
	"rko/definition"
	"rko/logger"
	"rko/metaheuristc"
	"rko/metaheuristc/solution"
	"rko/random"
	"sync"
)

type Solver struct {
	l            *logger.Log
	rg           *random.Generator
	env          definition.Environment
	solutionPool *solution.Pool
	solvers      []definition.Solver
}

func (s *Solver) Solve() any {
	var wg sync.WaitGroup

	for i, sv := range s.solvers {
		fmt.Printf("Running solver %s (%d)\n", sv.Name(), i+1)
		wg.Add(1)
		go metaheuristc.Worker(sv, &metaheuristc.Configuration{Id: i + 1}, &wg)
	}

	wg.Wait()

	rk := s.solutionPool.BestSolution()
	return s.env.Decode(rk.RK)
}
