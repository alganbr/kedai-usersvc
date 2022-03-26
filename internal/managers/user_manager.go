package managers

import (
	"github.com/alganbr/kedai-usersvc/internal/models"
	"github.com/alganbr/kedai-usersvc/internal/repos"
	"github.com/alganbr/kedai-utils/datetime"
	"github.com/alganbr/kedai-utils/errors"
	"net/http"
)

type IUserManager interface {
	Get(int64) (*models.User, *errors.Error)
	Create(*models.UserRq) (*models.User, *errors.Error)
	Update(int64, *models.UserRq) (*models.User, *errors.Error)
	Patch(int64, *models.UserRq) (*models.User, *errors.Error)
}

type UserManager struct {
	userRepo        repos.IUserRepository
	passwordManager IPasswordManager
}

func NewUserManager(userRepo repos.IUserRepository, passwordManager IPasswordManager) IUserManager {
	return &UserManager{
		userRepo:        userRepo,
		passwordManager: passwordManager,
	}
}

func (mgr *UserManager) Get(id int64) (*models.User, *errors.Error) {
	user, err := mgr.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mgr *UserManager) Create(rq *models.UserRq) (*models.User, *errors.Error) {
	if validateErr := rq.Validate(); validateErr != nil {
		return nil, validateErr
	}

	user := &models.User{
		Email:     rq.Email,
		FirstName: rq.FirstName,
		LastName:  rq.LastName,
		CreatedAt: datetime.GetUtcNow(),
		CreatedBy: rq.RequestedBy,
		UpdatedAt: datetime.GetUtcNow(),
		UpdatedBy: rq.RequestedBy,
	}

	err := mgr.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	err = mgr.passwordManager.Create(&models.UserPasswordRq{
		Id:          user.Id,
		Password:    rq.Password,
		RequestedBy: rq.RequestedBy,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mgr *UserManager) Update(id int64, rq *models.UserRq) (*models.User, *errors.Error) {
	if validateErr := rq.Validate(); validateErr != nil {
		return nil, validateErr
	}

	user, err := mgr.Get(id)
	if err != nil {
		return nil, err
	}

	updated := false
	if rq.Email != user.Email {
		user.Email = rq.Email
		updated = true
	}
	if rq.FirstName != user.FirstName {
		user.FirstName = rq.FirstName
		updated = true
	}
	if rq.LastName != user.LastName {
		user.LastName = rq.LastName
		updated = true
	}
	if !updated {
		return nil, &errors.Error{
			Code:    http.StatusNotModified,
			Message: "Not modified",
		}
	}

	user.UpdatedAt = datetime.GetUtcNow()
	user.UpdatedBy = rq.RequestedBy

	err = mgr.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mgr *UserManager) Patch(id int64, rq *models.UserRq) (*models.User, *errors.Error) {
	if validateErr := rq.Validate(); validateErr != nil {
		return nil, validateErr
	}

	user, err := mgr.Get(id)
	if err != nil {
		return nil, err
	}

	if rq.Email == "" {
		rq.Email = user.Email
	}
	if rq.FirstName == "" {
		rq.FirstName = user.FirstName
	}
	if rq.LastName == "" {
		rq.LastName = user.LastName
	}

	user, err = mgr.Update(id, rq)
	if err != nil {
		return nil, err
	}

	return user, nil
}
