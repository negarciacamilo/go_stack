package go_stack

import "sync"

type threadSafeStack struct {
	data []any
	mtx  sync.RWMutex
}

// Don't NewThreadSafeStack use this, not tested yet
func NewThreadSafeStack[T any]() Stack {
	return &threadSafeStack{data: []any{}}
}

func (t *threadSafeStack) IsEmpty() bool {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return len(t.data) == 0
}

func (t *threadSafeStack) Push(elem any) {
	t.mtx.Lock()
	t.data = append(t.data, elem)
	t.mtx.Unlock()
}

func (t *threadSafeStack) Pop() any {
	defer t.mtx.Unlock()
	if !t.IsEmpty() {
		t.mtx.Lock()
		elem := t.data[len(t.data)-1]
		t.data = t.data[:len(t.data)-1]
		return elem
	}
	return nil
}

func (t *threadSafeStack) Size() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return len(t.data)
}

func (t *threadSafeStack) Peek() any {
	defer t.mtx.RUnlock()
	if !t.IsEmpty() {
		t.mtx.RLock()
		return t.data[t.Size()-1]
	}
	return nil
}

func (t *threadSafeStack) Clear() {
	t.data = []any{}
}
