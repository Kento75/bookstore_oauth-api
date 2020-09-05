package access_token

import (
	"fmt"
	"time"
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
