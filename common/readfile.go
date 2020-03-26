package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func convertStringSliceAoIntSlice(ss []string) []int {
	res := make([]int, len(ss))
	for i, s := range ss {
		val, _ := strconv.Atoi(s)
		res[i] = val
	}
	return res
}

func GetUnweightedGraphFromFile(filename string) *UnweightedGraphInfo {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := make([]string, 0)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	g := &UnweightedGraphInfo{}
	g.Edges = make([][]int, 0)
	for i, content := range res {
		c := convertStringSliceAoIntSlice(strings.Split(content, " "))
		if i == 0 {
			// Through out this course
			// the first line should always have two values:
			// number of verticals & number of edges
			g.V, g.E = c[0], c[1]
			continue
		}
		g.Edges = append(g.Edges, c)
	}

	return g
}
