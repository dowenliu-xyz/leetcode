package lc0146_1

type LinkedNode struct {
	Key  int
	Val  int
	Pre  *LinkedNode
	Next *LinkedNode
}

type LRUCache struct {
	cap   int
	size  int
	table map[int]*LinkedNode
	dummy *LinkedNode
}

func Constructor(capacity int) LRUCache {
	dummy := &LinkedNode{}
	dummy.Pre, dummy.Next = dummy, dummy
	return LRUCache{
		cap:   capacity,
		size:  0,
		table: make(map[int]*LinkedNode),
		dummy: dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.table[key]
	if !ok {
		return -1
	}
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
	node.Pre, node.Next, this.dummy.Next, this.dummy.Next.Pre = this.dummy, this.dummy.Next, node, node
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.table[key]
	if !ok {
		node = &LinkedNode{Key: key, Val: value}
		if this.cap == this.size {
			removing := this.dummy.Pre
			this.dummy.Pre, removing.Pre.Next = removing.Pre, this.dummy
			removing.Next, removing.Pre = nil, nil
			delete(this.table, removing.Key)
			this.size--
		}
		node.Pre, node.Next, this.dummy.Next, this.dummy.Next.Pre = this.dummy, this.dummy.Next, node, node
		this.table[key] = node
		this.size++
		return
	}
	node.Val = value
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
	node.Pre, node.Next, this.dummy.Next, this.dummy.Next.Pre = this.dummy, this.dummy.Next, node, node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
