package main

import (
	"fmt"
	"log"
)

type Graph struct {
	list map[int][]int
}

func (c *Graph) addEdge(from int, to int) {

	if !ifExist(c, from, to) {
		c.list[from] = append(c.list[from], to)
		c.list[to] = append(c.list[to], from)
	} else {
		log.Println(" <-X- DECLINED [", from, " -> ", to, "] -X-> ")
	}

}

func (c *Graph) deleteEdge(from int) {
	if ifExist(c, from, -1) {
		delete(c.list, from)
		log.Println("Item deleted :", from)
	} else {
		log.Println("Already not exist this item: ", from)
	}
}

func ifExist(c *Graph, from int, to int) bool {

	_, ok := c.list[from]
	if ok {
		if to == -1 {
			return true
		}
		forList := c.list[from]
		forLen := len(c.list[from])
		for i := 0; i < forLen; i++ {
			if forList[i] == to {
				return true
			}
		}
	}
	return false
}

func (c *Graph) showGraph() {
	for vertex, edge := range c.list {
		fmt.Println(vertex, " -> ", edge)
	}
	fmt.Println("-----------------")
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

	myNewGraph.deleteEdge(3)

	myNewGraph.showGraph()
}
