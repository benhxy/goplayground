package goleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NaryNode struct {
	Val      int
	Children []*NaryNode
}
