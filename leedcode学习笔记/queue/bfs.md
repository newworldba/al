## BFS
广度优先搜索; 常见问题为在一个局域网节点中找出两点之间的最短路径


## 解题方式
1. 将根节点点放入队列中
2. 将根节点的关联节点（姑且称之为根2节点）放入队列中
3. 移除根节点
4. 将队列中的节点的关联节点（根3节点）放入队列中
5. 移除根2节点
6. 重复4，5步骤，直到找到目标节点，层数即为最短路径

## 注意点
1. 为保证不会出现死循环，需要将已放入队列中的节点写入哈希表中，遍历节点时，如果发现已经遍历过，就不应再放入队列
2. 放入新层的节点时，先将上一层的节点移除，这样才能不重复引入节点


