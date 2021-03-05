package access_token

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestExpirationConst(t *testing.T) {

	assert.EqualValues(t,24,expirationTime,"expiration should be 24 hour" )
}

func TestGetNewAccessToken(t *testing.T) {
	at :=GetNewAccessToken()

	assert.False(t,at.IsExpired(),"new token shouldn't expired")
	assert.EqualValues(t, "",at.AccessToken,"token shouldn't empty")
	assert.True(t, at.UserId==0,"new token user id should be 0")

}

func TestAccessToken_IsExpired(t *testing.T) {
	at :=AccessToken{}
	assert.True(t, at.IsExpired(),"empty token should by default expired ")

	at.Expires=time.Now().UTC().Add(3*time.Hour).Unix()
	assert.False(t, at.IsExpired(),"token expiring in 3 hour ")
}