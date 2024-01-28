package trees

import "fmt"

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack1 := []int{}
	stack2 := []int{}

	stack1 = getLeafNodes(root1, stack1)
	stack2 = getLeafNodes(root2, stack2)

	fmt.Printf("stack1: %v, stack2:%v", stack1, stack2)

	if len(stack1) != len(stack2) {
		return false
	}

	for i, val := range stack1 {
		if val != stack2[i] {
			return false
		}
	}

	return true
}

func getLeafNodes(root *TreeNode, stack []int) []int {
	if root == nil {
		return stack
	}
	if root.Left == nil && root.Right == nil {
		stack = append(stack, root.Val)
		return stack
	}

	stack = getLeafNodes(root.Left, stack)
	stack = getLeafNodes(root.Right, stack)
	return stack
}
