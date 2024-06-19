package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		// Move the accessed element to the front of the list
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		// Update the value if the key exists
		elem.Value.(*entry).value = value
		this.list.MoveToFront(elem)
	} else {
		// Insert the new entry
		newElem := this.list.PushFront(&entry{key, value})
		this.cache[key] = newElem

		// Check if capacity exceeded, remove the least recently used entry if needed
		if this.list.Len() > this.capacity {
			oldest := this.list.Back()
			delete(this.cache, oldest.Value.(*entry).key)
			this.list.Remove(oldest)
		}
	}
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1

	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1

	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
