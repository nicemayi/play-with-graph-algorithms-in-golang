package chapter02

import (
	"fmt"
	"graph/common"
	"strings"
)

type AdjMatrix struct {
	v   int
	e   int
	adj [][]int
}

func (g *AdjMatrix) String() string {
	res := []string{fmt.Sprintf("V = %d, E = %d", g.v, g.e)}
	for _, row := range g.adj {
		res = append(res, fmt.Sprintf("%v", row))
	}
	return strings.Join(res, "\n")
}

func NewAdjMatrix(filename string) *AdjMatrix {
	gInfo := common.GetUnweightedGraphFromFile(filename)

	adjMatrix := &AdjMatrix{}
	adjMatrix.v = gInfo.V
	adjMatrix.e = gInfo.E
	adjMatrix.adj = make([][]int, adjMatrix.v)
	for i := range adjMatrix.adj {
		adjMatrix.adj[i] = make([]int, adjMatrix.v)
	}

	for _, row := range gInfo.Edges {
		i, j := row[0], row[1]
		adjMatrix.adj[i][j] = 1
	}

	return adjMatrix
}
