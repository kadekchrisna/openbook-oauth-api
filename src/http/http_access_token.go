package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	tokens "github.com/kadekchrisna/openbook-oauth-api/src/domain/access_token"
	services "github.com/kadekchrisna/openbook-oauth-api/src/services/access_token"
	"github.com/kadekchrisna/openbook-oauth-api/src/utils/errors"
)

type (
	AccessTokenHandler interface {
		GetById(*gin.Context)
		CreateAccessToken(c *gin.Context)
	}
	accessTokenHandler struct {
		Service services.Service
	}
)

// NewServices NewServices
func NewServices(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		Service: service,
	}
}

// CreateAccessToken CreateAccessToken
func (s *accessTokenHandler) CreateAccessToken(c *gin.Context) {
	var at tokens.AccessTokenRequest
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&at); err != nil {
		resErr := errors.NewBadRequestError("Invalid json object")
		c.JSON(resErr.Status, resErr)
		return
	}
	fmt.Println(at)
	fmt.Println(at)
	token, errToken := s.Service.CreateAccessToken(at)
	if errToken != nil {
		c.JSON(errToken.Status, errToken)
		return
	}
	c.JSON(http.StatusOK, token)
	return

}

// GetById GetById
func (s *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := s.Service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
	return

}

//
// func (s *accessTokenHandler) UpdateExpired(c *gin.Context) {
// 	accessToken, err := s.Service.GetById(strings.TrimSpace(c.Param("access_token_id")))
// 	if err != nil {
// 		c.JSON(err.Status, err)
// 		return
// 	}
// 	if err := s.Service.UpdateExpired() {

// 	}
// }
