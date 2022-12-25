package utils_test

import (
	"hacktiv8-final-project-3/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpError(t *testing.T) {
	stringError := utils.NewHttpError("Message Test", "Test Error")

	assert.NotEmpty(t, stringError.Message)
	assert.NotEmpty(t, stringError.Trace)

	testStruct := struct {
		username string
	}{
		username: "wahyu_test",
	}

	structError := utils.NewHttpError("Error", testStruct)

	assert.NotEmpty(t, structError.Trace)
	assert.Equal(t, structError.Trace, testStruct)
}
