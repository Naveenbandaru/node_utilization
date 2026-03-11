import (
	"fmt"
)

type Node struct {
	id          int
	utilization int
}

func createNodes() []Node {
	return []Node{
		{1, 52},
		{2, 48},
		{3, 44},
		{4, 41},
		{5, 39},
	}
}

func printUtilization(nodes []Node) {
	for _, n := range nodes {
		fmt.Println("Node", n.id, "Utilization", n.utilization)
	}
}

func calculateAverage(nodes []Node) int {
	sum := 0
	for _, n := range nodes {
		sum += n.utilization
	}
	return sum / len(nodes)
}

func main() {
	nodes := createNodes()
	printUtilization(nodes)
	avg := calculateAverage(nodes)
	fmt.Println("Average Utilization", avg)
}
