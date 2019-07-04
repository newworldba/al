package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	grid = [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]int) (num int) {
	// 遍历所有节点
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			//只要有1就肯定存在岛屿,将关联岛屿置空，下次遍历将不会有该岛屿的连接点
			if grid[i][j] == 1 {
				search(grid, i, j)
				num++
			}
		}
	}
	return num
}

//将关联的岛屿全部置空
func search(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == 0 {
		return
	}
	//中心点置空
	grid[i][j] = 0
	//上下左右的陆地全部置空
	search(grid, i-1, j)
	search(grid, i+1, j)
	search(grid, i, j-1)
	search(grid, i, j+1)
}
