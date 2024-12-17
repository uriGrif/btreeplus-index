package mybtree

import "fmt"

type leafNode[K keyAble, T Indexable[K]] struct {
	degree   int
	keys     []K
	values   []*T
	prevNode *leafNode[K, T]
	nextNode *leafNode[K, T]
}

func (n *leafNode[K, T]) printKeys() {
	fmt.Printf("%v", n.keys)
}

func (n *leafNode[K, T]) getMinKey() K {
	return n.keys[0]
}

func (n *leafNode[K, T]) getMaxKey() K {
	return n.keys[len(n.keys)-1]
}

func (n *leafNode[K, T]) get(key K) *T {
	for _, v := range n.values {
		if (*v).GetKey() == key {
			return v
		}
	}
	return nil
}

func (n *leafNode[K, T]) insert(value *T, root *node[K, T], parentStack stack[node[K, T]]) {

	if len(n.keys) == n.degree {
		n.split(value, root, parentStack)
		return
	}

	valueKey := (*value).GetKey()

	insertPos := orderedInsert(&n.keys, valueKey)
	insertAt(&n.values, value, insertPos)
}

func (n *leafNode[K, T]) split(value *T, root *node[K, T], parentStack stack[node[K, T]]) {
	valueKey := (*value).GetKey()

	// create new node
	newNode := &leafNode[K, T]{
		degree:   n.degree,
		keys:     make([]K, 0, n.degree),
		values:   make([]*T, 0, n.degree),
		prevNode: n,
		nextNode: n.nextNode,
	}
	var splitPos int = n.degree / 2
	for i := splitPos + 1; i < n.degree; i++ {
		newNode.keys = append(newNode.keys, n.keys[i])
		newNode.values = append(newNode.values, n.values[i])
	}
	n.keys = n.keys[:splitPos+1]
	n.values = n.values[:splitPos+1]

	// add new value to corresponding node
	if valueKey <= n.keys[len(n.keys)-1] {
		n.insert(value, root, parentStack)
	} else {
		newNode.insert(value, root, parentStack)
	}

	// update parent node
	parentNode, ok := parentStack.pop()
	if !ok {
		// this is the root node
		parentNode = &internalNode[K, T]{
			degree: n.degree,
			keys:   make([]K, 0, n.degree),
			values: make([]node[K, T], 0, n.degree),
		}
		*root = parentNode
	}
	p := (parentNode).(*internalNode[K, T])
	p.handleChildrenSplit(root, parentStack, n, newNode)
}
