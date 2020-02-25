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
	length := 0
	for i := range p.SortedLibraries {
		if len(p.SortedLibraries[i].Books) == 0 {
			break
		}
		length += 1
	}
	_, err := io.WriteString(w, fmt.Sprintf("%d\n", length))
	if err != nil {
		return err
	}

	score := 0
	for i := 0; i < length; i++ {
		_, err = io.WriteString(w, fmt.Sprintf("%d %d\n", p.SortedLibraries[i].ID, len(p.SortedLibraries[i].Books)))
		if err != nil {
			return err
		}
		str := make([]string, 0)
		for j := range p.SortedLibraries[i].Books {
			str = append(str, fmt.Sprintf("%d", p.SortedLibraries[i].Books[j].ID))
			score += p.SortedLibraries[i].Books[j].Score
		}
		_, err = io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(str, " ")))
		if err != nil {
			return err
		}
	}
	// fmt.Printf("Score after writting: %d\n", score)
	return nil
}
