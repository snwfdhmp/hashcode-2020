package main

import "fmt"

type Plan struct {
	SortedLibraries []Library
}

func (p *Plan) String() string {
	out := ""
	for i := range p.SortedLibraries {
		out += fmt.Sprintf("%d: %#v\n", i, p.SortedLibraries[i])
	}
	return out
}
