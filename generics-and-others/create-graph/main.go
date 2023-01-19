package main

import "fmt"

type Graph struct {
	list map[int][]int
}

func (c *Graph) addEdge(from int, to int) {
	c.list[from] = append(c.list[from], to)
	c.list[to] = append(c.list[to], from)
}

func (c *Graph) showGraph() {
	for vertex, edge := range c.list {
		fmt.Println(vertex, " -> ", edge)
	}
}

func main() {

	myNewGraph := &Graph{
		list: make(map[int][]int),
	}
	myNewGraph.addEdge(1, 2)
	myNewGraph.addEdge(2, 3)
	myNewGraph.addEdge(2, 4)
	myNewGraph.addEdge(2, 5)
	myNewGraph.addEdge(5, 1)

	myNewGraph.showGraph()
	// 1  ->  [2 5]
	// 2  ->  [1 3 4 5]
	// 3  ->  [2]
	// 4  ->  [2]
	// 5  ->  [2 1]
}
