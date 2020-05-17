package goleetcode

import "container/list"

// Learning: Using different forms of for loop.
func postorder(root *NaryNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var stack = list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		var current = stack.Remove(stack.Back()).(*NaryNode)
		result = append(result, current.Val)

		for i := range current.Children {
			stack.PushBack(current.Children[i])
		}
	}

	reverse(&result)
	return result
}

// Learning: To use a struct, we need to dereference the pointer passed in.
// Learning: In Go it's ok to exchange values using tuple assignment. This is because during the evaluation of the right side, the values are copied into a tuple.
func reverse(input *[]int) {
	data := *input
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
