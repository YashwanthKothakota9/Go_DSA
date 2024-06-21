package main

import (
	"fmt"
	"math"
	"strings"
)

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func diameterHelper(node *TreeNode[int], diameter int) (int, int) {
	if node == nil {
		return 0, diameter
	} else {
		//compute height of each sub tree
		lh, diameter := diameterHelper(node.Left, diameter)
		rh, diameter := diameterHelper(node.Right, diameter)

		diameter = max(diameter, (lh + rh))

		return max(lh, rh) + 1, diameter
	}
}

func diameterOfBinaryTree(root *TreeNode[int]) int {
	if root == nil {
		return 0
	}
	diameter := 0
	_, diameter = diameterHelper(root, diameter)
	return diameter
}

func height[T any](node *TreeNode[T]) int {
	if node == nil {
		return 0
	} else {
		return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
	}
}

// Driver code
func main() {
	listOfTrees := [][]*TreeNode[int]{
		{&TreeNode[int]{Data: 2}, &TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 4}, &TreeNode[int]{Data: 3}, &TreeNode[int]{Data: 5}, &TreeNode[int]{Data: 6}, &TreeNode[int]{Data: 7}},
		{&TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 2}, &TreeNode[int]{Data: 3}, &TreeNode[int]{Data: 4}, &TreeNode[int]{Data: 5}, &TreeNode[int]{Data: 6}, &TreeNode[int]{Data: 7}, &TreeNode[int]{Data: 8}, &TreeNode[int]{Data: 9}},
		{&TreeNode[int]{Data: 45}, &TreeNode[int]{Data: 32}, &TreeNode[int]{Data: 23}, &TreeNode[int]{Data: 21}, &TreeNode[int]{Data: 19}, &TreeNode[int]{Data: 19}, &TreeNode[int]{Data: 18}, &TreeNode[int]{Data: 1}},
		{&TreeNode[int]{Data: 5}, &TreeNode[int]{Data: 3}, &TreeNode[int]{Data: 4}, &TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 2}, &TreeNode[int]{Data: 6}, &TreeNode[int]{Data: 7}, &TreeNode[int]{Data: 8}, &TreeNode[int]{Data: 9}},
		{&TreeNode[int]{Data: 9}, &TreeNode[int]{Data: 7}, nil, nil, &TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 8}, &TreeNode[int]{Data: 10}, nil, &TreeNode[int]{Data: 12}},
	}

	// Create the binary trees using the BinaryTree class
	inputTrees := make([]*BinaryTree[int], len(listOfTrees))
	for i, listOfNodes := range listOfTrees {
		inputTrees[i] = NewBinaryTree[int](listOfNodes)
	}

	// Print the input trees
	for i, tree := range inputTrees {
		fmt.Printf("%d.\tOriginal tree:\n", i+1)
		// DisplayTree() is a custom function used for prinitng purposes
		displayTree(tree.root)
		fmt.Printf("\n\tDiameter of tree: %d\n", diameterOfBinaryTree(tree.root))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}

}
func displayNode[T any](output []string, linkAbove []string, node *TreeNode[T], level int, p int, linkChar string) {
	if node == nil {
		return
	}
	h := len(output)
	SP := " "
	if p < 0 {
		extra := strings.Repeat(" ", int(math.Abs(float64(p))))
		for i := 0; i < len(output); i++ {
			if len(output[i]) > 0 {
				output[i] = extra + output[i]
			}
		}
		for i := 0; i < len(linkAbove); i++ {
			if len(linkAbove[i]) > 0 {
				linkAbove[i] = extra + linkAbove[i]
			}
		}
	}
	if level < h-1 {
		p = int(math.Max(float64(p), float64(len(output[level+1]))))
	}
	if level > 0 {
		p = int(math.Max(float64(p), float64(len(output[level-1]))))
	}
	p = int(math.Max(float64(p), float64(len(output[level]))))

	if node.Left != nil {
		LeftData := SP + fmt.Sprintf("%v", node.Left.Data) + SP
		displayNode(output, linkAbove, node.Left, level+1, p-len(LeftData), "L")
		p = int(math.Max(float64(p), float64(len(output[level+1]))))
	}

	space := p - len(output[level])
	if space > 0 {
		output[level] += strings.Repeat(" ", space)
	}
	nodeData := SP + fmt.Sprintf("%v", node.Data) + SP
	output[level] += nodeData

	space = p + len(SP) - len(linkAbove[level])
	if space > 0 {
		linkAbove[level] += strings.Repeat(" ", space)
	}
	linkAbove[level] += linkChar

	if node.Right != nil {
		displayNode(output, linkAbove, node.Right, level+1, len(output[level]), "R")
	}
}

func displayTree[T any](rootNode *TreeNode[T]) {
	if rootNode == nil {
		fmt.Println("\tnil")
	}
	h := height(rootNode)
	output := make([]string, h)
	linkAbove := make([]string, h)
	for i := 0; i < h; i++ {
		output[i] = ""
		linkAbove[i] = ""
	}
	displayNode(output, linkAbove, rootNode, 0, 5, " ")

	for i := 1; i < h; i++ {
		for j := 0; j < len(linkAbove[i]); j++ {
			if linkAbove[i][j] != ' ' {
				size := len(output[i-1])
				if size < j+1 {
					output[i-1] += strings.Repeat(" ", j+1-size)
				}
				jj := j
				if linkAbove[i][j] == 'L' {
					for ; output[i-1][jj] == ' '; jj++ {
					}
					for k := j + 1; k < jj; k++ {
						output[i-1] = output[i-1][:k] + "_" + output[i-1][k+1:]
					}
				} else if linkAbove[i][j] == 'R' {
					for output[i-1][jj] == ' ' {
						jj--
					}
					for k := j - 1; k > jj; k-- {
						output[i-1] = output[i-1][:k] + "_" + output[i-1][k+1:]
					}
				}
				linkAbove[i] = linkAbove[i][:j] + "|" + linkAbove[i][j+1:]
			}
		}
	}

	for i := 0; i < h; i++ {
		if i != 0 {
			fmt.Println("\t" + linkAbove[i])
		}
		fmt.Println("\t" + output[i])
	}
}
