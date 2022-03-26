package repos

import (
	"context"
	"github.com/alganbr/kedai-usersvc/internal/databases"
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type IUserRepository interface {
	Get(int64) (*models.User, *errors.Error)
	GetByEmail(string) (*models.User, *errors.Error)
	Create(*models.User) *errors.Error
	Update(*models.User) *errors.Error
}

type UserRepository struct {
	db *databases.DB
}

func NewUserRepository(db *databases.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Get(id int64) (*models.User, *errors.Error) {
	var user models.User
	err := repo.db.Pool.QueryRow(context.Background(), getUserQuery, id).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
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

	return &user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*models.User, *errors.Error) {
	var user models.User
	err := repo.db.Pool.QueryRow(context.Background(), getUserByEmailQuery, email).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
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

	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) *errors.Error {
	err := repo.db.Pool.QueryRow(context.Background(), createUserQuery,
		user.Email,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
		user.UpdatedBy,
	).Scan(&user.Id)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (repo *UserRepository) Update(user *models.User) *errors.Error {
	_, err := repo.db.Pool.Exec(context.Background(), updateUserQuery,
		user.Email,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.UpdatedBy,
		user.Id,
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
	getUserQuery = `
		SELECT 
			id,
			email,
			first_name,
			last_name,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM users
		WHERE id = $1
	`

	getUserByEmailQuery = `
		SELECT 
			id,
			email,
			first_name,
			last_name,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM users
		WHERE email = $1
	`

	createUserQuery = `
		INSERT INTO users (
			email,
			first_name,
			last_name,
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
			$6,
			$7
		) RETURNING id
	`

	updateUserQuery = `
		UPDATE users
		SET
			email = $1,
			first_name = $2,
			last_name = $3,
			updated_at = $4,
			updated_by = $5
		WHERE id = $6
	`
)
