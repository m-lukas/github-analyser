package util_test //black-box testing

import (
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	t.Run("not able to pop item from queue", func(t *testing.T) {
		queue := &util.Queue{}
		queue.Elements = append(queue.Elements, "test1")
		queue.Elements = append(queue.Elements, "test2")
		queue.Elements = append(queue.Elements, "test3")
		item1 := queue.Pop()
		assert.Equal(t, "test1", item1)
		item2 := queue.Pop()
		assert.Equal(t, "test2", item2)
		item3 := queue.Pop()
		assert.Equal(t, "test3", item3)
	})
	t.Run("not able to push item to queue", func(t *testing.T) {
		queue := &util.Queue{}
		queue.Push("test1")
		queue.Push("test2")
		queue.Push("test3")
		assert.Equal(t, "test1", queue.Elements[0])
		assert.Equal(t, "test2", queue.Elements[1])
		assert.Equal(t, "test3", queue.Elements[2])
	})
	t.Run("queue integration", func(t *testing.T) {
		var lenght int
		var item interface{}

		queue := &util.Queue{}
		queue.Push("test1")
		queue.Push("test2")
		item = queue.Pop()
		assert.Equal(t, "test1", item)
		item = queue.Pop()
		assert.Equal(t, "test2", item)

		lenght = len(queue.Elements)
		assert.Equal(t, 0, lenght)

		queue.Push("test3")
		lenght = len(queue.Elements)
		assert.Equal(t, 1, lenght)
		item = queue.Pop()
		assert.Equal(t, "test3", item)

		lenght = len(queue.Elements)
		assert.Equal(t, 0, lenght)
	})
}
