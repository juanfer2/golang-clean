package utilities

import (
	"errors"
	"os"
	"testing"

	"github.com/juanfer2/golang-clean/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {

	os.Setenv("MODE", "TEST")
	assert := assert.New(t)
	value := GetEnv("MODE")

	assert.Equal(value, "TEST", "they should be truthy")
	t.Errorf("Error loading .env file fila")

	tests := []struct {
		name         string
		mockFunc     func()
		expected     string
		expectingErr bool
	}{
		{
			name: "test enviroment exist",
			mockFunc: func() {
				utilities.ValidateEnvFile = func() (err error) {
					err = errors.New("Error loading .env file")

					return err
				}
			},
			expected:     "",
			expectingErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()

			if tc.expectingErr {
				tt.Errorf("Error expectation not met, want %v", tc.expectingErr)
			}
		})
	}
}
