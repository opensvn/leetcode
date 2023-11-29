/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	path := make([]int, len(graph))
	var traverse func(graph [][]int, n, index int)
	traverse = func(graph [][]int, n, index int) {
		path[index] = n
		if n == len(graph)-1 {
			pathCopy := make([]int, index+1)
			copy(pathCopy, path)
			res = append(res, pathCopy)
		}
		for _, neighbour := range graph[n] {
			traverse(graph, neighbour, index+1)
		}
	}
	traverse(graph, 0, 0)
	return res
}

// @lc code=end
