package main

import (
	"container/list"
	"fmt"
	"sync"
)

// LRUCache 结构体包含哈希表和双向链表
type LRUCache struct {
	capacity int
	// 存储键和值
	cache map[int]*list.Element
	// 双向链表，按访问顺序排序
	list *list.List
	// 锁，保证线程安全
	lock sync.RWMutex
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get 获取缓存中的值
func (c *LRUCache) Get(key int) (int, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.cache == nil {
		return 0, false
	}

	elem, exists := c.cache[key]
	if !exists {
		return 0, false
	}

	// 将访问的元素移动到双向链表的末尾，表示最近使用
	c.list.MoveToBack(elem)

	return elem.Value.(int), true
}

// Put 添加或更新缓存中的值
func (c *LRUCache) Put(key int, value int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.cache == nil {
		c.cache = make(map[int]*list.Element)
		c.list = list.New()
	}

	if c.capacity <= 0 {
		return
	}

	// 如果键已存在，更新值并移动到链表末尾
	if elem, exists := c.cache[key]; exists {
		c.list.MoveToBack(elem)
		elem.Value = value
		return
	}

	// 如果缓存已满，移除链表头部的元素（最不常用的元素）
	if c.list.Len() == c.capacity {
		oldest := c.list.Front()
		c.list.Remove(oldest)
		delete(c.cache, oldest.Value.(int))
	}

	// 添加新元素到缓存和链表
	back := c.list.PushBack(value)
	c.cache[key] = back
}

func main() {
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 输出 1

	lru.Put(3, 3)           // 淘汰 key 2
	fmt.Println(lru.Get(2)) // 输出 false

	lru.Put(4, 4)           // 淘汰 key 1
	fmt.Println(lru.Get(1)) // 输出 false
	fmt.Println(lru.Get(3)) // 输出 3
	fmt.Println(lru.Get(4)) // 输出 4
}
