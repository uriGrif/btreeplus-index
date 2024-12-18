package mybtree

import (
	"fmt"
	"testing"
)

type indexableInt int

func (n indexableInt) GetKey() int {
	return int(n)
}

func TestInsertion(t *testing.T) {
	tree := BtreeIndex[int, indexableInt]{
		Degree: 3,
		Root:   nil,
		Unique: true,
	}
	values := []indexableInt{1, 2, 4, 5, 6, 3, 16, 10, 12, 7, 8, 13, 14, 17, 18, 19, 20}
	for _, v := range values {
		tree.Insert(&v)
		fmt.Printf("inserted: %v - Tree:\n", v)
		tree.levelTraverse(func(n node[int, indexableInt]) {
			n.printKeys()
		})
		fmt.Printf("\n")
	}
}
