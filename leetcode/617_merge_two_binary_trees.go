package goleetcode

// Learning: To return a struct pointer, return the address of the struct (&variable)
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	newTree := TreeNode{}
	newTree.Val = t1.Val + t2.Val
	newTree.Left = mergeTrees(t1.Left, t2.Left)
	newTree.Right = mergeTrees(t1.Right, t2.Right)

	return &newTree
}
