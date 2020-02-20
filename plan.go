package main

import (
	"fmt"
	"io"
	"strings"
)

type Plan struct {
	SortedLibraries []Library
}

func (p *Plan) String() string {
	out := ""
	for i := range p.SortedLibraries {
		if len(p.SortedLibraries[i].Books) == 0 {
			continue
		}
		out += fmt.Sprintf("%d: %#v\n", i, p.SortedLibraries[i])
	}
	return out
}

func (p *Plan) Write(w io.Writer) error {
	_, err := io.WriteString(w, fmt.Sprintf("%d\n", len(p.SortedLibraries)))
	if err != nil {
		return err
	}

	for i := range p.SortedLibraries {
		_, err = io.WriteString(w, fmt.Sprintf("%d %d\n", p.SortedLibraries[i].ID, len(p.SortedLibraries[i].Books)))
		if err != nil {
			return err
		}
		str := make([]string, 0)
		for j := range p.SortedLibraries[i].Books {
			str = append(str, fmt.Sprintf("%d", p.SortedLibraries[i].Books[j].ID))
		}
		_, err = io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(str, " ")))
		if err != nil {
			return err
		}
	}
	return nil
}
