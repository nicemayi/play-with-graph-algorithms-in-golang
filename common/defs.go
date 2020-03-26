package common

import (
	"fmt"
)

type UnweightedGraphInfo struct {
	V     int
	E     int
	Edges [][]int
}

func (g *UnweightedGraphInfo) String() string {
	return fmt.Sprintf(
		"Unweighted graph, with:\n\t%d verticals\n\t%d edges\nAll edges: \n\t%v",
		g.V, g.E, g.Edges,
	)
}
