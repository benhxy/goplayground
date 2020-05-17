package goleetcode

// Learning: Use range to loop over map.
func destCity(paths [][]string) string {
	const notDestination int = 1
	const maybeDestination int = 0

	visited := make(map[string]int)

	for _, path := range paths {
		visited[path[0]] = notDestination

		_, exist := visited[path[1]]
		if !exist {
			visited[path[1]] = maybeDestination
		}
	}

	var destination string
	for k, v := range visited {
		if v == maybeDestination {
			destination = k
		}
	}

	return destination
}
