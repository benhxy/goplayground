package goleetcode

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// process left subtree, return root and rightmost leaf
	// process right subtree, return root and rightmost leaf
	// link root in between lefttree's rightmost leaf
	// set root.Left to null to de-link
	// return left tree root and right tree rightmost leaf
	// terminal condition: when node has no children, return itself, itself
	newRoot, _ := helper(root)
	return newRoot
}

// Learning: Pointer operation (implicit dereferencing) and multiple return values.
func helper(root *TreeNode) (min *TreeNode, max *TreeNode) {
	min = root
	max = root

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		leftMin, leftMax := helper(root.Left)
		min = leftMin
		leftMax.Right = root
		max = root
	}

	if root.Right != nil {
		rightMin, rightMax := helper(root.Right)
		max.Right = rightMin
		max = rightMax
	}

	root.Left = nil

	return
}
