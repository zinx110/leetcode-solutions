package LRUCache

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*node
	left     *node
	right    *node
}

func Constructor(capacity int) LRUCache {

	lruCache := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*node),
	}

	lruCache.left = &node{0, 0, nil, nil}
	lruCache.right = &node{0, 0, nil, nil}
	lruCache.left.next = lruCache.right
	lruCache.right.prev = lruCache.left

	return lruCache
}

func (this *LRUCache) remove(n *node) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
	delete(this.cache, n.key)
}

func (this *LRUCache) insert(n *node) {
	prev := this.right.prev
	prev.next = n
	n.prev = prev
	n.next = this.right
	this.right.prev = n
	this.cache[n.key] = n
}

func (this *LRUCache) Get(key int) int {
	n, exists := this.cache[key]

	if !exists {
		return -1
	}

	this.remove(n)
	this.insert(n)
	return n.value

}

func (this *LRUCache) Put(key int, value int) {
	n, exists := this.cache[key]
	if exists {
		this.remove(n)
		this.insert(n)
		n.value = value
		return
	}

	if this.size < this.capacity {
		this.size++
	} else {
		first := this.left.next
		this.remove(first)
	}
	n = &node{key, value, nil, nil}
	this.insert(n)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
