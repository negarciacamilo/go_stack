package go_stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[string]()
	require.NotNil(t, s)
	require.IsType(t, &stack{}, s)
}

func TestPushShouldIncrementSize(t *testing.T) {
	s := NewStack[string]()
	s.Push("test")
	s.Push("a")
	s.Push("is")
	s.Push("this")
	require.Equal(t, 4, s.Size())
}

func TestPopFromStack(t *testing.T) {
	s := NewStack[string]()
	s.Push("test")
	s.Push("a")
	s.Push("is")
	s.Push("this")
	require.Equal(t, "this", s.Pop())
	require.Equal(t, "is", s.Pop())
	require.Equal(t, "a", s.Pop())
	require.Equal(t, "test", s.Pop())
	require.True(t, s.IsEmpty())
	require.Equal(t, 0, s.Size())
}

func TestPeek(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	require.Equal(t, 3, s.Size())
	require.Equal(t, 3, s.Peek())
}

func TestClear(t *testing.T) {
	s := NewStack[bool]()
	s.Push(true)
	s.Push(true)
	s.Push(false)
	require.Equal(t, 3, s.Size())
	s.Clear()
	require.Equal(t, 0, s.Size())
}
