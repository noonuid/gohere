package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	var testCases = []struct {
		vertexes []int
		edges    [][]int
		distWant map[int]int
		pathWant map[int]int
	}{
		{
			vertexes: []int{},
			edges:    [][]int{},
			distWant: nil,
			pathWant: nil,
		},
		{
			vertexes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			edges: [][]int{
				{0, 1, 4},
				{0, 7, 8},
				{1, 2, 8},
				{1, 7, 11},
				{2, 3, 7},
				{2, 5, 4},
				{2, 8, 2},
				{3, 4, 9},
				{3, 5, 14},
				{4, 5, 10},
				{5, 6, 2},
				{6, 7, 1},
				{6, 8, 6},
				{7, 8, 7},
			},
			distWant: map[int]int{
				0: 0,
				1: 4,
				2: 12,
				3: 19,
				4: 21,
				5: 11,
				6: 9,
				7: 8,
				8: 14,
			},
			pathWant: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 5,
				5: 6,
				6: 7,
				7: 0,
				8: 2,
			},
		},
	}

	for caseIndex, testCase := range testCases {
		var graph = NewGraph(false)
		for _, edge := range testCase.edges {
			graph.AddWeightedEdge(edge[0], edge[1], edge[2])
		}
		distGot, pathGot := graph.Dijkstra(0)

		if !reflect.DeepEqual(distGot, testCase.distWant) || !reflect.DeepEqual(pathGot, testCase.pathWant) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ndistGot: %v, distWant: %v\npathGot: %v, pathWant: %v",
				caseIndex, testCase, distGot, testCase.distWant, pathGot, testCase.pathWant)
		}
	}
}

func TestFloyd(t *testing.T) {
	var testCases = []struct {
		vertexes []int
		edges    [][]int
		distWant map[int]map[int]int
		pathWant map[int]map[int]int
	}{
		{
			vertexes: []int{},
			edges:    [][]int{},
			distWant: nil,
			pathWant: nil,
		},
		{
			vertexes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			edges: [][]int{
				{0, 1, 4},
				{0, 7, 8},
				{1, 2, 8},
				{1, 7, 11},
				{2, 3, 7},
				{2, 5, 4},
				{2, 8, 2},
				{3, 4, 9},
				{3, 5, 14},
				{4, 5, 10},
				{5, 6, 2},
				{6, 7, 1},
				{6, 8, 6},
				{7, 8, 7},
			},
			distWant: map[int]map[int]int{
				0: {0: 0, 1: 4, 2: 12, 3: 19, 4: 21, 5: 11, 6: 9, 7: 8, 8: 14},

				1: {0: 4, 1: 0, 2: 8, 3: 15, 4: 22, 5: 12, 6: 12, 7: 11, 8: 10},

				2: {0: 12, 1: 8, 2: 0, 3: 7, 4: 14, 5: 4, 6: 6, 7: 7, 8: 2},

				3: {0: 19, 1: 15, 2: 7, 3: 0, 4: 9, 5: 11, 6: 13, 7: 14, 8: 9},

				4: {0: 21, 1: 22, 2: 14, 3: 9, 4: 0, 5: 10, 6: 12, 7: 13, 8: 16},

				5: {0: 11, 1: 12, 2: 4, 3: 11, 4: 10, 5: 0, 6: 2, 7: 3, 8: 6},

				6: {0: 9, 1: 12, 2: 6, 3: 13, 4: 12, 5: 2, 6: 0, 7: 1, 8: 6},

				7: {0: 8, 1: 11, 2: 7, 3: 14, 4: 13, 5: 3, 6: 1, 7: 0, 8: 7},

				8: {0: 14, 1: 10, 2: 2, 3: 9, 4: 16, 5: 6, 6: 6, 7: 7, 8: 0},
			},
			pathWant: map[int]map[int]int{
				0: {0: -1, 1: -1, 2: 1, 3: 2, 4: 6, 5: 6, 6: 7, 7: -1, 8: 2},

				1: {0: -1, 1: -1, 2: -1, 3: 2, 4: 5, 5: 2, 6: 7, 7: -1, 8: 2},

				2: {0: 1, 1: -1, 2: -1, 3: -1, 4: 5, 5: -1, 6: 5, 7: 6, 8: -1},

				3: {0: 2, 1: 2, 2: -1, 3: -1, 4: -1, 5: 2, 6: 5, 7: 6, 8: 2},

				4: {0: 6, 1: 5, 2: 5, 3: -1, 4: -1, 5: -1, 6: 5, 7: 6, 8: 5},

				5: {0: 6, 1: 2, 2: -1, 3: 2, 4: -1, 5: -1, 6: -1, 7: 6, 8: 2},

				6: {0: 7, 1: 7, 2: 5, 3: 5, 4: 5, 5: -1, 6: -1, 7: -1, 8: -1},

				7: {0: -1, 1: -1, 2: 6, 3: 6, 4: 6, 5: 6, 6: -1, 7: -1, 8: -1},

				8: {0: 2, 1: 2, 2: -1, 3: 2, 4: 5, 5: 2, 6: -1, 7: -1, 8: -1},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		var graph = NewGraph(false)
		for _, edge := range testCase.edges {
			graph.AddWeightedEdge(edge[0], edge[1], edge[2])
		}
		distGot, pathGot := graph.Floyd()

		if !reflect.DeepEqual(distGot, testCase.distWant) || !reflect.DeepEqual(pathGot, testCase.pathWant) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ndistGot: %v, distWant: %v\npathGot: %v, pathWant: %v",
				caseIndex, testCase, distGot, testCase.distWant, pathGot, testCase.pathWant)
		}
	}
}
