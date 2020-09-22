package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token/utils/errors"
)

// REST Provider

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	// 認証タイプ パスワードの場合
	case grantTypePassword:
		break
		// 認証タイプ クレデンシャルの場合
	case grantTypeClientCredentials:
		break
	default:
		return errors.BadRequestError("invalid grant_type parameter")
	}

	// TODO: validate parameters for each grant_type
	return nil
}

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

func GetNewAccessToken() AccessToken {
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
