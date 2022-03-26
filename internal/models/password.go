package models

import (
	"github.com/alganbr/kedai-utils/errors"
	"net/http"
	"strings"
	"time"
)

type UserPassword struct {
	Id        int64     `json:"id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type UserPasswordRq struct {
	Id          int64  `json:"id"`
	Password    string `json:"password"`
	RequestedBy string `json:"requested_by"`
}

type ValidatePasswordRq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rq *UserPasswordRq) Validate() *errors.Error {
	rq.Password = strings.TrimSpace(strings.ToLower(rq.Password))
	if len(rq.Password) < 8 {
		return &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Password length must be at least 8 characters long",
		}
	}
	if rq.RequestedBy == "" {
		return &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Requested by cannot be empty",
		}
	}
	return nil
}
