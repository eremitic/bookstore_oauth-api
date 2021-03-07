package db

import (
	"github.com/eremitic/bookstore_oauth-api/src/clients/cassandra"
	"github.com/eremitic/bookstore_oauth-api/src/domain/access_token"
	"github.com/eremitic/bookstore_oauth-api/src/domain/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {

	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}

	defer session.Close()

	return nil, errors.NewInternalErr("database no implemented")
}
