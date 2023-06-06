package lc0146_3

type linkedListNode struct {
	key  int
	val  int
	pre  *linkedListNode
	next *linkedListNode
}

type LRUCache struct {
	capacity int
	table    map[int]*linkedListNode
	dummy    *linkedListNode
}

func Constructor(capacity int) LRUCache {
	dummy := &linkedListNode{}
	dummy.pre, dummy.next = dummy, dummy
	return LRUCache{
		capacity: capacity,
		table:    make(map[int]*linkedListNode),
		dummy:    dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.table[key]
	if !ok {
		return -1
	}
	this.pop(node)
	this.pushHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.table[key]
	if ok {
		node.val = value
		this.pop(node)
	} else {
		node = &linkedListNode{key: key, val: value}
		if len(this.table) == this.capacity {
			this.pop(this.dummy.pre)
		}
	}
	this.pushHead(node)
}

func (this *LRUCache) pop(node *linkedListNode) {
	node.pre.next, node.next.pre = node.next, node.pre
	delete(this.table, node.key)
}

func (this *LRUCache) pushHead(node *linkedListNode) {
	this.dummy.next, this.dummy.next.pre, node.pre, node.next = node, node, this.dummy, this.dummy.next
	this.table[node.key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
