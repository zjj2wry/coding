package main

import "fmt"

const loadFactor = 0.75

type HashTable struct {
	items              []Item
	size               int
	colisionResolution func(key int) int
}

type Item struct {
	key   int
	value interface{}
}

func NewHashTabel(size int) *HashTable {
	items := make([]Item, size)
	ht := &HashTable{
		items: items,
		size:  size,
	}
	ht.colisionResolution = ht.linearProbe
	return ht
}

func (h *HashTable) hash(key int) int {
	return key % h.size
}

func (h *HashTable) balanceFactor() float64 {
	var i float64
	for _, item := range h.items {
		if item.value != nil {
			i++
		}
	}
	return i / float64(h.size)
}

// 线性探测，在数据非常接近的时候容易导致聚集，导致 hash 表的效率变低
func (h *HashTable) linearProbe(key int) int {
	i := h.hash(key + 1)
	for h.items[i].value != nil && h.items[i].key != key {
		i = h.hash(i + 1)
		if h.balanceFactor() >= loadFactor {
			return -1
		}
	}
	return i
}

// 二次探测，解决数据集非常接近时的聚集问题
func (h *HashTable) quadraticProbe(key int) int {
	i := 1
	index := h.hash(key + i*i)
	for h.items[index].value != nil && h.items[index].key != key {
		index = h.hash(index + i*i)
		if h.balanceFactor() >= loadFactor {
			return -1
		}
	}
	return index
}

func (h *HashTable) Set(key int, value interface{}) {
	i := h.hash(key)
	item := h.items[i]
	if item.value == nil {
		h.items[i] = Item{
			key:   key,
			value: value,
		}
		return
	}
	if item.key == key {
		return
	}

	colision := h.colisionResolution(key)
	if colision < 0 {
		h.rehash()
		h.Set(key, value)
		return
	}

	h.items[colision] = Item{
		key:   key,
		value: value,
	}
	return
}

func (h *HashTable) rehash() {
	h.size *= 2
	items := h.items
	h.items = make([]Item, h.size)
	for _, item := range items {
		if item.value != nil {
			h.Set(item.key, item.value)
		}
	}
	return
}

func (h *HashTable) Get(key int) (interface{}, bool) {
	i := h.hash(key)
	item := h.items[i]
	if item.value == nil {
		return nil, false
	}
	if item.key == key {
		return item.value, true
	}

	colision := h.colisionResolution(key)
	if h.items[colision].value != nil && h.items[colision].key == key {
		return h.items[colision].value, true
	}
	return nil, false
}

func main() {
	hashmap := NewHashTabel(8)
	for i := 0; i < 16; i++ {
		hashmap.Set(i, i)
	}
	for i := 0; i < 16; i++ {
		item, ok := hashmap.Get(i)
		fmt.Println(item, ok)
	}
	hashmap = NewHashTabel(8)
	hashmap.colisionResolution = hashmap.quadraticProbe
	for i := 0; i < 16; i++ {
		hashmap.Set(i, i)
	}
	for i := 0; i < 16; i++ {
		item, ok := hashmap.Get(i)
		fmt.Println(item, ok)
	}
}
