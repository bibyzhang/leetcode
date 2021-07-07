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

	l.LinkList.head.next = tail
	l.LinkList.tail.prev = head

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
	//if(l.LinkList.head.next==nil) {
	//	l.LinkList.head.next = newNode
	//	l.LinkList.tail.prev = newNode
	//	newNode.prev = l.LinkList.head
	//	newNode.next = l.LinkList.tail
	//} else {
		//current := l.LinkList.head
		////寻找最后一个节点
		//for current.next != nil {
		//	current = current.next
		//}
		current := l.LinkList.tail.prev
		newNode.next = current.next
		newNode.prev = current
		current.next = newNode
		l.LinkList.tail.prev = newNode
	//}

	l.LinkList.size++
	//fmt.Printf("%v", l.LinkList)

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

	//if(l.LinkList.size==1) {
	//	l.LinkList.head = &LinkNode{}
	//	l.LinkList.tail = &LinkNode{}
	//} else {
		//非最后一个元素,最后一个元素next为nil,将prev的next置为nil即可
		//if(node.next!=nil) {
			node.prev.next = node.next
			node.next.prev = node.prev
		//} else {
		//	node.prev.next = nil
		//}
		//l.LinkList.tail.prev = node
	//}

	l.hashmap[key] = nil
	l.LinkList.size--
}

//删除最旧的节点
//最旧的节点即是头部节点
func (l *LRUCache) RemoveLastLink() {
	//fmt.Printf("%v", l.LinkList.head.next.key)
	node := l.LinkList.head.next
	key := node.key
	node.next.prev = node.prev
	node.prev.next = node.next
	//l.LinkList.head = node
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
	obj := Constructor(1)
	obj.Put(2,1)
	param_1 := obj.Get(2)
	obj.Put(3,2)
	param_2 := obj.Get(2)
	param_3 := obj.Get(3)

	//测试4
	//["LRUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	//obj := Constructor(2)
	//param_1 := obj.Get(2)
	//obj.Put(2,6)
	//param_2 := obj.Get(1)
	//obj.Put(1,5)
	//obj.Put(1,2)
	//obj.Get(2)
	//param_3 := obj.Get(1)
	//param_4 := obj.Get(2)

//["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
//	[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]

	//
	fmt.Printf("%v", param_1)
	fmt.Printf("%v", param_2)
	fmt.Printf("%v", param_3)
	//fmt.Printf("%v", param_4)
	//fmt.Printf("%v", param_5)
}