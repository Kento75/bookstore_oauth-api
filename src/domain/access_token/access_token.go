package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token/utils/errors"
)

// REST Provider

const (
	expirationTime = 24
)

/**
 * sample ↓
 * Web frontend - Client-Id: 123
 * Android App - Client-Id: 234
 */
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.BadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.BadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.BadRequestError("invalid expiration time")
	}

	return nil
}

func GetNoewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

/**
 * 有効期限チェック
 * Expiresが過去の場合 => true
 * 未来日の場合 => false
 */
func (accessToken AccessToken) IsExpired() bool {
	now := time.Now().UTC()

	expirationTime := time.Unix(accessToken.Expires, 0)
	fmt.Println(expirationTime)

	return expirationTime.Before(now)
}
