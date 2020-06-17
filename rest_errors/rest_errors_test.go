package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		err.Message(), err.Status(), "internal_server_error", err.Causes()), err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		err.Message(), err.Status(), "bad_request", err.Causes()), err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		err.Message(), err.Status(), "not_found", err.Causes()), err.Error())
}
func TestNewUnauthorized(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		err.Message(), err.Status(), "unauthorized", err.Causes()), err.Error())
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("this is the message", http.StatusInternalServerError, "internal_server_error", nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		err.Message(), err.Status(), "internal_server_error", err.Causes()), err.Error())
	assert.Nil(t, err.Causes())
}
