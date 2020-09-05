package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}

	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestAccessToken(t *testing.T) {
	at := GetNoewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token token should not be nil")

	assert.EqualValues(t, "", at.AccessToken, "brand new access token token should not have defined access token id")

	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	// パターン1：期限が存在しない場合
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	// パターン1：期限が未来の場合
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}
