package goleetcode

// Learning: Use ... to append items to slice.
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
