package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		key         string
		value       string
		expected    string
		expecterErr string
	}{
		"correct": {
			key:         "Authorization",
			value:       "ApiKey 12345",
			expected:    "12345",
			expecterErr: "",
		},
		"wrong header": {
			key:         "NotAuthorization",
			value:       "ApiKey 12345",
			expected:    "",
			expecterErr: "no authorization header included",
		},
		"wrong tag": {
			key:         "Authorization",
			value:       "NotApiKey 12345",
			expected:    "",
			expecterErr: "malformed authorization header",
		},
		"no value": {
			key:         "Authorization",
			value:       "ApiKey",
			expected:    "",
			expecterErr: "malformed authorization header",
		},
		"CI failure": {
			key:         "Authorization",
			value:       "ApiKey 123",
			expected:    "1234",
			expecterErr: "malformed authorization header",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)
			got, err := GetAPIKey(header)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
			if err != nil && err.Error() != tc.expecterErr {
				t.Fatalf("expected: %v, got: %v", tc.expecterErr, err)
			}
		})
	}
}
