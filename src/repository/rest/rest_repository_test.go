package rest

import (
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())

}

func TestLoginUserTimeout(t *testing.T) {

}
