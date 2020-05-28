package tokens

import (
	"fmt"
	"strings"
	"time"

	cryptoutils "github.com/kadekchrisna/openbook-oauth-api/src/utils/crypto"
	"github.com/kadekchrisna/openbook-oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type (
	AccessToken struct {
		AccessToken string `json:"access_token"`
		UserId      int64  `json:"user_id"`
		ClientId    int64  `json:"client_id"`
		Expires     int64  `json:"expires"`
	}
	AccessTokenRequest struct {
		GrantType string `json:"grant_type"`
		Scope     string `json:"scope"`

		Username string `json:"email"`
		Password string `json:"password"`

		// Used for client_credentials grant type
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
)

func (t *AccessTokenRequest) Validation() *errors.ResErr {
	switch t.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant type parameter")
	}
	return nil
}

// Validation Validation
func (t *AccessToken) Validation() *errors.ResErr {
	t.AccessToken = strings.TrimSpace(t.AccessToken)
	if t.AccessToken == "" {
		return errors.NewBadRequestError("Access token not valid")
	}
	if t.UserId == 0 {
		return errors.NewBadRequestError("User does not valid")
	}
	if t.ClientId == 0 {
		return errors.NewBadRequestError("Client does not valid")
	}
	return nil
}

// GetAccessToken as
func GetAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired asd
func (t *AccessToken) IsExpired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}

func (t *AccessToken) Generate() {
	t.AccessToken = cryptoutils.GetMD5(fmt.Sprintf("at-%d-%d-ran", t.UserId, t.Expires))
}
