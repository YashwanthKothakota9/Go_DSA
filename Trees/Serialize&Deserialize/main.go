package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var MARKER = "M"
var m = 1

/*
	serializeRec is a recursive function that is used

by Serialize function as a helper function
*/
func serializeRec(node *TreeNode[int], stream *[]string) {
	// Adding marker to stream if the node is nil
	if node == nil {
		(*stream) = append((*stream), MARKER+strconv.Itoa(m))
		m += 1
		return
	}

	// Adding node to stream
	(*stream) = append((*stream), strconv.Itoa(node.Data))

	// Doing a pre-order tree traversal for serialization
	serializeRec(node.Left, stream)
	serializeRec(node.Right, stream)
}

func Serialize(root *TreeNode[int]) []string {
	stream := make([]string, 0)
	serializeRec(root, &stream)
	return stream
}

// deserialize is used to deserialize integer array into a binary tree.
func deserialize(stream *[]string) *TreeNode[int] {
	// dequeue first element form array
	val := (*stream)[0]
	*stream = (*stream)[1:len(*stream)] // reframing the slice takes O(1) time complexity

	// Return nil when a marker is encountered
	if string(val[0]) == MARKER {
		return nil
	}

	// Creating new Binary Tree Node from current value from stream
	tempVal, _ := strconv.Atoi(val)
	node := &TreeNode[int]{Data: tempVal}

	// Doing a pre-order tree traversal for deserialization
	node.Left, node.Right = deserialize(stream), deserialize(stream)

	// Return node if it exists
	return node
}

// Driver code
func main() {
	listOfTrees := [][]*TreeNode[int]{
		{&TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 2}, &TreeNode[int]{Data: 3}, &TreeNode[int]{Data: 4}, &TreeNode[int]{Data: 5}},
		{&TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 350}},
		{&TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 350}},
		{&TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 350}, &TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 25}},
		{&TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 350}},
		{&TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 350}},
		{&TreeNode[int]{Data: 350}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 100}},
		{&TreeNode[int]{Data: 1}, nil, &TreeNode[int]{Data: 2}, nil, &TreeNode[int]{Data: 3}, nil, &TreeNode[int]{Data: 4}, nil, &TreeNode[int]{Data: 5}},
		{},
	}

	// Create the binary trees using the BinaryTree class
	inputTrees := make([]*BinaryTree[int], len(listOfTrees))
	for i, listOfNodes := range listOfTrees {
		inputTrees[i] = NewBinaryTree[int](listOfNodes)
	}

	for i, tree := range inputTrees {

		fmt.Printf("%d.\tBinary tree:\n", i+1)
		displayTree(tree.root)

		fmt.Printf("\n\tMarker used for NULL nodes in serialization/deserialization: %s*\n", MARKER)

		// Serialization
		ostream := serialize(tree.root)
		fmt.Print("\n\tSerialized integer array:\n")
		fmt.Print("\t", strings.Replace(fmt.Sprint(ostream), " ", ", ", -1))

		// Deserialization
		deserializedRoot := deserialize(&ostream)
		fmt.Print("\n\n\tDeserialized binary tree:\n")
		displayTree(deserializedRoot)

		fmt.Printf("%s\n", strings.Repeat("-", 100))
		m = 1
	}
}

func height[T any](node *TreeNode[T]) int {
	if node == nil {
		return 0
	} else {
		return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
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
