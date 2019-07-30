/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/
package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i, rec(i))
	i = 2
	fmt.Println(i, rec(i))
	i = 3
	fmt.Println(i, rec(i))
	i = 4
	fmt.Println(i, rec(i))
	i = 5
	fmt.Println(i, rec(i))
	i = 6
	fmt.Println(i, rec(i))
}

var h = make(map[int]int)

//递归方法
func rec(i int) int {
	//如果当前的台阶数<1, 则该方式无效，并终止条件
	if i < 1 {
		return 0
	}
	//已知条件，如果只有1阶台阶，只有一种方法
	if i == 1 {
		return 1
	}
	//已知条件，如果只有2阶台阶，只有两种方法
	if i == 2 {
		return 2
	}
	//记录当前台阶的方法，如果不记录，时间复杂度是O(2^n)；增加记录后，时间复杂度是O(n)
	if v, ok := h[i]; ok {
		return v
	}
	//递归的主体思路：当前的台阶数等于前一个台阶的方式+前两个台阶的方式
	return rec(i-1) + rec(i-2)
}
