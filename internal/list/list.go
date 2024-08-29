package list

// Node is an element in a linked list.
type Node[T any] struct {
	prev, next *Node[T]
	list       *List[T]

	Value T
}

// Next returns the next item in the list.
func (e *Node[T]) Next() *Node[T] {
	if e.list == nil || e.next == &e.list.root {
		return nil
	}
	return e.next
}

// Prev returns the previous item in the list.
func (e *Node[T]) Prev() *Node[T] {
	if e.list == nil || e.prev == &e.list.root {
		return nil
	}
	return e.prev
}

// List implements a generic linked list based off of container/list. This
// contains the minimimum functionally required for an lru cache.
type List[T any] struct {
	root Node[T]
	len  int
}

// NewList creates a new linked list.
func NewList[T any]() *List[T] {
	l := &List[T]{}
	l.Init()
	return l
}

// Init intializes the list with no elements.
func (l *List[T]) Init() {
	l.root = Node[T]{}
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
}

// Len is the number of elements in the list.
func (l *List[T]) Len() int {
	return l.len
}

// InsertNewNode adds a new value to the front of the list.
func (l *List[T]) InsertNewNode(value T) *Node[T] {
	node := &Node[T]{Value: value}
	at := &l.root
	node.prev = at
	node.next = at.next
	node.prev.next = node
	node.next.prev = node
	node.list = l
	l.len++

	return node
}

// GetLastNode returns the last element in the list.
func (l *List[T]) GetLastNode() *Node[T] {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

// Remove removes the given element from the list.
func (l *List[T]) Remove(e *Node[T]) T {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.list.len--
	e.next = nil
	e.prev = nil
	e.list = nil
	return e.Value
}
