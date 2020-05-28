package tokens

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpiationTime(t *testing.T) {
	// Line 11-14 are the same as the line 15
	// if expirationTime != 24 {
	// 	t.Error("Expiration time must be 24 hours")
	// }
	assert.EqualValues(t, 24, expirationTime, "Expiration time must be 24 hours")
}

func TestGetAccessToken(t *testing.T) {
	// Line 19-23 are the same as the line 24
	at := GetAccessToken(12)
	// if at.IsExpired() {
	// 	t.Error("Access token expired")
	// }
	assert.False(t, at.IsExpired(), "Access token expired")

	// if at.AccessToken != "" {
	// 	t.Error("Access token doesnt exist")
	// }
	assert.EqualValues(t, "", at.AccessToken, "Access token doesnt exist")
	// if at.UserId != 0 {
	// 	t.Error("User doesnt exist")
	// }
	assert.True(t, at.UserId == 0, "User doesnt exist")
}

func TestAccessTokenExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("Access token expired")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token expired")
	}

}
