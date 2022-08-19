package main

type LinkedNode struct {
	Key   int
	Value int
	Prev  *LinkedNode
	Next  *LinkedNode
}

type LRUCache struct {
	size     int
	capacity int
	list     *LinkedNode
	mp       map[int]*LinkedNode
}

func Constructor(capacity int) LRUCache {
	dummyNode := &LinkedNode{}
	dummyNode.Next, dummyNode.Prev = dummyNode, dummyNode
	return LRUCache{
		size:     0,
		capacity: capacity,
		list:     dummyNode,
		mp:       map[int]*LinkedNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.mp[key]
	if !ok {
		return -1
	}
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	node.Prev, node.Next = this.list, this.list.Next
	this.list.Next.Prev, this.list.Next = node, node
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.mp[key]; ok {
		node.Value = value
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		node.Prev, node.Next = this.list, this.list.Next
		this.list.Next.Prev, this.list.Next = node, node
		return
	}
	if this.capacity == this.size {
		last := this.list.Prev
		last.Prev.Next, last.Next.Prev = last.Next, last.Prev
		delete(this.mp, last.Key)
		this.size--
	}
	node := &LinkedNode{
		Key:   key,
		Value: value,
		Prev:  this.list,
		Next:  this.list.Next,
	}
	this.list.Next.Prev, this.list.Next = node, node
	this.mp[key] = node
	this.size++
}
