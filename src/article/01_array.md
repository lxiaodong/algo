# [array 数据结构详细介绍](https://time.geekbang.org/column/article/40961?utm_source=pinpaizhuanqu&utm_medium=geektime&utm_campaign=guanwang&utm_term=guanwang&utm_content=0511)<i class="fa fa-external-link"></i>
>&#128227; [【数组算法题库】](./../database/01_array.md)

## 1. 数组是什么?：
数组是一种线性表数据结构，是用一组连续的内存空间来存储一组具有相同数据类型的数据。

### 1.1 什么是线性表？
线性表就是数据排成像一条线一样的结构。每个线性表只有前后两个方向。例如数组、链表、队列、栈

>注意：线性表的反面就是非线性表，与之相反的是，数据不单单是可以排成一条线一样的结构，而且数组直之间不仅是前后两个方向。例如二叉树、堆

### 1.2 连续的内存空间&相同的数据结构
连续的内存空间和相同的数据结构提供一个“杀手锏”：随机访问，也就是根据数组的下标进行访问。（`随机访问，可以通过数组值的类型计算空间地址, 具体可以参考coed/go下的intro.go`）
> 弊端：虽然可以随机访问，但是删除，插入等操作破坏了连续性，需要大量的数据搬移操作。

## 2. 时间复杂度
### 2.1 随机访问时间复杂度
因为数组支持随机访问，所以根据下标可以计算出空间地址，随意时间复杂度是O(1)

### 2.2 插入操作时间复杂度
- 保证数组的顺序，那么插入数据就需要移动插入位置之后的所有数据，所以最好的时间复杂度是O(1), 最坏时间复杂度是O(N), 平均时间复杂度O(N)
- 不保证数组的顺序，那么插入数据的时候可以与当前插入位置的元素进行位置互换`k, n-1 = n-1, k`, 那么时间复杂度是O(1)

### 2.3 删除操作时间复杂度
- 保证数组的顺序，那么删除数据就需要移动删除元素的位置之后的所有数据，所以最好的时间复杂度是O(1), 最坏时间复杂度是O(N), 平均时间复杂度O(N)
- 如果不保证连续性，可以批量删除，这样就会减少搬移的数据


