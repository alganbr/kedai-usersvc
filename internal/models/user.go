package models

import (
	"github.com/alganbr/kedai-utils/errors"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type UserRq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RequestedBy string `json:"requested_by"`
}

func (rq *UserRq) Validate() *errors.Error {
	rq.Email = strings.TrimSpace(strings.ToLower(rq.Email))
	if rq.Email == "" {
		return &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid email address",
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
