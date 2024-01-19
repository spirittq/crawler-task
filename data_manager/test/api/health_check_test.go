package api_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {

	expectedReponse := []byte(`{"healthCheck": "OK"}`)

	resp, _ := http.Get("http://localhost:3000/health_check")
	require.Equal(t, http.StatusOK, resp.StatusCode)
	bodyBytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, expectedReponse, bodyBytes)
}
