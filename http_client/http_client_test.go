package http_client_test

import (
	"testing"

	"github.com/eds000n/hashing/http_client"
)

func TestDoRequest(t *testing.T) {
	testCases := []struct {
		url              string
		expectedResponse [16]byte
		expectedErr      error
	}{
		{
			url:              "http://adjust.com", // FIXME: flaky test, should choose an url which is guaranteed not to change
			expectedResponse: [16]byte{70, 164, 226, 199, 157, 231, 3, 243, 190, 59, 71, 242, 108, 235, 32, 94},
		},
	}
	for _, tc := range testCases {
		response, err := http_client.DoRequest(tc.url)

		if response != tc.expectedResponse {
			t.Fatalf("expected: %v, got: %v", tc.expectedResponse, response)
		}

		if err != tc.expectedErr {
			t.Fatalf("expected error: %v, got: %v", tc.expectedErr, err)
		}
	}
}
