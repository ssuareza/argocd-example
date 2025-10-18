package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name     string
		envVar   string
		envValue string
		expected string
	}{
		{
			name:     "OUTPUT environment variable set",
			envVar:   "OUTPUT",
			envValue: "test_value",
			expected: "Hello from argocd-example! Reading variable from test_value\n",
		},
		{
			name:     "OUTPUT environment variable not set",
			envVar:   "OUTPUT",
			envValue: "",
			expected: "Hello from argocd-example! Reading variable from OUTPUT environment variable not set\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.envVar, tt.envValue)
				defer os.Unsetenv(tt.envVar)
			} else {
				os.Unsetenv(tt.envVar) // Ensure it's not set from previous tests
			}

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				output := os.Getenv("OUTPUT")
				if output == "" {
					output = "OUTPUT environment variable not set"
				}
				io.WriteString(w, "Hello from argocd-example! Reading variable from "+output+"\n")
			})

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}
