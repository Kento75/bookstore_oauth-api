package db

import (
	"github.com/Kento75/bookstore_oauth-api/src/clients/cassandra"
	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token"
	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token/utils/errors"
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

	// TODO:
	return nil, errors.InternalServerError("database connection not implemented yet")
}
