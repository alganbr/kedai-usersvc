package repos

import (
	"context"
	"github.com/alganbr/kedai-usersvc/internal/databases"
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type IPasswordRepository interface {
	Get(int64) (*models.UserPassword, *errors.Error)
	Create(*models.UserPassword) *errors.Error
	Update(*models.UserPassword) *errors.Error
}

type PasswordRepository struct {
	db *databases.DB
}

func NewPasswordRepository(db *databases.DB) IPasswordRepository {
	return &PasswordRepository{
		db: db,
	}
}

func (repo *PasswordRepository) Get(id int64) (*models.UserPassword, *errors.Error) {
	var userPassword models.UserPassword
	err := repo.db.Pool.QueryRow(context.Background(), getPasswordQuery, id).Scan(
		&userPassword.Id,
		&userPassword.Password,
		&userPassword.CreatedAt,
		&userPassword.CreatedBy,
		&userPassword.UpdatedAt,
		&userPassword.UpdatedBy,
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

	return &userPassword, nil
}

func (repo *PasswordRepository) Create(userPassword *models.UserPassword) *errors.Error {
	_, err := repo.db.Pool.Exec(context.Background(), createPasswordQuery,
		userPassword.Id,
		userPassword.Password,
		userPassword.CreatedAt,
		userPassword.CreatedBy,
		userPassword.UpdatedAt,
		userPassword.UpdatedBy,
	)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (repo *PasswordRepository) Update(userPassword *models.UserPassword) *errors.Error {
	_, err := repo.db.Pool.Exec(context.Background(), updatePasswordQuery,
		userPassword.Password,
		userPassword.UpdatedAt,
		userPassword.UpdatedBy,
		userPassword.Id,
	)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

const (
	getPasswordQuery = `
		SELECT 
			id,
			password,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM user_password
		WHERE id = $1
	`

	createPasswordQuery = `
		INSERT INTO user_password (
			id,
			password,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`

	updatePasswordQuery = `
		UPDATE user_password
		SET
			password = $1,
			updated_at = $2,
			updated_by = $3
		WHERE id = $4
	`
)
