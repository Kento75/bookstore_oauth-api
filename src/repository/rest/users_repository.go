package rest

import (
	"encoding/json"
	"time"

	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token/utils/errors"
	"github.com/Kento75/bookstore_oauth-api/src/domain/users"
	"github.com/federicoleon/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	// README -> https://github.com/federicoleon/golang-restclient
	// bookstore_users-api
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)

		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
