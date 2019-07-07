/**
  你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

示例 1:

输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:

输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:

输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:

输入: deadends = ["0000"], target = "8888"
输出：-1

*/
package main

import "fmt"
import "container/list"

func main() {
	var deadends []string
	var target string
	deadends = []string{"0000"}
	target = "1111"
	deadends = []string{"0201", "0101", "0102", "1212", "2002"}
	target = "0202"
	fmt.Println(deadends, target, openLock(deadends, target))
}

func openLock(deadends []string, target string) int {
	deads := make(map[string]bool)
	//死亡密码数组改为映射格式，便于查找
	for _, d := range deadends {
		deads[d] = true
	}
	start, step := "0000", 0
	queue := list.New()
	tmpA := byte('0' - 1)
	tmpB := byte('9' + 1)
	//bfs 需要注意的是不要有死循环，使用一个映射记录已经遍历过的节点
	used := make(map[string]bool)
	queue.PushBack(start)
	used[start] = true

	for queue.Len() > 0 {
		//在遍历时，queue的长度是有变化的，所以需要在外部计算queue的长度,而不是在for内部
		l := queue.Len()
		for i := 0; i < l; i++ {
			//从queue中取出节点，进行计算
			v, _ := queue.Remove(queue.Front()).(string)
			if deads[v] {
				continue
			}
			if v == target {
				return step
			}
			tmpP := []byte(v)
			tmpM := []byte(v)
			for j := 0; j < 4; j++ {
				tmpP[j]++
				tmpM[j]--
				//转动的结果如果超出枚举值，需要初始化
				if tmpP[j] == tmpB {
					tmpP[j] = '0'
				}
				if tmpM[j] == tmpA {
					tmpM[j] = '9'
				}
				tmpPS := string(tmpP)
				if _, ok := used[tmpPS]; !ok {
					queue.PushBack(tmpPS)
					used[tmpPS] = true
				}
				tmpMS := string(tmpM)
				if _, ok := used[tmpMS]; !ok {
					queue.PushBack(tmpMS)
					used[tmpMS] = true
				}
				tmpM[j] = []byte(v)[j]
				tmpP[j] = []byte(v)[j]
			}
		}
		step++
	}
	return -1
}
