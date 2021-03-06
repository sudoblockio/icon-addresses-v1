package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// List test
func TestAddressesEndpointList(t *testing.T) {
	assert := assert.New(t)

	addressesServiceURL := os.Getenv("ADDRESSES_SERVICE_URL")
	if addressesServiceURL == "" {
		addressesServiceURL = "http://localhost:8000"
	}
	addressesServiceRestPrefx := os.Getenv("ADDRESSES_SERVICE_REST_PREFIX")
	if addressesServiceRestPrefx == "" {
		addressesServiceRestPrefx = "/api/v1"
	}

	resp, err := http.Get(addressesServiceURL + addressesServiceRestPrefx + "/addresses")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List limit and skip test
func TestAddressesEndpointListLimitSkip(t *testing.T) {
	assert := assert.New(t)

	addressesServiceURL := os.Getenv("ADDRESSES_SERVICE_URL")
	if addressesServiceURL == "" {
		addressesServiceURL = "http://localhost:8000"
	}
	addressesServiceRestPrefx := os.Getenv("ADDRESSES_SERVICE_REST_PREFIX")
	if addressesServiceRestPrefx == "" {
		addressesServiceRestPrefx = "/api/v1"
	}

	resp, err := http.Get(addressesServiceURL + addressesServiceRestPrefx + "/addresses?limit=20&skip=20")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List is contract test
func TestAddressesEndpointListIsContract(t *testing.T) {
	assert := assert.New(t)

	addressesServiceURL := os.Getenv("ADDRESSES_SERVICE_URL")
	if addressesServiceURL == "" {
		addressesServiceURL = "http://localhost:8000"
	}
	addressesServiceRestPrefx := os.Getenv("ADDRESSES_SERVICE_REST_PREFIX")
	if addressesServiceRestPrefx == "" {
		addressesServiceRestPrefx = "/api/v1"
	}

	resp, err := http.Get(addressesServiceURL + addressesServiceRestPrefx + "/addresses?is_contract=true")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}
