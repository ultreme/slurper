package slurper

import (
	"math/rand"
	"strings"
)

// Slurper slurps
type Slurper interface {
	Slurp(string) string
}

// Slurp implements the slurper interface
type Slurp struct {
}

// Slurp slurps
func (s *Slurp) Slurp(in string) string {
	var res string

	words := strings.Split(in, " ")
	for idx := range(words) {
		if rand.Int() % 2 == 1 {
			res += "slurp "
		}
		res += words[idx]
		res += " "
	}

	return res
}

// NewSlurper returns a brand new Slurper
func NewSlurper() Slurper {
	return &Slurp{}
}
