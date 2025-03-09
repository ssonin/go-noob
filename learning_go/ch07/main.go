package main

import (
	"io"
	"os"
	"slices"
	"strings"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamA string, scoreA int, teamB string, scoreB int) {
	if _, ok := l.Teams[teamA]; !ok {
		return
	}
	if _, ok := l.Teams[teamB]; !ok {
		return
	}
	if scoreA > scoreB {
		l.Wins[teamA]++
	} else if scoreB > scoreA {
		l.Wins[teamB]++
	}
}

func (l League) Ranking() []string {
	result := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		result = append(result, k)
	}
	slices.SortFunc(result, func(a, b string) int {
		return l.Wins[b] - l.Wins[a]
	})
	return result
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	io.WriteString(w, strings.Join(r.Ranking(), "\n"))
}

func main() {
	teamA := Team{"WASP", []string{"Joey"}}
	teamB := Team{"DUFUS", []string{"Ross"}}

	l := League{
		Teams: map[string]Team{
			"WASP":  {"WASP", []string{"Joey"}},
			"DUFUS": {"DUFUS", []string{"Ross"}},
		},
		Wins: map[string]int{}}
	l.MatchResult(teamA.Name, 1, teamB.Name, 0)
	RankPrinter(l, os.Stdout)
}
