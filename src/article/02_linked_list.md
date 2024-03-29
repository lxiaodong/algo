# [linked list 数据结构详细介绍](https://time.geekbang.org/column/article/41013?utm_source=pinpaizhuanqu&utm_medium=geektime&utm_campaign=guanwang&utm_term=guanwang&utm_content=0511)<i class="fa fa-external-link"></i>

>&#128227; [【链表算法题库】](./../database/02_linked_list.md)

## 1.链表是什么？
链表它由节点（Node）的集合组成，每个节点包含数据元素和指向下一个节点的“指针”。链表中的节点可以在内存中离散分布，它们通过指针相互连接，形成一个逻辑上的序列。
> 注意：数组必须是连续的内存，假如申请“100M”的内存，如果连续内存不满100M，数组就没有办法申请，而链表却可以。

## 2.常见的链表结构
### 2.1 单向链表（Singly Linked List）
每个节点只包含一个指向下一个节点的指针。最后一个节点指向空值（null）表示链表的结束。
![Alt text](./../../static/images/02_linked_list/singly_likned_list.png)
> 时间复杂度：
> 查找：O(N)
> 插入：O(1)，注意这里单指插入操作不包含寻址
> 删除：O(1)，注意这里单指删除操作不包含寻址
> 在某个元素后插入或删除某个元素：O(N)

### 2.2 双向链表（Doubly Linked List）
每个节点包含指向前一个节点和后一个节点的两个指针。双向链表可以从头到尾遍历，也可以从尾到头遍历。
![Alt text](./../../static/images/02_linked_list/doubly_linked_list.png)

> 时间复杂度：
> 查找：O(1), 注意可以支持 O(1) 时间复杂度的情况下找到前驱结点。
> 
> 删除：O(1), 注意这是指定删除给定指针指向的结点，不是删除某个值的结点（O(N)）
> 
> 插入：O(1), 注意这是指定指针指向的结点前插入，不是在某个值的结点前插入（O(N)）

### 2.3 循环链表（Circular Linked List）
最后一个节点的指针指向头节点，形成一个循环。循环链表可以从任何节点开始遍历，一直回到起始节点。
> 循环链表又分单循环和双循环

#### 2.3.1 单循环
![Alt text](./../../static/images/02_linked_list/singly_circular_linked_list.png)
#### 2.3.2 双循环
![Alt text](./../../static/images/02_linked_list/double_circular_linked_list.png
)


## 3. 链表与数组的区别
### 3.1 时间复杂度
| 操作      | 链表 | 数组 |
| --------- | ---- | ---- |
| 插入/删除 | O(1) | O(N) |
| 随机访问  | O(N) | O(1) |

### 3.2 CPU缓存
- 数组是连续的内存空间，借助CPU缓存机制，预读数组中的数据，所以访问率高。

- 链表在内存中并非连续存储，所以对CPU缓存不友好，没有办法预读。

### 3.3 大小
> 数组:
> 1. 数组的大小固定，一经声明就需要占用固定的连续的内存空间，如果声明的数组过大，系统没有足够连续的内存，就会提示“”内存不足“。
> 2. 如果声明的数组过小，就会出现不够用的情况，这时只能再申请一个更大的内存空间，把原数组拷贝进去，非常费时。
  
> 链表: 本身没有大小的限制，天然的支持动态扩容。


## 链表(Linked List)使用场景？
> 链表使用场景lru, fifo, lfu 这几个不好记，那就记书房的书多常用的清理方式。
### LRU 最近最少使用策略

### FIFO先进先出策略
### LFU 最少使用策略 
