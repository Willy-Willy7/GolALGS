package main

import "fmt"

type Graph struct {
	vertices int
	adjList  [][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		adjList:  make([][]int, v),
	}
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjList[src] = append(g.adjList[src], dest)
	g.adjList[dest] = append(g.adjList[dest], src) // For undirected graph
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.vertices)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")

		for _, neighbor := range g.adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}
