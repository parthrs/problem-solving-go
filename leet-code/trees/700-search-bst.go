package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	// edge case
	if root == nil || root.Val == val {
		return root
	}

	// Recursive case(s)
	// We don't need to check if the left
	// or right subtrees are nil here, they
	// will be the first thing caught (in the edge
	// case section) on the recursive call
	if val < root.Val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}
