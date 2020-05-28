package tokens

import (
	tokens "github.com/kadekchrisna/openbook-oauth-api/src/domain/access_token"
	"github.com/kadekchrisna/openbook-oauth-api/src/repository/db"
	"github.com/kadekchrisna/openbook-oauth-api/src/repository/rest"
	"github.com/kadekchrisna/openbook-oauth-api/src/utils/errors"
)

type (
	// Repository interface {
	// 	GetById(string) (*tokens.AccessToken, *errors.ResErr)
	// 	UpdateExpired(tokens.AccessToken) *errors.ResErr
	// 	CreateAccessToken(tokens.AccessToken) *errors.ResErr
	// }
	Service interface {
		GetById(string) (*tokens.AccessToken, *errors.ResErr)
		UpdateExpired(tokens.AccessToken) *errors.ResErr
		CreateAccessToken(tokens.AccessTokenRequest) (*tokens.AccessToken, *errors.ResErr)
	}
	service struct {
		restUsersRepo rest.RestUsersRepository
		dbRepo        db.DbRepository
	}
)

// NewService NewService
func NewService(repo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: repo,
		dbRepo:        dbRepo,
	}
}

// CreateAccessToken CreateAccessToken
// func (s *service) CreateAccessToken(at tokens.AccessToken) *errors.ResErr {
// 	if err := s.dbRepo.CreateAccessToken(at); err != nil {
// 		return err
// 	}
// 	return nil
// }
func (s *service) CreateAccessToken(request tokens.AccessTokenRequest) (*tokens.AccessToken, *errors.ResErr) {
	if err := request.Validation(); err != nil {
		return nil, err
	}
	user, errLogin := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if errLogin != nil {
		return nil, errLogin
	}
	at := tokens.GetAccessToken(user.Id)
	at.Generate()
	if err := s.dbRepo.CreateAccessToken(at); err != nil {
		return nil, err
	}
	return &at, nil
}

// GetById GetById
func (s *service) GetById(accessTokenId string) (*tokens.AccessToken, *errors.ResErr) {
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("id must not empty")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil

}

func (s *service) UpdateExpired(at tokens.AccessToken) *errors.ResErr {
	return s.dbRepo.UpdateExpired(at)
}
