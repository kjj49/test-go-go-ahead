package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host     = "app:8080"
	basePath = "http://" + host
)

func TestMain(m *testing.M) {
	err := checkBasePathAvailability(20)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func checkBasePathAvailability(attempts int) error {
	var err error

	for attempts > 0 {
		resp, err := http.Get(basePath)
		if err == nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}

		log.Printf("Integration tests: URL %s is not available, attempts left: %d", basePath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

func TestGetCurrencyByParamsIntegration(t *testing.T) {
	t.Parallel()

	// Test case: Valid parameters
	resp, err := http.Get(basePath + "/currency?val=USD&date=22.05.2024")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test case: Missing val parameter
	resp, err = http.Get(basePath + "/currency?date=22.05.2024")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Test case: Invalid currency code
	resp, err = http.Get(basePath + "/currency?val=US&date=22.05.2024")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Test case: Invalid date format
	resp, err = http.Get(basePath + "/currency?val=USD&date=2024-05-22")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
