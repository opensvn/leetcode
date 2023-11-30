/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	var traverse func(graph [][]int, n int)
	traverse = func(graph[][]int, n int) {
		if onPath[n] {
			hasCycle = true
		}
		if visited[n] || onPath[n] {
			return
		}

		visited[n] = true
		onPath[n] = true
		for _, neighber := range graph[n] {
			traverse(graph, neighber)
		}
		onPath[n] = false
	}

	graph := buildGraph(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}
	return !hasCycle
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	res := make([][]int, numCourses)
	for _, p := range prerequisites {
		from := p[1]
		to := p[0]
		res[from] = append(res[from], to)
	}
	return res
}
// @lc code=end
