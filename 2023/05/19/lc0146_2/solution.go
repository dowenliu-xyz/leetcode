package lc0146_2

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	cap   int
	table map[int]*Node
	dummy *Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.Pre, dummy.Next = dummy, dummy
	return LRUCache{
		cap:   capacity,
		table: make(map[int]*Node),
		dummy: dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.table[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.pushHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.table[key]; ok {
		node.Val = value
		this.remove(node)
		this.pushHead(node)
		return
	}
	if len(this.table) == this.cap {
		tail := this.dummy.Pre
		this.remove(tail)
		delete(this.table, tail.Key)
		tail.Pre, tail.Next = nil, nil // help GC
	}
	node := &Node{Key: key, Val: value}
	this.pushHead(node)
	this.table[key] = node
}

func (this *LRUCache) remove(node *Node) {
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
}

func (this *LRUCache) pushHead(node *Node) {
	this.dummy.Next, this.dummy.Next.Pre, node.Pre, node.Next = node, node, this.dummy, this.dummy.Next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
