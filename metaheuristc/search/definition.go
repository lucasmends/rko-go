package search

import (
	"rko/metaheuristc"
	"rko/random"
)

type Type = int

const (
	Swap Type = iota
	Mirror
	Farey
	RVND
)

type Local interface {
	SetRG(rg *random.Generator)
	Search(rko *metaheuristc.RandomKeyValue)
}
