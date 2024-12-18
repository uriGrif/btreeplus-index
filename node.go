package mybtree

type keyAble interface {
	comparable
	~int | ~uint64 | ~string | ~float64
}

type Indexable[K keyAble] interface {
	GetKey() K
}

type node[K keyAble, T Indexable[K]] interface {
	printKeys()
	get(K) *T
	insert(*T, *node[K, T], stack[node[K, T]])
	delete(K)
	getMinKey() K
	getMaxKey() K
	valuesAmount() int
}
