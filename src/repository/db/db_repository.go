package db

import (
	"fmt"

	"github.com/kadekchrisna/openbook-oauth/src/client/cassandra"
	tokens "github.com/kadekchrisna/openbook-oauth/src/domain/access_token"
	"github.com/kadekchrisna/openbook-oauth/src/utils/errors"
)

const (
	querySelectAccessTokenByAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken              = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpiredToken             = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type (
	DbRepository interface {
		GetById(string) (*tokens.AccessToken, *errors.ResErr)
		UpdateExpired(tokens.AccessToken) *errors.ResErr
		CreateAccessToken(tokens.AccessToken) *errors.ResErr
	}
	dbRepository struct{}
)

// New New
func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) CreateAccessToken(at tokens.AccessToken) *errors.ResErr {
	// if err := at.Validation(); err != nil {
	// 	return err
	// }
	session := cassandra.GetSession()
	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return nil
}

// GetById GetById
func (r *dbRepository) GetById(id string) (*tokens.AccessToken, *errors.ResErr) {
	var result tokens.AccessToken

	session := cassandra.GetSession()
	if err := session.Query(querySelectAccessTokenByAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err.Error() == "not found" {
			return nil, errors.NewNotFoundError(err.Error())
		}
		fmt.Println(err.Error())
		return nil, errors.NewInternalServerError(err.Error())
	}
	// defer session.Close()

	return &result, nil
}

// UpdateExpired UpdateExpired
func (r *dbRepository) UpdateExpired(at tokens.AccessToken) *errors.ResErr {
	if err := at.Validation(); err != nil {
		return err
	}
	session := cassandra.GetSession()
	if err := session.Query(queryUpdateExpiredToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return nil
}
