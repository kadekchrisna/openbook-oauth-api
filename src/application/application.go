package app

import (
	"github.com/gin-gonic/gin"
	http "github.com/kadekchrisna/openbook-oauth/src/http"
	"github.com/kadekchrisna/openbook-oauth/src/repository/db"
	"github.com/kadekchrisna/openbook-oauth/src/repository/rest"
	tokens "github.com/kadekchrisna/openbook-oauth/src/services/access_token"
)

var (
	router = gin.Default()
)

// StartApplication StartApplication
func StartApplication() {
	// session, err := cassandra.GetSession()
	// if err != nil {
	// 	panic(err)
	// }
	// session.Close()

	serivces := http.NewServices(tokens.NewService(rest.NewRestUsersRepository(), db.New()))

	router.GET("/access-token/gen/:access_token_id", serivces.GetById)
	router.POST("/access-token/gen", serivces.CreateAccessToken)

	router.Run(":8080")
}
