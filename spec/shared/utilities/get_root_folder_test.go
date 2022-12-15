package utilities

import (
	"regexp"
	"testing"

	"github.com/juanfer2/golang-clean/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetRootFolder(t *testing.T) {
	assert := assert.New(t)
	folder := utilities.GetRootFolder()
	match, _ := regexp.MatchString("golang-clean", folder)

	assert.Equal(match, true, "they should be truthy")
}
