package main

import (
	"fmt"
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
			name:     "ENVIRONMENT variable set",
			envVar:   "ENVIRONMENT",
			envValue: "test_value",
			expected: "Hello from argocd-example! Reading variable from test_value\n",
		},
		{
			name:     "ENVIRONMENT environment variable not set",
			envVar:   "ENVIRONMENT",
			envValue: "",
			expected: "Hello from argocd-example! Reading variable from ENVIRONMENT variable not set\n",
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
				env := os.Getenv("ENVIRONMENT")
				if env == "" {
					env = "ENVIRONMENT variable not set"
				}
				fmt.Fprintf(w, "Hello from argocd-example! Reading variable from "+env+"\n")
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
