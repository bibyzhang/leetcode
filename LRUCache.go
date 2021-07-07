//https://leetcode-cn.com/problems/lru-cache/

package main

import "fmt"

//LRU HashLink
type LRUCache struct {
	hashmap map[int]*LinkNode
	LinkList LinkList
}

//链表
type LinkList struct {
	head *LinkNode//指向第一个节点,通过遍历节点next到所有节点
	tail *LinkNode//指向最后一个节点,通过遍历prev到所有节点
	size int
	capacity int
}

//链表节点
type LinkNode struct {
	key int//对应map中的key
	value int
	prev *LinkNode
	next *LinkNode
}

func Constructor(capacity int) LRUCache {
	var LRUCache LRUCache
	LRUCache.InitLink(capacity)
	return LRUCache
}

func (this *LRUCache) Get(key int) int {
	if(this.hashmap[key]==nil) {
		return -1
	}
	//放到最新
	this.MakeNodeRecenty(key)
	return this.hashmap[key].value
}

func (this *LRUCache) Put(key int, value int)  {
	if(this.hashmap[key]!=nil) {
		this.RemoveLink(key)
	}
	this.AddLink(key,value)
}

//初始化链表
func (l *LRUCache) InitLink(capacity int) {
	l.hashmap = make(map[int]*LinkNode)

	head := &LinkNode{}

	tail := &LinkNode{}

	l.LinkList.head = head
	l.LinkList.tail = tail

	l.LinkList.size = 0
	l.LinkList.capacity = capacity
}

//添加链表节点
//添加到链表尾部
//尾部节点即是最新的节点
func (l *LRUCache) AddLink(key int, value int) {
	//存储已满,删除旧的一个节点
	if(l.LinkList.size>=l.LinkList.capacity) {
		l.RemoveLastLink()
	}

	newNode := &LinkNode{}
	newNode.key = key
	newNode.value = value
	//第一个节点
	if(l.LinkList.head.next==nil) {
		l.LinkList.head.next = newNode
		l.LinkList.tail.prev = newNode
		newNode.prev = l.LinkList.head
	} else {
		current := l.LinkList.head
		//寻找最后一个节点
		for current.next != nil {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
		current.next.prev = current
	}

	l.LinkList.size++
	l.hashmap[key] = newNode
}

//将节点放到最新位置
func (l *LRUCache) MakeNodeRecenty(key int) {
	node := l.hashmap[key]
	value := node.value
	l.RemoveLink(key)
	l.AddLink(key, value)
}

//删除指定key
func (l *LRUCache) RemoveLink(key int) {
	node := l.hashmap[key]

	if(l.LinkList.size==1) {
		l.LinkList.head = &LinkNode{}
		l.LinkList.tail = &LinkNode{}
	} else {
		//非最后一个元素,最后一个元素next为nil,将prev的next置为nil即可
		if(node.next!=nil) {
			node.prev.next = node.next
			node.next.prev = node.prev
		} else {
			node.prev.next = nil
		}
	}

	l.hashmap[key] = nil
	l.LinkList.size--
}

//删除最旧的节点
//最旧的节点即是头部节点
func (l *LRUCache) RemoveLastLink() {
	node := l.LinkList.head.next
	key := node.key
	l.LinkList.head = l.LinkList.head.next
	l.hashmap[key] = nil
	l.LinkList.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	//capacity := 10
	//key := 1
	//vaule := 1

	//初始化链表
	//var LRUCache LRUCache
	//LRUCache.InitLink(capacity)
	//LRUCache.AddLink(1,1)
	//LRUCache.AddLink(2,2)
	//LRUCache.AddLink(3,3)
	//LRUCache.AddLink(4,4)
	//fmt.Println("v%",LRUCache.LinkList.size)

	//测试1
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//obj := Constructor(2)
	//obj.Put(1,1)
	//obj.Put(2,2)
	//param_1 := obj.Get(1)
	//obj.Put(3,3)
	//param_2 := obj.Get(2)
	//obj.Put(4,4)
	//param_3 := obj.Get(1)
	//param_4 := obj.Get(3)
	//param_5 := obj.Get(4)

	//obj.Get(1)
	//obj.Put(3,3)
	//obj.Get(2)
	//obj.Put(4,4)
	//obj.Get(1)
	//obj.Get(3)
	//obj.Get(4)

//[null,null,null,1,null,-1,null,-1,3,4]

	//["LRUCache","get"]
	//	[[1],[0]]

	//obj := Constructor(capacity)
	//obj.Put(2)

	//obj.Put(1,1)
	//obj.Put(2,2)
	//obj.Put(3,3)
	//obj.Put(4,4)
	//obj.Put(5,5)
	//obj.Put(6,6)
	//obj.Put(7,7)
	//obj.Put(8,8)
	//obj.Put(9,9)
	//obj.Put(10,10)
	//obj.Put(11,11)
	////param_1 := obj.Get(1)
	//param_1 := obj.Get(3)

	//测试3
	//["LRUCache","put","get","put","get","get"]
	//[[1],[2,1],[2],[3,2],[2],[3]]
	//obj := Constructor(1)
	//obj.Put(2,1)
	//param_1 := obj.Get(2)
	//obj.Put(3,2)
	//param_2 := obj.Get(2)
	//param_3 := obj.Get(3)

	//测试4
	//["LRUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	obj := Constructor(2)
	param_1 := obj.Get(2)
	obj.Put(2,6)
	param_2 := obj.Get(1)
	obj.Put(1,5)
	obj.Put(1,2)
	obj.Get(2)
	param_3 := obj.Get(1)
	param_4 := obj.Get(2)

	//
	fmt.Printf("%v", param_1)
	fmt.Printf("%v", param_2)
	fmt.Printf("%v", param_3)
	fmt.Printf("%v", param_4)
	//fmt.Printf("%v", param_5)
}