package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	id          int
	utilization int
}

func generateNodes(count int) []Node {
	nodes := make([]Node, 0)
	for i := 1; i <= count; i++ {
		nodes = append(nodes, Node{i, rand.Intn(15) + 75})
	}
	return nodes
}

func rebalance(nodes []Node) []Node {
	for i := range nodes {
		if nodes[i].utilization < 85 {
			nodes[i].utilization += 3
		}
	}
	return nodes
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nodes := generateNodes(5)
	nodes = rebalance(nodes)

	for _, n := range nodes {
		fmt.Println("Node", n.id, "Utilization", n.utilization)
	}
}
