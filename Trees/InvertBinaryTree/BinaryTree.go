package main

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T any] struct {
	root *TreeNode[T]
}

func NewBinaryTree[T any](ListOfNodes []*TreeNode[T]) *BinaryTree[T] {
	tree := &BinaryTree[T]{}
	tree.root = tree.createBinaryTree(ListOfNodes)
	return tree
}

func (tree *BinaryTree[T]) createBinaryTree(ListOfNodes []*TreeNode[T]) *TreeNode[T] {
	if len(ListOfNodes) == 0 {
		return nil
	}
	//root node of binary tree
	root := &TreeNode[T]{Data: ListOfNodes[0].Data}

	queue := []*TreeNode[T]{root}

	i := 1
	for i < len(ListOfNodes) {
		curr := queue[0]
		queue = queue[1:]

		// If the node is not nil, create a new TreeNode object for its Left child,
		// set it as the Left child of the current node, and add it to the queue
		if ListOfNodes[i] != nil {
			curr.Left = &TreeNode[T]{Data: ListOfNodes[i].Data}
			queue = append(queue, curr.Left)
		}

		i++

		// If there are more ListOfNodes in the list and the next node is not nil,
		// create a new TreeNode object for its Right child, set it as the Right child
		// of the current node, and add it to the queue
		if i < len(ListOfNodes) && ListOfNodes[i] != nil {
			curr.Right = &TreeNode[T]{Data: ListOfNodes[i].Data}
			queue = append(queue, curr.Right)
		}

		i++
	}

	return root
}
