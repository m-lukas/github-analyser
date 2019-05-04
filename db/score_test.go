package db

//INTEGRATION TEST FOR SCORECONFIG -> takes ~0.1 seconds
/*
func Test_Score(t *testing.T) {

	var err error

	dbRoot = &DatabaseRoot{}

	t.Run("scoreconfig initialization doesn't work", func(t *testing.T) {
		err = dbRoot.initScoreConfig()
		require.Nil(t, err, "failed to initialize score config")

		scoreConfig := dbRoot.ScoreConfig
		require.NotNil(t, scoreConfig)

		assert.Equal(t, 1.0, scoreConfig.FollowingK)
		assert.Equal(t, 1.0, scoreConfig.ForksK)
		assert.Equal(t, 1.0, scoreConfig.ReposW)
	})

	dbRoot = nil
}
*/
