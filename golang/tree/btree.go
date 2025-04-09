package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BTreeNode[K constraints.Ordered] struct {
	Value K
	Left  *BTreeNode[K]
	Right *BTreeNode[K]
}

func (n *BTreeNode[K]) TraversePreorder(process func(a K)) {
	process(n.Value)
	if n.Left != nil { n.Left.TraversePreorder(process) }
	if n.Right != nil { n.Right.TraversePreorder(process) }
}

func (n *BTreeNode[K]) TraverseInorder(process func(a K)) {
	if n.Left != nil { n.Left.TraverseInorder(process) }
	process(n.Value)
	if n.Right != nil { n.Right.TraverseInorder(process) }
}

func (n *BTreeNode[K]) TraversePostorder(process func(a K)) {
	if n.Left != nil { n.Left.TraversePostorder(process) }
	if n.Right != nil { n.Right.TraversePostorder(process) }
	process(n.Value)
}

func (n *BTreeNode[K]) TraverseDepthFirst(process func (a K)) {
	queue := make([]*BTreeNode[K], 0)
	queue = append(queue, n)

	i := 0
	for i < len(queue) {
		node := queue[i]
		i++

		process(node.Value)

		if node.Left != nil { queue = append(queue, node.Left) }
		if node.Right != nil { queue = append(queue, node.Right) }
	}
}

func (n *BTreeNode[K]) Add(value K) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &BTreeNode[K] { Value: value }
		} else {
			n.Left.Add(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &BTreeNode[K]{ Value: value }
		} else {
			n.Right.Add(value)
		}
	}
}

func (n *BTreeNode[K]) Find(target K) *BTreeNode[K] {
	switch {
	case n.Value == target:
		return n
	case n.Value > target:
		if n.Left == nil { return nil }
		return n.Left.Find(target)
	case n.Value < target:
		if n.Right == nil { return nil }
		return n.Right.Find(target)
	default:
		return nil
	}
}

func TreeShowcase() {
	nodeA := BTreeNode[string] { Value: "A" }
	nodeC := BTreeNode[string] { Value: "C" }
	nodeB := BTreeNode[string] { Value: "B", Left: &nodeA, Right: &nodeC }
	nodeE := BTreeNode[string] { Value: "E" }
	nodeD := BTreeNode[string] { Value: "D", Left: &nodeB, Right: &nodeE }

	print := func(a string) () {
		fmt.Println(a);
	}

	counter := 0
	stepCounterPrint := func (a string) {
		counter++
		fmt.Printf("step: %d, value: %s\n", counter, a)
	}

	fmt.Println("Preorder traverse")
	nodeD.TraversePreorder(print)

	fmt.Println("\nPreorder traverse subtree")
	nodeB.TraversePreorder(stepCounterPrint)
	
	fmt.Println("\nInorder traverse")
	nodeD.TraverseInorder(print)
	
	fmt.Println("\nPostorder traverse")
	nodeD.TraversePostorder(print)
	
	fmt.Println("\nDepth first traverse")
	nodeD.TraverseDepthFirst(print)

	printInt := func (value int) {
		fmt.Println(value)
	}
	root := BTreeNode[int] { Value: 5 }
	root.Add(2)
	root.Add(1)
	root.Add(4)
	root.Add(3)
	root.Add(6)
	root.Add(9)
	root.Add(10)
	root.Add(7)
	root.Add(8)
	fmt.Println("\nInorder traverse")
	root.TraverseInorder(printInt)

	fmt.Println("Searching")
	fmt.Println(root.Find(7))
	fmt.Println(root.Find(12))
}