package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegisterRequest struct {
	ID              uuid.UUID `json:"-"`
	PhoneNumber     string    `json:"phone_number"`
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required"`
	Password        string    `json:"password" binding:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserRegisterResponse struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	User  UserRegisterResponse `json:"user"`
	Token string               `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
}

type UserUpdateRequest struct {
	ID          uuid.UUID `json:"-"`
	PhoneNumber string    `json:"phone_number"`
	Name        string    `json:"name"`
	Email       string    `json:"-"`
}

type UploadPhotoRequest struct {
	ID        uuid.UUID             `json:"-"`
	PhotoLink string                `json:"-"`
	Photo     *multipart.FileHeader `form:"photo"`
}

type UpdatePasswordRequest struct {
	ID              uuid.UUID `json:"-"`
	OldPassword     string    `json:"old_password" binding:"required"`
	NewPassword     string    `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}
