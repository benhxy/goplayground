package goleetcode

import "container/list"

func preorderRecursive(root *NaryNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)

	for i := 0; i < len(root.Children); i++ {
		result = append(result, preorderRecursive(root.Children[i])...)
	}

	return result
}

// Learning: Using list as stack.
func preorderIterative(root *NaryNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		current := stack.Remove(stack.Back()).(*NaryNode)
		result = append(result, current.Val)
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack.PushBack(current.Children[i])
		}
	}

	return result
}
