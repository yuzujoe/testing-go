package providers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoutryRestClientError(t *testing.T) {
	// Execution:
	coutry, err := GetCoutry("AR")

	// Validation:
	assert.Nil(t, coutry)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid restclient error when getting coutry AR", err.Message)
}
func TestGetCoutryNotFound(t *testing.T) {
	coutry, err := GetCoutry("AR")

	assert.Nil(t, coutry)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "country not found", err.Message)
}
func TestGetCoutryInvalidErrorIterface(t *testing.T) {
	coutry, err := GetCoutry("AR")

	assert.Nil(t, coutry)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when getting country AR", err.Message)
}
func TestGetCoutryInvalidJsonResponse(t *testing.T) {
	coutry, err := GetCoutry("AR")

	assert.Nil(t, coutry)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal countryndata for AR", err.Message)
}
func TestGetCoutryNotError(t *testing.T) {
	coutry, err := GetCoutry("AR")

	assert.Nil(t, err)
	assert.NotNil(t, coutry)
	assert.EqualValues(t, "AR", coutry.ID)
	assert.EqualValues(t, "AR", coutry.Name)
	assert.EqualValues(t, "AR", coutry.TimeZone)
	assert.EqualValues(t, 24, len(coutry.States))
}
