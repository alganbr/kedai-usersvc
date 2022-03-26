package managers

import (
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-usersvc/internal/repos"
	"github.com/alganbr/kedai-utils/datetime"
	"github.com/alganbr/kedai-utils/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type IPasswordManager interface {
	Create(*models.UserPasswordRq) *errors.Error
	Update(*models.UserPasswordRq) *errors.Error
	Validate(*models.ValidatePasswordRq) (*models.User, *errors.Error)
}

type PasswordManager struct {
	passwordRepo repos.IPasswordRepository
	userRepo     repos.IUserRepository
}

func NewPasswordManager(passwordRepo repos.IPasswordRepository, userRepo repos.IUserRepository) IPasswordManager {
	return &PasswordManager{
		passwordRepo: passwordRepo,
		userRepo:     userRepo,
	}
}

func (mgr *PasswordManager) Create(rq *models.UserPasswordRq) *errors.Error {
	if validateErr := rq.Validate(); validateErr != nil {
		return validateErr
	}

	userPassword := &models.UserPassword{
		Id:        rq.Id,
		Password:  hashPassword(rq.Password),
		CreatedAt: datetime.GetUtcNow(),
		CreatedBy: rq.RequestedBy,
		UpdatedAt: datetime.GetUtcNow(),
		UpdatedBy: rq.RequestedBy,
	}

	err := mgr.passwordRepo.Create(userPassword)
	if err != nil {
		return err
	}

	return nil
}

func (mgr *PasswordManager) Update(rq *models.UserPasswordRq) *errors.Error {
	if validateErr := rq.Validate(); validateErr != nil {
		return validateErr
	}

	userPassword := &models.UserPassword{
		Id:        rq.Id,
		Password:  hashPassword(rq.Password),
		CreatedAt: datetime.GetUtcNow(),
		CreatedBy: rq.RequestedBy,
		UpdatedAt: datetime.GetUtcNow(),
		UpdatedBy: rq.RequestedBy,
	}

	err := mgr.passwordRepo.Update(userPassword)
	if err != nil {
		return err
	}

	return nil
}

func (mgr *PasswordManager) Validate(rq *models.ValidatePasswordRq) (*models.User, *errors.Error) {
	user, err := mgr.userRepo.GetByEmail(rq.Email)
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid user and password",
		}
	}
	userPassword, err := mgr.passwordRepo.Get(user.Id)
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid user and password",
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword.Password), []byte(rq.Password)); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid user and password",
		}
	}
	return user, nil
}

func hashPassword(password string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPasswordBytes)
}
