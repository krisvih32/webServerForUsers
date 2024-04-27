package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOST(t *testing.T) {
	assert := assert.New(t)
	response, err := http.Post("localhost:8080?email=nelsonmichael?example.com&phoneNo=001-992-365-5248&firstName=Colleen&la", "", nil)
	assert.ErrorIs(err, http.ErrHandlerTimeout, "timed out")
	assert.Equal(response.StatusCode, http.StatusBadRequest, fmt.Sprintf("status bad request: %s", response.Body))
	assert.Equal(response.StatusCode, http.StatusInternalServerError, fmt.Sprintf("internal server error: %s", response.Body))
}
