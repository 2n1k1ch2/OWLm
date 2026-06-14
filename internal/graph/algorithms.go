package graph

func DFS(graph *Graph, start *Node) []*Node {
	result := make([]*Node, 0)
	visited := make(map[string]struct{})
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if _, ok := visited[node.Name]; ok {
			return
		}
		visited[node.Name] = struct{}{}
		result = append(result, node)
		for _, dep := range node.Dependencies {
			if next, ok := graph.GetNode(dep.Target.Name); ok {
				dfs(next)
			}
		}
	}
	dfs(start)
	return result
}

func BFS(graph *Graph, start *Node) map[*Node]uint32 {

	visited := make(map[*Node]uint32)
	queue := make([]*Node, 0)
	visited[start] = 0
	for len(queue) > 0 {
		current := queue[0]
		currentCycle := visited[current]
		queue = queue[1:]
		for _, dep := range current.Dependencies {
			node, ok := graph.GetNode(dep.Target.Name)
			if _, marked := visited[node]; !marked && ok {
				visited[node] = currentCycle + 1
			}
		}
	}
	return visited
}
