package access_token

import (
	"strings"

	"github.com/Kento75/bookstore_oauth-api/src/domain/access_token/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
