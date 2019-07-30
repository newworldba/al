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
	//i是起始位置
	i := 0
	//n是台阶总数
	n := 1
	fmt.Println(n, tree(i, n))
	n = 2
	fmt.Println(n, tree(i, n))
	n = 3
	fmt.Println(n, tree(i, n))
	/*
	  以前以n=4为例，介绍下树的结构
	               0  //0分出1和2
	          1         2  // 2分出3和4
	       2     3   3     4  // 4==总台阶数，终止条件
	     3   4  4 5 4 5
	    4 5
	*/
	n = 4
	fmt.Println(i, n, tree(i, n))
	n = 5
	fmt.Println(i, n, tree(i, n))
	n = 6
	fmt.Println(i, n, tree(i, n))
}

//记录表
var h = make(map[int]int)

//树方法
func tree(i, n int) int {
	//终止条件 当前的台阶数不能大于总台阶数
	if i > n {
		return 0
	}
	//终止条件 当前的台阶等于总台阶，则算是走完了一次
	if i == n {
		return 1
	}
	//在记录表中检查下，当前的方式是否尝试过
	if v, ok := h[i]; ok {
		return v
	}
	//树方法的主体思路， 一个节点可以分拆出两种可能
	return tree(i+1, n) + tree(i+2, n)
}
