package utils_test

import (
	"os"
	"server/utils"
	"testing"
)

func TestLoadEnvironment(t *testing.T) {
	tests := []struct {
		name       string
		envVarName string
		want       string
	}{
		{
			name:       "Test loading TEST_ENV_VAR",
			envVarName: "TEST_ENV_VAR",
			want:       "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.LoadEnvironment()

			got := os.Getenv(tt.envVarName)

			if got != tt.want {
				t.Errorf("os.Getenv(%q) = %v; want %v", tt.envVarName, got, tt.want)
			}
		})
	}
}
