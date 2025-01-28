package main

import "fmt"

type TreeNode struct {
	data     int
	children []*TreeNode
}

func (tn *TreeNode) AddChild(i *TreeNode) {
	tn.children = append(tn.children, i)
}

func (tn *TreeNode) Display(l int) {
	for i := 0; i < l; i++ {
		fmt.Print("  ")
	}
	fmt.Println(tn.data)

	for _, j := range tn.children {
		j.Display(l + 1)
	}
}

func main() {
	tn := &TreeNode{data: 1}
	child1 := &TreeNode{data: 2}
	child2 := &TreeNode{data: 3}
	tn.AddChild(child1)
	tn.AddChild(child2)

	child1.AddChild(&TreeNode{data: 4})
	child1.AddChild(&TreeNode{data: 5})
	child2.AddChild(&TreeNode{data: 6})
	child2.AddChild(&TreeNode{data: 7})
	tn.Display(1)
}
