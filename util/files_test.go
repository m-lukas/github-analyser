package util

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	write_test = "./test/test_write.txt"
	read_test  = "./test/test_read.txt"
)

func Test_Files(t *testing.T) {
	t.Run("not able to write file", func(t *testing.T) {
		var err error

		err = WriteFile(write_test, []string{"test1", "test2", "test3"})
		fmt.Println(err)
		assert.Nil(t, err)

		err = os.Remove(write_test)
		assert.Nil(t, err)
	})
	t.Run("not able to read file", func(t *testing.T) {
		expected := []string{"hallo", "hello", "salut"}

		output, err := readFile(read_test)
		assert.Nil(t, err)

		assert.Equal(t, expected, output)
	})
	t.Run("doesn't recognise file format", func(t *testing.T) {
		assert.True(t, hasFileFormat("./util/test/test_write.txt", "txt"))
		assert.True(t, hasFileFormat("file.txt", "txt"))
		assert.False(t, hasFileFormat("./something/image.png", "jpg"))
		assert.True(t, hasFileFormat("./whatever/document.docx", "docx"))
	})
	t.Run("integration: failed to retrive input array", func(t *testing.T) {
		var output []string
		var err error
		var expected = []string{"hallo", "hello", "salut"}

		output, err = ReadInputFiles([]string{read_test})
		assert.Nil(t, err)
		assert.Equal(t, expected, output)

		output, err = ReadInputFiles([]string{read_test, "file_name.png"})
		assert.Error(t, err)

		output, err = ReadInputFiles([]string{read_test, read_test})
		assert.Nil(t, err)
		assert.Equal(t, expected, output)
	})

}
