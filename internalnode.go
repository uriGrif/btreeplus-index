package mybtree

import "fmt"

type internalNode[K keyAble, T Indexable[K]] struct {
	degree int
	keys   []K
	values []node[K, T]
}

func (n *internalNode[K, T]) printKeys() {
	fmt.Printf("%v", n.keys)
}

func (n *internalNode[K, T]) getMinKey() K {
	return n.keys[0]
}

func (n *internalNode[K, T]) getMaxKey() K {
	return n.keys[len(n.keys)-1]
}

func (n *internalNode[K, T]) nextNodePosition(key K) int {
	// basically a binary search, but not for an exact value
	var lo, hi int = 0, len(n.keys) - 1
	for lo < hi {
		var i int = (hi + lo) / 2

		if key < n.keys[i] {
			hi = i
		} else {
			lo = i + 1
		}
	}
	if lo >= len(n.keys) || key > n.keys[lo] {
		return -1
	}
	return lo
}

func (n *internalNode[K, T]) get(key K) *T {
	pos := n.nextNodePosition(key)
	if pos == -1 {
		return nil
	}
	return n.values[n.nextNodePosition(key)].get(key)
}

func (n *internalNode[K, T]) insert(value *T, root *node[K, T], parentStack stack[node[K, T]]) {
	pos := n.nextNodePosition((*value).GetKey())
	if pos == -1 {
		// I have to update the largest key
		n.keys[len(n.values)-1] = (*value).GetKey()
		pos = len(n.values) - 1
	}
	parentStack.push(n)
	n.values[pos].insert(value, root, parentStack)
}

func (n *internalNode[K, T]) split(root *node[K, T], parentStack stack[node[K, T]], child1, child2 node[K, T]) {
	// create new node
	newNode := &internalNode[K, T]{
		degree: n.degree,
		keys:   make([]K, 0, n.degree),
		values: make([]node[K, T], 0, n.degree),
	}
	var splitPos int = n.degree / 2
	for i := splitPos + 1; i < n.degree; i++ {
		newNode.keys = append(newNode.keys, n.keys[i])
		newNode.values = append(newNode.values, n.values[i])
	}
	n.keys = n.keys[:splitPos+1]
	n.values = n.values[:splitPos+1]

	// add new value to corresponding node
	newKey := child2.getMinKey()

	if newKey <= n.keys[len(n.keys)-1] {
		insertPos := orderedInsert(&n.keys, newKey)
		insertAt(&n.values, child1, insertPos)
	} else {
		insertPos := orderedInsert(&newNode.keys, newKey)
		insertAt(&newNode.values, child1, insertPos)
	}

	oldKey := child2.getMaxKey()
	if oldKey <= n.keys[len(n.keys)-1] {
		updatePos := n.nextNodePosition(oldKey)
		if updatePos == -1 {
			// value didn't exist
			n.values = append(n.values, child2)
		} else {
			n.values[updatePos] = child2
		}
	} else {
		updatePos := newNode.nextNodePosition(oldKey)
		if updatePos == -1 {
			// value didn't exist
			newNode.values = append(newNode.values, child2)
		} else {
			newNode.values[updatePos] = child2

		}
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
	p := (parentNode).(*internalNode[K, T]) // The parent node will always be an internalNode
	p.handleChildrenSplit(root, parentStack, n, newNode)
}

func (n *internalNode[K, T]) handleChildrenSplit(root *node[K, T], parentStack stack[node[K, T]], child1, child2 node[K, T]) {
	key1 := child2.getMinKey()

	if len(n.keys) == n.degree {
		// split this node
		n.split(root, parentStack, child1, child2)
		return
	}

	key2 := child2.getMaxKey()
	updatePos := n.nextNodePosition(key2)
	if updatePos == -1 {
		// value didn't exist
		n.keys = append(n.keys, key2)
		n.values = append(n.values, child2)
	} else {
		n.values[updatePos] = child2
	}

	insertPos := orderedInsert(&n.keys, key1)
	insertAt(&n.values, child1, insertPos)
}
