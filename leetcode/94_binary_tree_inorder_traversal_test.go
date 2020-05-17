package goleetcode

import "testing"

func TestInorderTraversal(t *testing.T) {
	root := TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	expected := []int{1, 3, 2}
	result := inorderTraversal(&root)
	if !Equal(result, expected) {
		t.Errorf("Expected %v, actual %v", expected, result)
	}
}

func TestInorderTraversalWithNilNode(t *testing.T) {
	expected := make([]int, 0)
	result := inorderTraversal(nil)
	if !Equal(result, expected) {
		t.Errorf("Expected %v, actual %v", expected, result)
	}
}
