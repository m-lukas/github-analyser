package util_test //black-box testing

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/assert"
)

const (
	slice_size = 5000
)

func Test_Contains(t *testing.T) {
	slice1 := []string{"Hello", "test123", "bliblablo", "m-lukas", "greetings", "CODE University", "testcase", "fun"}
	slice2 := []string{"Hamburg", "berlin", "Munich", "Görlitz", "Germany", "United States"}
	slice3 := []string{"sindresorhus", "Ouka", "bitliner", "m-lukas", "Urhengulas", "alexmorten", "nat", "steve"}

	testTable := []struct {
		Slice  []string
		Query  string
		Output bool
	}{
		{slice1, "test123", true},
		{slice1, "code University", false},
		{slice1, "Frankfurt", false},
		{slice2, "Berlin", false},
		{slice2, "Görlitz", true},
		{slice2, "United States", true},
		{slice3, "sindresorhus", true},
		{slice3, "alexmorten1", false},
		{slice3, "ouka", false},
	}

	t.Run("normal: contains", func(t *testing.T) {
		for _, test := range testTable {
			assert.Equal(t, test.Output, util.Contains(test.Slice, test.Query))
		}
	})
	t.Run("binary: contains", func(t *testing.T) {
		for _, test := range testTable {
			assert.Equal(t, test.Output, util.BinaryContains(test.Slice, test.Query))
		}
	})
}

func Benchmark_Contains(b *testing.B) {
	slice, contained := generateStringSlice(slice_size)

	b.Run("normal: benchmark", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			util.Contains(slice, contained)
		}
	})
	b.Run("binary: benchmark", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			util.BinaryContains(slice, contained)
		}
	})
}

func generateStringSlice(lenght int) ([]string, string) {

	var resSlice []string

	dictionary := []string{"hello", "test", "lukas", "code", "berlin"}
	for i := 0; i < lenght; i++ {
		base := dictionary[rand.Intn(len(dictionary)-1)]
		item := fmt.Sprintf("%s%d", base, i)
		resSlice = append(resSlice, item)
	}

	contains := resSlice[rand.Intn(len(resSlice)-1)]

	return resSlice, contains

}
