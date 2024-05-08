package service

import (
	"log"
	"net/http"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/helper"
	"github.com/CRobinDev/Gemastik/pkg/validator"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(req model.UserRegisterRequest) (model.ServiceResponse, error)
	// RegisterAdmin(req model.UserRegisterRequest) (model.ServiceResponse, error)
	LoginUser(req model.UserLoginRequest) (model.ServiceResponse, error)
	// LoginAdmin(req model.UserLoginRequest) (model.ServiceResponse, error)
	// UploadPhoto(req model.UploadPhotoRequest) (model.ServiceResponse, error)
	// UpdatePassword(req model.UpdatePasswordRequest) (model.ServiceResponse, error)
	// UpdateProfile(req model.UserUpdateRequest) (model.ServiceResponse, error)
	GetUser(req model.UserParam) (entity.User, error)
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

	user := entity.User{
		ID:          uuid.New(),
		Email:       req.Email,
		Password:    hashPass,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
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
	log.Println(req.Password)
	log.Println(user.Password)

	if err := helper.ComparePassword(user.Password, req.Password); err != nil {
		log.Println(err)
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrInvalidPassword.Error(),
			Data:    nil,
		}, err
	}

	token, err := helper.CreateJWTToken(user.ID)
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

func (us *UserService) GetUser(req model.UserParam) (entity.User, error) {
	return us.UserRepository.FindByParam(req)
}
