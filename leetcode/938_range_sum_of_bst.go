package goleetcode

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	return sum
}
