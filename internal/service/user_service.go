package service

import (
	"log"
	"net/http"
	"time"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/helper"
	"github.com/CRobinDev/Gemastik/pkg/supabase"
	"github.com/CRobinDev/Gemastik/pkg/validator"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(req model.UserRegisterRequest) (model.ServiceResponse, error)
	LoginUser(req model.UserLoginRequest) (model.ServiceResponse, error)
	UpdatePassword(req model.UpdatePasswordRequest) (model.ServiceResponse, error)
	UpdateProfile(req model.UserUpdateRequest) (model.ServiceResponse, error)
	UploadPhoto(req model.UploadPhotoRequest) (model.ServiceResponse, error)
	GetUser(req model.UserParam) (model.ServiceResponse, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Register(req model.UserRegisterRequest) (model.ServiceResponse, error) {
	if err := ValidateRequestRegister(req); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}, err
	}

	_, err := us.UserRepository.FindByEmail(req.Email)
	if err == nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrEmailAlreadyUsed.Error(),
			Data:    nil,
		}, err
	}

	hashPass, err := helper.HashPassword(req.Password)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrHashPassword.Error(),
			Data:    err.Error(),
		}, err
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	user := entity.User{
		ID:          uuid.New(),
		Email:       req.Email,
		Password:    hashPass,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}

	if err := us.UserRepository.Create(user); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrFailedCreateUser.Error(),
			Data:    err.Error(),
		}, err
	}

	resp := model.UserRegisterResponse{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully register user",
		Data:    resp,
	}, nil
}

func (us *UserService) LoginUser(req model.UserLoginRequest) (model.ServiceResponse, error) {
	user, err := us.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrEmailNotFound.Error(),
			Data:    nil,
		}, err
	}

	if err := helper.ComparePassword(user.Password, req.Password); err != nil {
		log.Println(err)
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrInvalidPassword.Error(),
			Data:    err.Error(),
		}, err
	}

	token, err := helper.SignJWTToken(user.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrGenerateToken.Error(),
			Data:    err.Error(),
		}, err
	}

	resp := model.UserLoginResponse{
		User: model.UserRegisterResponse{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		},
		Token: token,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully login",
		Data:    resp,
	}, nil
}

func ValidateRequestUpdatePassword(req model.UpdatePasswordRequest) error {
	switch {
	case req.OldPassword == "" || !validator.ValidatePassword(req.OldPassword):
		return errors.ErrInvalidPassword
	case req.NewPassword == "" || !validator.ValidatePassword(req.NewPassword):
		return errors.ErrInvalidPassword
	}
	return nil
}

func ValidateRequestRegister(req model.UserRegisterRequest) error {
	switch {
	case req.Name == "":
		return errors.ErrNameRequired
	case req.Email == "" || !validator.ValidateEmail(req.Email):
		return errors.ErrInvalidEmail
	case req.Password == "" || !validator.ValidatePassword(req.Password):
		return errors.ErrInvalidPassword
	case req.Name == "":
		return errors.ErrUsernameRequired
	case req.PhoneNumber == "" || !validator.ValidatePhone(req.PhoneNumber):
		return errors.ErrInvalidPhoneNumber
	}
	return nil
}

func (us *UserService) GetUser(req model.UserParam) (model.ServiceResponse, error) {
	user, err := us.UserRepository.FindByParam(req)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrEmailNotFound.Error(),
			Data:    nil,
		}, err
	}
	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully get user",
		Data:    user,
	}, nil
}

func (us *UserService) UpdatePassword(req model.UpdatePasswordRequest) (model.ServiceResponse, error) {
	if err := ValidateRequestUpdatePassword(req); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrInvalidPassword.Error(),
			Data:    nil,
		}, err
	}

	user, err := us.UserRepository.FindByID(req.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrEmailNotFound.Error(),
			Data:    nil,
		}, err
	}

	hashPass, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrHashPassword.Error(),
			Data:    err.Error(),
		}, err
	}

	if hashPass == user.Password {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrInvalidPassword.Error(),
			Data:    nil,
		}, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	user.Password = hashPass
	user.UpdatedAt = updatedAt
	if err := us.UserRepository.UpdatePassword(user); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrFailedUpdatePassword.Error(),
			Data:    err.Error(),
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully update password",
		Data:    nil,
	}, nil
}

func (us *UserService) UpdateProfile(req model.UserUpdateRequest) (model.ServiceResponse, error) {
	user, err := us.UserRepository.FindByID(req.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrUserNotFound.Error(),
			Data:    nil,
		}, err
	}

	switch {
	case req.Name == "":
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrNameRequired.Error(),
			Data:    nil,
		}, err
	case req.PhoneNumber == "":
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrPhoneNumberRequired.Error(),
			Data:    nil,
		}, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	log.Println("woi", updatedAt)

	user.Name = req.Name
	user.PhoneNumber = req.PhoneNumber
	user.UpdatedAt = updatedAt

	if err := us.UserRepository.UpdateProfile(user); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrFailedUpdatePassword.Error(),
			Data:    err.Error(),
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully update profile",
		Data:    nil,
	}, nil
}

func (us *UserService) UploadPhoto(req model.UploadPhotoRequest) (model.ServiceResponse, error) {
	user, err := us.UserRepository.FindByID(req.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrEmailNotFound.Error(),
			Data:    nil,
		}, err
	}
	supabaseStorage := supabase.NewSupabaseStorage()
	if user.PhotoLink != "" {
		if err := supabaseStorage.Delete(user.PhotoLink); err != nil {
			return model.ServiceResponse{
				Code:    http.StatusInternalServerError,
				Error:   true,
				Message: errors.ErrInternalServer.Error(),
				Data:    nil,
			}, err
		}
	}

	url, err := supabaseStorage.Upload(req.Photo)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrFailedUploadPhoto.Error(),
			Data:    err.Error(),
		}, err
	}

	log.Println("woi", url)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	err = us.UserRepository.UpdatePhoto(entity.User{
		ID:        user.ID,
		PhotoLink: url,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrFailedUploadPhoto.Error(),
			Data:    nil,
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully upload photo",
		Data:    url,
	}, nil
}
