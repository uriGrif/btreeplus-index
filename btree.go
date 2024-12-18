package mybtree

import (
	"errors"
	"sync"
)

type BtreeIndex[K keyAble, T Indexable[K]] struct {
	mu     sync.RWMutex
	Degree int
	Root   node[K, T]
	Unique bool
}

func (t *BtreeIndex[K, T]) Get(key K) *T {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.Root == nil {
		return nil
	}
	return t.Root.get(key)
}

func (t *BtreeIndex[K, T]) unsyncGet(key K) *T {
	if t.Root == nil {
		return nil
	}
	return t.Root.get(key)
}

func (t *BtreeIndex[K, T]) Insert(value *T) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.Root == nil {
		t.Root = &leafNode[K, T]{
			degree:   t.Degree,
			keys:     make([]K, 0, t.Degree),
			values:   make([]*T, 0, t.Degree),
			prevNode: nil,
			nextNode: nil,
		}
	}
	if t.Unique && t.unsyncGet((*value).GetKey()) != nil {
		return errors.New("this index does not allow duplicated keys")
	}
	t.Root.insert(value, &t.Root, make([]node[K, T], 0))
	return nil
}

func (t *BtreeIndex[K, T]) Delete(key K) error {
	// this is a simplified delete: it only merges the nodes if the leaf node where the entry was deleted is empty
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.unsyncGet(key) == nil {
		return errors.New("can't delete key beacuse it doesn't exist")
	}
	t.Root.delete(key)
	return nil
}

func (t *BtreeIndex[K, T]) levelTraverse(f func(node[K, T])) {
	var q queue[node[K, T]] = make([]node[K, T], 0)
	q.push(t.Root)
	recursiveBfs(t.Root, &q)
	for _, n := range q {
		f(n)
	}
}

func recursiveBfs[K keyAble, T Indexable[K]](n node[K, T], q *queue[node[K, T]]) {
	intNode, ok := (n).(*internalNode[K, T])
	if ok {
		for _, v := range intNode.values {
			q.push(v)
		}
		for _, v := range intNode.values {
			recursiveBfs(v, q)
		}
		return
	}
}
