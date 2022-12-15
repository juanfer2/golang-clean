package utilities

import (
	"os"
	"testing"

	"github.com/juanfer2/golang-clean/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("MODE", "TEST")
	assert := assert.New(t)
	value := utilities.GetEnv("MODE")

	assert.Equal(value, "TEST", "they should be truthy")
}
