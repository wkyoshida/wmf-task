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
			name:     "multiple titles",
			titles:   []string{"foo", "bar"},
			expected: "foo|bar",
		},
		{
			name:     "single title",
			titles:   []string{"foo"},
			expected: "foo",
		},
		{
			name:     "empty title",
			titles:   []string{},
			expected: "",
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
			name:        "common short description",
			content:     "foobar {{Short description|Canadian computer scientist}} foobar",
			expectedRes: "{{Short description|Canadian computer scientist}}",
			expectedErr: nil,
		},
		{
			name:        "empty short description",
			content:     "{{Short description| }}",
			expectedRes: "{{Short description| }}",
			expectedErr: nil,
		},
		{
			name:        "no short description",
			content:     "foobar",
			expectedRes: "",
			expectedErr: nil,
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
			name:         "common short description",
			shortDescRaw: "{{Short description|Canadian computer scientist}}",
			expectedRes:  "Canadian computer scientist",
			expectedErr:  nil,
		},
		{
			name:         "empty short description",
			shortDescRaw: "{{Short description| }}",
			expectedRes:  " ",
			expectedErr:  nil,
		},
		{
			name:         "no short description",
			shortDescRaw: "foobar",
			expectedRes:  "",
			expectedErr:  nil,
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
