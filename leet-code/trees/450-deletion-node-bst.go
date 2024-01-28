package trees

func DeleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return node
	}

	if key < node.Val { // key is in left subtree
		node.Left = DeleteNode(node.Left, key)
	} else if key > node.Val { // key is in right subtree
		node.Right = DeleteNode(node.Right, key)
	} else { // current node is the one to be deleted
		if node.Left == nil && node.Right == nil { // Case 1, No children
			node = nil
		} else if node.Left == nil { // Case 2, Only one child node, right
			node = node.Right
		} else if node.Right == nil { // Case 2, Only one child node, left
			node = node.Left
		} else { // Case 3, Both children
			t := findMin(node.Right)
			// we don't want to replace the node entirely with minNode,
			// "node" L & R pointers could get replaced
			node.Val = t.Val
			node.Right = DeleteNode(node.Right, t.Val)
		}
	}

	return node
}

// findMin will find the minimum node from a given
// (right subtree) node by recursively calling self
// Idea: The elem for an subtree will be found
// by continously traversing the left subtrees (given
// the BST property)
func findMin(node *TreeNode) *TreeNode {
	if node.Left != nil {
		return findMin(node.Left)
	}
	return node
}
