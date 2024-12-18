package mybtree

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	var s stack[int]

	// Test pushing elements
	s.push(1)
	s.push(2)
	s.push(3)

	if len(s) != 3 {
		t.Errorf("expected stack length 3, got %d", len(s))
	}

	// Test popping elements
	value, ok := s.pop()
	if !ok || value != 3 {
		t.Errorf("expected pop value 3, got %d", value)
	}

	value, ok = s.pop()
	if !ok || value != 2 {
		t.Errorf("expected pop value 2, got %d", value)
	}

	value, ok = s.pop()
	if !ok || value != 1 {
		t.Errorf("expected pop value 1, got %d", value)
	}

	// Test popping from empty stack
	_, ok = s.pop()
	if ok {
		t.Errorf("expected pop to return false for empty stack")
	}
}

func TestQueuePushPop(t *testing.T) {
	var q queue[int]

	// Test pushing elements
	q.push(1)
	q.push(2)
	q.push(3)

	if len(q) != 3 {
		t.Errorf("expected queue length 3, got %d", len(q))
	}

	// Test popping elements
	value, ok := q.pop()
	if !ok || value != 1 {
		t.Errorf("expected pop value 1, got %d", value)
	}

	value, ok = q.pop()
	if !ok || value != 2 {
		t.Errorf("expected pop value 2, got %d", value)
	}

	value, ok = q.pop()
	if !ok || value != 3 {
		t.Errorf("expected pop value 3, got %d", value)
	}

	// Test popping from empty queue
	_, ok = q.pop()
	if ok {
		t.Errorf("expected pop to return false for empty queue")
	}
}

func TestInsertAt(t *testing.T) {
	s := []int{1, 2, 4}

	// Test inserting in the middle
	insertAt(&s, 3, 2)
	expected := []int{1, 2, 3, 4}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}

	// Test inserting at the beginning
	insertAt(&s, 0, 0)
	expected = []int{0, 1, 2, 3, 4}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}

	// Test inserting at the end
	insertAt(&s, 5, len(s))
	expected = []int{0, 1, 2, 3, 4, 5}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}
}

type intKey int

func (i intKey) Less(other intKey) bool {
	return i < other
}

func TestOrderedInsert(t *testing.T) {
	s := []intKey{1, 3, 5}

	// Test inserting in the middle
	pos := orderedInsert(&s, 4)
	expected := []intKey{1, 3, 4, 5}
	if pos != 2 {
		t.Errorf("expected position 2, got %d", pos)
	}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}

	// Test inserting at the beginning
	pos = orderedInsert(&s, 0)
	expected = []intKey{0, 1, 3, 4, 5}
	if pos != 0 {
		t.Errorf("expected position 0, got %d", pos)
	}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}

	// Test inserting at the end
	pos = orderedInsert(&s, 6)
	expected = []intKey{0, 1, 3, 4, 5, 6}
	if pos != len(expected)-1 {
		t.Errorf("expected position %d, got %d", len(expected)-1, pos)
	}
	for i, v := range s {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, s)
		}
	}
}
