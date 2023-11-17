/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type Node struct {
	prev *Node
	next *Node
	key int
	value int
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (d *DoublyLinkedList) AddLast(node *Node) {
	node.prev = d.tail.prev
	node.next = d.tail
	d.tail.prev.next = node
	d.tail.prev = node
	d.size++
}

func (d *DoublyLinkedList) Remove(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	d.size--
}

func (d *DoublyLinkedList) RemoveFirst() *Node {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	d.Remove(first)
	return first
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

type LRUCache struct {
	keys map[int]*Node
	cache *DoublyLinkedList
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := &DoublyLinkedList{
		head: &Node{},
		tail: &Node{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return LRUCache{
		keys: make(map[int]*Node, capacity),
		cache: cache,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.keys[key]
	if !ok {
		return -1
	}
	this.makeRecently(key)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.keys[key]; ok {
		this.remove(key)
		this.addRecently(key, value)
		return
	}

	if this.cache.Size() == this.capacity {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

func (this *LRUCache) makeRecently(key int) {
	if node, ok := this.keys[key]; ok {
		this.cache.Remove(node)
		this.cache.AddLast(node)
	}
}

func (this *LRUCache) addRecently(key, value int) {
	node := &Node{key: key, value: value}
	this.cache.AddLast(node)
	this.keys[key] = node
}

func (this *LRUCache) remove(key int) {
	if node, ok := this.keys[key]; ok {
		this.cache.Remove(node)
		delete(this.keys, key)
	}
}

func (this *LRUCache) removeLeastRecently() {
	node := this.cache.RemoveFirst()
	if node != nil {
		delete(this.keys, node.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
