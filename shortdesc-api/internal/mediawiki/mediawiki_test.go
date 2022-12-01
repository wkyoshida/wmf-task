package mediawiki

import (
	"strings"
	"testing"

	"shortdesc-api/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	var testConfig config.MWInstanceConfig
	err := Init(testConfig)
	assert.NoError(err)
	assert.NotNil(client)
	assert.True(strings.Contains(client.UserAgent, userAgent))

	err = Init(testConfig)
	assert.Equal(ErrClientExists, err)
}

func TestConstructTitlesParam(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name     string
		titles   []string
		expected string
	}{
		{
			"multiple titles",
			[]string{"foo", "bar"},
			"foo|bar",
		},
		{
			"single title",
			[]string{"foo"},
			"foo",
		},
		{
			"empty title",
			[]string{},
			"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := constructTitlesParam(tc.titles)
			assert.Equal(tc.expected, res)
		})
	}
}

func TestParseShortDescRaw(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name        string
		content     string
		expectedRes string
		expectedErr error
	}{
		{
			"common short description",
			"foobar {{Short description|Canadian computer scientist}} foobar",
			"{{Short description|Canadian computer scientist}}",
			nil,
		},
		{
			"empty short description",
			"{{Short description| }}",
			"{{Short description| }}",
			nil,
		},
		{
			"no short description",
			"foobar",
			"",
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseShortDescRaw(tc.content)
			assert.Equal(tc.expectedRes, res)
			assert.Equal(tc.expectedErr, err)
		})
	}
}

func TestParseShortDesc(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name         string
		shortDescRaw string
		expectedRes  string
		expectedErr  error
	}{
		{
			"common short description",
			"{{Short description|Canadian computer scientist}}",
			"Canadian computer scientist",
			nil,
		},
		{
			"empty short description",
			"{{Short description| }}",
			" ",
			nil,
		},
		{
			"no short description",
			"foobar",
			"",
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseShortDesc(tc.shortDescRaw)
			assert.Equal(tc.expectedRes, res)
			assert.Equal(tc.expectedErr, err)
		})
	}
}
