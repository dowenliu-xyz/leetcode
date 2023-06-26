package lc0146

type cacheNode struct {
	Key  int
	Val  int
	Next *cacheNode
	Pre  *cacheNode
}

type LRUCache struct {
	capacity int
	mp       map[int]*cacheNode
	dummy    *cacheNode
}

func Constructor(capacity int) LRUCache {
	dummy := &cacheNode{}
	dummy.Next, dummy.Pre = dummy, dummy
	return LRUCache{
		capacity: capacity,
		mp:       make(map[int]*cacheNode),
		dummy:    dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.mp[key]
	if !ok {
		return -1
	}
	this.pop(node)
	this.pushHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.mp[key]
	if ok {
		this.pop(node)
		node.Val = value
	} else {
		if len(this.mp) == this.capacity {
			this.pop(this.dummy.Pre)
		}
		node = &cacheNode{
			Key: key,
			Val: value,
		}
	}
	this.pushHead(node)
}

func (this *LRUCache) pop(node *cacheNode) {
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
	delete(this.mp, node.Key)
}

func (this *LRUCache) pushHead(node *cacheNode) {
	this.dummy.Next, this.dummy.Next.Pre, node.Pre, node.Next = node, node, this.dummy, this.dummy.Next
	this.mp[node.Key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
