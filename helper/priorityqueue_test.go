package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue[int, string]()

	require.Equal(t, 0, queue.Len())

	t.Run("PopEmpty", func(t *testing.T) {
		obj, p, ok := queue.Pop()
		require.Equal(t, "", obj)
		require.Equal(t, 0, p)
		require.False(t, ok)
	})

	t.Run("PushPop", func(t *testing.T) {
		queue.Push(1, "asdf")
		require.Equal(t, 1, queue.Len())

		obj, p, ok := queue.Pop()
		require.Equal(t, "asdf", obj)
		require.Equal(t, 1, p)
		require.True(t, ok)

		require.Equal(t, 0, queue.Len())
	})

	t.Run("PushPushPopPop", func(t *testing.T) {
		queue.Push(34, "foo")
		queue.Push(4, "best")
		queue.Push(42, "bar")
		require.Equal(t, 3, queue.Len())

		obj1, p1, ok1 := queue.Pop()
		require.Equal(t, "best", obj1)
		require.Equal(t, 4, p1)
		require.True(t, ok1)
		require.Equal(t, 2, queue.Len())

		obj2, p2, ok2 := queue.Pop()
		require.Equal(t, "foo", obj2)
		require.Equal(t, 34, p2)
		require.True(t, ok2)
		require.Equal(t, 1, queue.Len())

		obj3, p3, ok3 := queue.Pop()
		require.Equal(t, "bar", obj3)
		require.Equal(t, 42, p3)
		require.True(t, ok3)
		require.Equal(t, 0, queue.Len())
	})
}
