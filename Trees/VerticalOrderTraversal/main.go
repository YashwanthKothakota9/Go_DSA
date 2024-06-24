package main

import (
	"fmt"
	"strings"
)

func VerticalOrder(root *TreeNode[int]) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	nodeList := make(map[int][]int)
	minColumn := 0
	maxIndex := 0

	queue := new(Queue)
	queue.Enqueue(root, 0)
	for !queue.IsEmpty() {
		node, column := queue.Dequeue()
		if node != nil {
			temp := nodeList[column]
			temp = append(temp, node.Data)
			nodeList[column] = temp

			minColumn = min(minColumn, column)
			maxIndex = max(maxIndex, column)

			queue.Enqueue(node.Left, column-1)
			queue.Enqueue(node.Right, column+1)
		}
	}

	result := make([][]int, 0)
	for x := minColumn; x < maxIndex+1; x++ {
		result = append(result, nodeList[x])
	}
	return result
}
func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// min function returns the minimum of the integers provided
func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
func arrayToString(arr [][]int) string {
	var builder strings.Builder

	builder.WriteString("[")

	for i, row := range arr {
		if i > 0 {
			builder.WriteString(", ")
		}

		builder.WriteString("[")
		for j, num := range row {
			if j > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%d", num))
		}
		builder.WriteString("]")
	}

	builder.WriteString("]")

	return builder.String()
}
