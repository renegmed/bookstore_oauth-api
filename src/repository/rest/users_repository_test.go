package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func newResponder(s int, c string, ct string) httpmock.Responder {
	resp := httpmock.NewStringResponse(s, c)
	resp.Header.Set("Content-Type", ct)

	return httpmock.ResponderFromResponse(resp)
}

func TestLoginUserNoError(t *testing.T) {

	defer httpmock.DeactivateAndReset()

	repository := NewRestUsersRepository()

	httpmock.ActivateNonDefault(repository.Client.GetClient())
	httpmock.RegisterResponder(
		"POST", "/users/login",
		newResponder(200, `
		{"id": 1, "first_name": "Fede", "last_name": "León", "email": "fedeleon.cba@gmail.com"}
		`, "application/json"))

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "León", user.LastName)
	assert.EqualValues(t, "fedeleon.cba@gmail.com", user.Email)
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	repository := NewRestUsersRepository()

	httpmock.ActivateNonDefault(repository.Client.GetClient())
	httpmock.RegisterResponder(
		"POST", "/users/loginx", // wrong url
		newResponder(200, `
		{"id": 1, "first_name": "Fede", "last_name": "León", "email": "fedeleon.cba@gmail.com"}
		`, "application/json"))

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid error interface when trying to post users login", err.Message())
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

	defer httpmock.DeactivateAndReset()

	repository := NewRestUsersRepository()

	httpmock.ActivateNonDefault(repository.Client.GetClient())
	httpmock.RegisterResponder(
		"POST", "/users/login",
		newResponder(500, ``, "application/json"))

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message())
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	repository := NewRestUsersRepository()

	httpmock.ActivateNonDefault(repository.Client.GetClient())
	httpmock.RegisterResponder(
		"POST", "/users/login",
		newResponder(404, ``, "application/json"))

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message())
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

	defer httpmock.DeactivateAndReset()

	repository := NewRestUsersRepository()

	httpmock.ActivateNonDefault(repository.Client.GetClient())
	httpmock.RegisterResponder(
		"POST", "/users/login",
		newResponder(200, `
			{"id": 1, "first_name": "Fede", "last_name: "León", "email": "fedeleon.cba@gmail.com"}  
		`, "application/json")) // NOTE: data has invalid json data structure - last_name has missing closing quote

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message())
}
