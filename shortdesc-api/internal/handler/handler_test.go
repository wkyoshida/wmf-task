package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name         string
		method       string
		target       string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "missing 'titles' parameter",
			method:       http.MethodGet,
			target:       "/shortdesc",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message": "Bad request. Query parameter 'titles' not present."}`,
		},
		{
			name:         "method not allowed",
			method:       http.MethodPost,
			target:       "/shortdesc",
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: `{"message": "Can't find method requested."}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				tc.method,
				tc.target,
				nil,
			)
			ServeHTTP(w, r)
			assert.Equal(tc.expectedCode, w.Code)
			assert.Equal(tc.expectedBody, w.Body.String())
		})
	}
}
