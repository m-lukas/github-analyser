package metrix

import (
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/assert"
)

func Test_Calculations(t *testing.T) {

	users := []*db.User{
		{
			Login:        "test1",
			Repositories: 20,
		},
		{
			Login:        "test2",
			Repositories: 2,
		},
		{
			Login:        "test3",
			Repositories: 55,
		},
		{
			Login:        "test4",
			Repositories: 14,
		},
		{
			Login:        "test5",
			Repositories: 67,
		},
	}

	t.Run("calculate k", func(t *testing.T) {
		expected := 0.2
		k := calcK(users, TYPE_REPOS)
		assert.Equal(t, expected, k)

		//cross-check
		k2 := util.CalcKFromY(50, 20)
		assert.Equal(t, k, k2)
	})

}
