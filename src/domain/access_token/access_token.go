package access_token

import (
	"github.com/eremitic/bookstore_oauth-api/src/domain/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix()}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(now)
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadReqErr("invalid token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadReqErr("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadReqErr("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadReqErr("invalid expire")
	}
	return nil
}
