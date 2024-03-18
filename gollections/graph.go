package gollections

import (
	"fmt"
	"math"
)

type Graph struct {
	directed bool
	Adj      [][]float64
}

type GraphOption func(this *Graph)
type AdjacencyList map[int]map[int]float64

func NewGraph(directed bool, opts ...GraphOption) *Graph {
	g := &Graph{directed: directed}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func WithAdjacencyList(list AdjacencyList) GraphOption {
	return func(this *Graph) {

		// Some adjacency lists might omit empty edge lists - count all nodes
		nodes := make(map[int]bool)
		for vertex, edges := range list {
			nodes[vertex] = true
			for neighbour := range edges {
				nodes[neighbour] = true
			}
		}
		n := len(nodes)

		this.Adj = make([][]float64, n+1)
		for i := range this.Adj {
			this.Adj[i] = make([]float64, n+1)
		}

		for i := 0; i < n+1; i++ {
			for j := 0; j < n+1; j++ {
				this.Adj[i][j] = math.Inf(1)
			}
			this.Adj[i][i] = 0.0
		}

		for vertex, edges := range list {
			for neighbour, weight := range edges {
				this.Adj[vertex][neighbour] = weight
				if !this.directed {
					this.Adj[neighbour][vertex] = weight
				}
			}
		}
	}
}

func (g *Graph) AddEdge(u, v int, w float64) {
	if u < 0 || u >= len(g.Adj) || v > 0 && v >= len(g.Adj) {
		return
	}

	g.Adj[u][v] = w
	if !g.directed {
		g.Adj[v][u] = w
	}
}

func (g *Graph) IsEdge(u, v int) bool {
	n := len(g.Adj)
	return u >= 0 && u < n && v >= 0 && v < n && g.Adj[u][v] != math.Inf(1)
}

func (g *Graph) HasNode(v int) bool {
	return v >= 0 && v < len(g.Adj)
}

// Check all possible paths between 2 nodes
func (g *Graph) FloydWarshall() [][]float64 {
	n := len(g.Adj)
	if n == 0 {
		return [][]float64{}
	}

	distance := make([][]float64, n)
	for i := range distance {
		distance[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			distance[i][j] = g.Adj[i][j]
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// Is it faster to go from i to j by detouring via k?
				detourDistance := distance[i][k] + distance[k][j]
				if detourDistance < distance[i][j] {
					distance[i][j] = detourDistance
				}
			}
		}
	}

	return distance
}

func (g *Graph) ConnectedComponents() [][]int {
	n := len(g.Adj)

	if n == 0 {
		return [][]int{}
	}

	components := [][]int{}
	visited := make([]bool, n)

	for i := 1; i < n; i++ {
		if visited[i] {
			continue
		}
		s := Stack[int]{[]int{i}}
		component := []int{}

		for !s.IsEmpty() {
			v := s.Pop()
			if !visited[*v] {
				visited[*v] = true
				component = append(component, *v)
				for j := 0; j < n; j++ {
					if *v != j && g.Adj[*v][j] != math.Inf(1) {
						fmt.Printf("%v ", j)
						s.Push(j)
					}
				}
			}
		}
		components = append(components, component)
	}

	return components
}

func (g *Graph) String() string {
	result := ""

	for i := range g.Adj {
		result += fmt.Sprintf("%d: %v\n", i, g.Adj[i])
	}
	return result
}
