package main

import (
	"strconv"
	"strings"
)

func levelOrderTraversal(root *TreeNode[int]) string {
	var result []string

	if root == nil {
		return "None"
	}

	currentQueue := new(Queue)
	currentQueue.Enqueue(root)

	for currentQueue.Size() > 0 {
		levelSize := currentQueue.Size()
		var levelNodes []string

		for i := 0; i < levelSize; i++ {
			temp := currentQueue.Dequeue()
			levelNodes = append(levelNodes, strconv.Itoa(temp.Data))

			if temp.Left != nil {
				currentQueue.Enqueue(temp.Left)
			}
			if temp.Right != nil {
				currentQueue.Enqueue(temp.Right)
			}
		}

		result = append(result, strings.Join(levelNodes, ", "))
	}

	return strings.Join(result, " : ")
}
