package go_LRU_algorithm

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedlinkedNode
	head, tail *DLinkedlinkedNode
}

type DLinkedlinkedNode struct {
	key, value int
	prev, next *DLinkedlinkedNode
}

func initDLinkedlinkedNode(key, value int) *DLinkedlinkedNode {
	return &DLinkedlinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedlinkedNode{},
		head:     initDLinkedlinkedNode(0, 0),
		tail:     initDLinkedlinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	linkedNode := this.cache[key]
	this.moveToHead(linkedNode)
	return linkedNode.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		linkedNode := initDLinkedlinkedNode(key, value)
		this.cache[key] = linkedNode
		this.addToHead(linkedNode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		linkedNode := this.cache[key]
		linkedNode.value = value
		this.moveToHead(linkedNode)
	}
}

func (this *LRUCache) addToHead(linkedNode *DLinkedlinkedNode) {
	linkedNode.prev = this.head
	linkedNode.next = this.head.next
	this.head.next.prev = linkedNode
	this.head.next = linkedNode
}

func (this *LRUCache) removelinkedNode(linkedNode *DLinkedlinkedNode) {
	linkedNode.prev.next = linkedNode.next
	linkedNode.next.prev = linkedNode.prev
}

func (this *LRUCache) moveToHead(linkedNode *DLinkedlinkedNode) {
	this.removelinkedNode(linkedNode)
	this.addToHead(linkedNode)
}

func (this *LRUCache) removeTail() *DLinkedlinkedNode {
	linkedNode := this.tail.prev
	this.removelinkedNode(linkedNode)
	return linkedNode
}
