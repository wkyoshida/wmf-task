package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assert := assert.New(t)

	var expected = env{
		ApiPort:        "3000",
		MainMWInstance: "en",
		MWInstanceConfigs: []MWInstanceConfig{
			{
				ApiUrl: "https://en.wikipedia.org/w/api.php",
				ID:     "en",
				User:   "",
				Pass:   "",
			},
			{
				ApiUrl: "https://pt.wikipedia.org/w/api.php",
				ID:     "pt",
				User:   "",
				Pass:   "",
			},
		},
	}

	err := Load("../../env.yaml")
	assert.NoError(err)
	assert.Equal(expected, Env)
}
