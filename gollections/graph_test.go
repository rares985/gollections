package gollections

import (
	"math"
	"reflect"
	"testing"
)

func TestAddEdgeUndirected(t *testing.T) {
	testCases := map[string]struct {
		g    *Graph
		want bool
		u    int
		v    int
	}{
		"empty graph": {
			g:    &Graph{},
			want: false,
			u:    1,
			v:    2,
		},
		"has edge": {
			g: NewGraph(false, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
				2: {3: 1, 1: 1},
				3: {1: 1, 2: 1},
			})),
			want: true,
			u:    1,
			v:    2,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.g.IsEdge(tc.u, tc.v)

			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestAddEdgeDirected(t *testing.T) {
	testCases := map[string]struct {
		g    *Graph
		want bool
		u    int
		v    int
	}{
		"empty graph": {
			g:    &Graph{},
			want: false,
			u:    1,
			v:    2,
		},
		"has edge ok": {
			g: NewGraph(true, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
			})),
			want: true,
			u:    1,
			v:    2,
		},
		"has edge inverted": {
			g: NewGraph(true, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
			})),
			want: false,
			u:    2,
			v:    1,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.g.IsEdge(tc.u, tc.v)

			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHasNode(t *testing.T) {
	testCases := map[string]struct {
		g    *Graph
		want bool
		node int
	}{
		"empty graph": {
			g:    &Graph{},
			want: false,
			node: 1,
		},
		"has node ok": {
			g: NewGraph(true, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
			})),
			want: true,
			node: 1,
		},
		"has node oob": {
			g: NewGraph(true, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
			})),
			want: false,
			node: 4,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.g.HasNode(tc.node)

			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestUndirectedFloydWarshall(t *testing.T) {
	testCases := map[string]struct {
		g    *Graph
		want [][]float64
	}{
		"empty graph": {
			g:    &Graph{},
			want: [][]float64{},
		},
		"fw ok": {
			g: NewGraph(false, WithAdjacencyList(AdjacencyList{
				1: {2: 1, 3: 1},
			})),
			want: [][]float64{
				{0, math.Inf(1), math.Inf(1), math.Inf(1)},
				{math.Inf(1), 0, 1, 1},
				{math.Inf(1), 1, 0, 2},
				{math.Inf(1), 1, 2, 0},
			},
		},
		"fw ok complex": {
			g: NewGraph(false, WithAdjacencyList(AdjacencyList{
				1: {2: 2, 3: 7, 4: 8, 5: 1},
				2: {5: 3, 4: 4},
				3: {1: 7, 4: 6},
				4: {2: 4, 3: 6, 1: 8, 5: 1},
				5: {1: 1, 4: 1, 2: 3},
			})),
			want: [][]float64{
				{0, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1)},
				{math.Inf(1), 0, 2, 7, 2, 1},
				{math.Inf(1), 2, 0, 9, 4, 3},
				{math.Inf(1), 7, 9, 0, 6, 7},
				{math.Inf(1), 2, 4, 6, 0, 1},
				{math.Inf(1), 1, 3, 7, 1, 0},
			},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.g.FloydWarshall()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(\n%+v) != Want(\n%+v)", got, tc.want)
			}
		})
	}
}
