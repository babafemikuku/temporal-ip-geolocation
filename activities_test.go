package iplocate_test

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/babafemikuku/temporal-ip-geolocation/iplocate"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}

func TestGetIP(t *testing.T) {
	testsuite := &testsuite.WorkflowTestSuite{}
	env := testsuite.NewTestActivityEnvironment()

	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("127.0.0.1\n")),
	}

	ipActivities := &iplocate.IPActivities{
		HTTPClient: &MockHTTPClient{
			Response: mockResponse,
		},
	}

	env.RegisterActivity(ipActivities)

	val, err := env.ExecuteActivity(ipActivities.GetIP)
	if err != nil {
		log.Fatalf("Expected error to be nil, but got: %v", err)
	}

	var ip string
	val.Get(&ip)

	expectedIP := "127.0.0.1"
	assert.Equal(t, ip, expectedIP)
}

func TestGetLocationInfo(t *testing.T) {
	testsuite := &testsuite.WorkflowTestSuite{}
	env := testsuite.NewTestActivityEnvironment()

	mockResponse := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
			"city": "Ogba",
			"regionName": "Lagos",
			"country": "Nigeria"
			}`)),
	}

	ipActivities := &iplocate.IPActivities{
		HTTPClient: &MockHTTPClient{
			Response: mockResponse,
		},
	}

	env.RegisterActivity(ipActivities)

	ip := "127.0.0.1"
	val, err := env.ExecuteActivity(ipActivities.GetLocationInfo, ip)
	if err != nil {
		log.Fatalf("Expected error to be nil, but got %v:", err)
	}

	var location string
	val.Get(&location)

	expectedLocation := "Ogba, Lagos, Nigeria"
	assert.Equal(t, location, expectedLocation)
}
