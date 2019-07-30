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
	n := 1
	fmt.Println(n, fib(n))
	n = 2
	fmt.Println(n, fib(n))
	n = 3
	fmt.Println(n, fib(n))
	n = 4
	fmt.Println(n, fib(n))
	n = 5
	fmt.Println(n, fib(n))
	n = 6
	fmt.Println(n, fib(n))
}

//斐波那契数列
func fib(n int) int {
	//已知条件 第一个是1
	l1 := 1
	if n == 1 {
		return l1
	}

	//已知条件 第二个是2
	l2 := 2
	if n == 2 {
		return l2
	}

	var r int
	for i := 3; i <= n; i++ {
		//每个数都等于前两个数相加
		r = l1 + l2
		//循环完一次后， 更新前两个数的值
		l1 = l2
		l2 = r
	}
	return r
}
