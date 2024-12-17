package mybtree

type stack[T any] []T

func (s *stack[T]) push(value T) {
	*s = append(*s, value)
}

func (s *stack[T]) pop() (T, bool) {
	var result T
	if len(*s) == 0 {
		return result, false
	}
	result = (*s)[len(*s)-1]
	if len(*s) > 1 {
		*s = (*s)[:len(*s)-2]
	} else {
		*s = (*s)[0:0]
	}
	return result, true
}

type queue[T any] []T

func (q *queue[T]) push(value T) {
	*q = append(*q, value)
}

func (q *queue[T]) pop() (T, bool) {
	var result T
	if len(*q) == 0 {
		return result, false
	}
	result = (*q)[0]
	*q = (*q)[1:]
	return result, true
}

func insertAt[K any](s *[]K, value K, pos int) {
	if pos == len(*s) {
		*s = append(*s, value)
		return
	}

	// copy last element at the end
	*s = append(*s, (*s)[len(*s)-1])
	// shift values
	for i := len(*s) - 2; i > pos; i-- {
		(*s)[i] = (*s)[i-1]
	}
	(*s)[pos] = value
}

func orderedInsert[K keyAble](s *[]K, value K) int {
	if len(*s) == 0 {
		*s = append(*s, value)
		return 0
	}
	var pos int = -1
	for i, v := range *s {
		if value < v {
			pos = i
			break
		}
	}

	if pos == -1 {
		// none of the elements is bigger than the new key, so it should be appended at the end
		pos = len(*s)
	}

	insertAt(s, value, pos)
	return pos
}
