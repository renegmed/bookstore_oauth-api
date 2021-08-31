package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"bookstore_oauth-api/src/domain/users"

	"github.com/renegmed/bookstore_utils-go/rest_errors"

	//"github.com/mercadolibre/golang-restclient/rest"
	resty "github.com/go-resty/resty/v2"
)

// var (
// 	// usersRestClient = rest.RequestBuilder{
// 	// 	BaseURL: "http://localhost:8082",
// 	// 	Timeout: 100 * time.Millisecond,
// 	// }
// 	userRestClient = resty.New()
// )

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {
	Client *resty.Client
}

func NewRestUsersRepository() *usersRepository {
	client := resty.New()
	return &usersRepository{client}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	//log.Println("...post /users/login\n", request)

	//response := usersRestClient.Post("/users/login", request)

	response, err := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "appication/json").
		SetBody(request).
		Post("/users/login")
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid error interface when trying to post users login", err)
	}

	log.Println("... status code 1", response.StatusCode())

	if response == nil || response.RawResponse == nil || response.StatusCode() == http.StatusNotFound {
		return nil, rest_errors.NewNotFoundError("invalid restclient response when trying to login user")
	}

	log.Println("... status code 2", response.StatusCode())

	if response.RawResponse.StatusCode > 299 { // there is an error
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	//log.Println("...response post /users/login\n", response)

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil
}
