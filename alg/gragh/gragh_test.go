package gragh

import (
	"fmt"
	"testing"
)

/*
	0-1-3
	| | |
	2-4-5
	  | |
	  7-6
*/
func creategrash() *Gragh {
	g := NewGragh(8)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)
	g.AddEdge(4, 7)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	return g
}
func TestBfs(t *testing.T) {
	g := creategrash()

	fmt.Println(g.Bfs(0, 7))
	fmt.Println(g.Bfs(0, 6))
}

func TestDfs(t *testing.T) {
	g := creategrash()
	fmt.Println(g.Dfs(0, 7))
}
