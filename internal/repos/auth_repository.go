package repos

import (
	"context"
	"github.com/alganbr/kedai-authsvc/internal/databases"
	"github.com/alganbr/kedai-authsvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type IAuthRepository interface {
	Get(string) (*models.AccessToken, *errors.Error)
	Create(*models.AccessToken) *errors.Error
}

type AuthRepository struct {
	db *databases.DB
}

func NewAuthRepository(db *databases.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) Get(id string) (*models.AccessToken, *errors.Error) {
	var accessToken models.AccessToken
	err := repo.db.Pool.QueryRow(context.Background(), getAccessTokenQuery, id).Scan(
		&accessToken.Id,
		&accessToken.UserId,
		&accessToken.Expires,
	)

	if err == pgx.ErrNoRows {
		return nil, &errors.Error{
			Code:    http.StatusNoContent,
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &accessToken, nil
}

func (repo *AuthRepository) Create(accessToken *models.AccessToken) *errors.Error {
	err := repo.db.Pool.QueryRow(context.Background(), createAccessTokenQuery,
		accessToken.UserId,
		accessToken.Expires,
	).Scan(&accessToken.Id)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

const (
	getAccessTokenQuery = `
		SELECT 
			id,
			user_id,
			expires
		FROM access_token 
		WHERE id = $1;
	`

	createAccessTokenQuery = `
		INSERT INTO access_token (
			user_id,
			expires
		) VALUES (
			$1,
			$2
		) RETURNING id;
	`
)
