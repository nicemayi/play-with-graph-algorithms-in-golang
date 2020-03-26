package chapter02

import (
	"fmt"
	"testing"
)

func TestAdjMatrix(t *testing.T) {
	adjMatrix := NewAdjMatrix("g.txt")
	fmt.Println(adjMatrix)
}
