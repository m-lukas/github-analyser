package util_test //black-box testing

import (
	"fmt"
	"os"
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/assert"
)

const (
	write_test = "./test/test_write.txt"
	read_test  = "./test/test_read.txt"
)

func Test_Files(t *testing.T) {

	t.Run("WriteFile(): not able to write file", func(t *testing.T) {
		var err error

		err = util.WriteFile(write_test, []string{"test1", "test2", "test3"})
		fmt.Println(err)
		assert.Nil(t, err)

		err = os.Remove(write_test)
		assert.Nil(t, err)
	})

	t.Run("ReadLines(): not able to read lines in file", func(t *testing.T) {
		expected := []string{"hallo", "hello", "salut"}

		output, err := util.ReadLines(read_test)
		assert.Nil(t, err)

		assert.Equal(t, expected, output)
	})

	t.Run("HasFileFormat(): doesn't recognise file format", func(t *testing.T) {
		assert.True(t, util.HasFileFormat("./util/test/test_write.txt", "txt"))
		assert.True(t, util.HasFileFormat("file.txt", "txt"))
		assert.False(t, util.HasFileFormat("./something/image.png", "jpg"))
		assert.True(t, util.HasFileFormat("./whatever/document.docx", "docx"))
	})

	t.Run("ReadInputFiles(): (integration) failed to retrive input array", func(t *testing.T) {
		var output []string
		var err error
		var expected = []string{"hallo", "hello", "salut"}

		output, err = util.ReadInputFiles([]string{read_test})
		assert.Nil(t, err)
		assert.Equal(t, expected, output)

		output, err = util.ReadInputFiles([]string{read_test, "file_name.png"})
		assert.Error(t, err)

		output, err = util.ReadInputFiles([]string{read_test, read_test})
		assert.Nil(t, err)
		assert.Equal(t, expected, output)
	})

	t.Run("ReadFile(): not able to read file (query)", func(t *testing.T) {
		expected := "hallo\nhello\nsalut"

		output, err := util.ReadFile(read_test)
		assert.Nil(t, err)

		assert.Equal(t, expected, output)
	})

}
