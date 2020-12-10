package ten

import (
	"sort"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

// import (
// 	"sort"
// 	"strconv"

// 	"github.com/joshsticks/AdventOfCode2020/inputhelper"
// )

// func Traverse(curr int, nums []int, numMap map[int]int, ch chan int) {
// 	count := 0
// 	if val, ok := numMap[curr+1]; ok {
// 		if val == len(nums)-1 {
// 			count++
// 		} else {
// 			c := make(chan int)

// 			go func() {
// 				Traverse(curr+1, nums, numMap, c)
// 				close(c)
// 			}()

// 			for i := range c {
// 				count += i
// 			}
// 		}
// 	}

// 	if val, ok := numMap[curr+2]; ok {
// 		if val == len(nums)-1 {
// 			count++
// 		} else {
// 			// Traverse(val, nums, numMap, ch)
// 			c := make(chan int)

// 			go func() {
// 				Traverse(curr+2, nums, numMap, c)
// 				close(c)
// 			}()

// 			for i := range c {
// 				count += i
// 			}
// 		}
// 	}

// 	if val, ok := numMap[curr+3]; ok {
// 		if val == len(nums)-1 {
// 			count++
// 		} else {
// 			// Traverse(val, nums, numMap, ch)
// 			c := make(chan int)

// 			go func() {
// 				Traverse(curr+3, nums, numMap, c)
// 				close(c)
// 			}()

// 			for i := range c {
// 				count += i
// 			}
// 		}
// 	}

// 	ch <- count
// }

// func SolvePart2() int {
// 	count := 0
// 	numberStrings := inputhelper.GetStrings("inputs/input10.txt")
// 	var numbers []int

// 	for _, x := range numberStrings {
// 		num, _ := strconv.Atoi(x)
// 		numbers = append(numbers, num)
// 	}

// 	sort.Ints(numbers)

// 	numMap := make(map[int]int)

// 	for i, x := range numbers {
// 		numMap[x] = i
// 	}

// 	ch := make(chan int)

// 	go func() {
// 		Traverse(0, numbers, numMap, ch)
// 		close(ch)
// 	}()

// 	for i := range ch {
// 		count += i
// 	}

// 	return count
// }

type Node struct {
	value int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
}

func (g *Graph) PathTo(s *Node, t *Node, q map[*Node]int) int {
	if s.value == t.value {
		return 1
	} else {
		if val, ok := q[s]; ok {
			return val
		} else {
			for _, u := range g.edges[*s] {
				q[s] += g.PathTo(u, t, q)
			}
		}
	}
	return q[s]
}
func SolvePart2() int {
	numberStrings := inputhelper.GetStrings("inputs/input10.txt")
	var numbers []int

	for _, x := range numberStrings {
		num, _ := strconv.Atoi(x)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	var g Graph
	q := make(map[int]*Node)
	for _, x := range numbers {
		n := Node{x}
		g.AddNode(&n)
		q[x] = &n
	}

	for key, item := range q {
		if x, ok := q[key+1]; ok {
			g.AddEdge(item, x)
		}
		if x, ok := q[key+2]; ok {
			g.AddEdge(item, x)
		}
		if x, ok := q[key+3]; ok {
			g.AddEdge(item, x)
		}

	}

	a := make(map[*Node]int)
	c := g.PathTo(q[0], q[numbers[len(numbers)-1]], a)

	return c
}
