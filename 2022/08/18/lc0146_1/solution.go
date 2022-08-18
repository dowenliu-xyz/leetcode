package main

type LinkedNode struct {
	Key   int
	Value int
	Prev  *LinkedNode
	Next  *LinkedNode
}

type LRUCache struct {
	cap  int
	size int
	list *LinkedNode
	mp   map[int]*LinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:  capacity,
		size: 0,
		list: &LinkedNode{},
		mp:   make(map[int]*LinkedNode, capacity),
	}
	cache.list.Prev, cache.list.Next = cache.list, cache.list
	return cache
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
	if this.size == this.cap {
		node := this.list.Prev
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		delete(this.mp, node.Key)
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
