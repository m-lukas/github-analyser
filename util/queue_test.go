package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	t.Run("not able to pop item from queue", func(t *testing.T) {
		queue := &Queue{}
		queue.elements = append(queue.elements, "test1")
		queue.elements = append(queue.elements, "test2")
		queue.elements = append(queue.elements, "test3")
		item1 := queue.Pop()
		assert.Equal(t, "test1", item1)
		item2 := queue.Pop()
		assert.Equal(t, "test2", item2)
		item3 := queue.Pop()
		assert.Equal(t, "test3", item3)
	})
	t.Run("not able to push item to queue", func(t *testing.T) {
		queue := &Queue{}
		queue.Push("test1")
		queue.Push("test2")
		queue.Push("test3")
		assert.Equal(t, "test1", queue.elements[0])
		assert.Equal(t, "test2", queue.elements[1])
		assert.Equal(t, "test3", queue.elements[2])
	})
	t.Run("queue integration", func(t *testing.T) {
		var lenght int
		var item interface{}

		queue := &Queue{}
		queue.Push("test1")
		queue.Push("test2")
		item = queue.Pop()
		assert.Equal(t, "test1", item)
		item = queue.Pop()
		assert.Equal(t, "test2", item)

		lenght = len(queue.elements)
		assert.Equal(t, 0, lenght)

		queue.Push("test3")
		lenght = len(queue.elements)
		assert.Equal(t, 1, lenght)
		item = queue.Pop()
		assert.Equal(t, "test3", item)

		lenght = len(queue.elements)
		assert.Equal(t, 0, lenght)
	})
}
