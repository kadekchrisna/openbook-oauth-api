package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kadekchrisna/openbook-oauth/src/domain/users"
	"github.com/kadekchrisna/openbook-oauth/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

type (
	RestUsersRepository interface {
		LoginUser(string, string) (*users.User, *errors.ResErr)
	}
	userRepository struct{}
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:3080",
		Timeout: 100 * time.Millisecond,
	}
)

// NewRestUsersRepository NewRestUsersRepository
func NewRestUsersRepository() RestUsersRepository {
	return &userRepository{}
}

func (s *userRepository) LoginUser(email string, password string) (*users.User, *errors.ResErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/user/auth", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid client to login user")
	}

	if response.StatusCode > 299 {
		var errRes errors.ResErr
		if err := json.Unmarshal(response.Bytes(), &errRes); err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface at login user")
		}
		return nil, &errRes
	}

	var user users.User
	fmt.Println(string(response.Bytes()))
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Invalid user interface at login user")
	}
	return &user, nil

}
