package repository

import (
	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	FindByID(id uuid.UUID) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	UpdateProfile(user entity.User) error
	UpdatePassword(user entity.User) error
	UpdatePhoto(user entity.User) error
	Delete(id uuid.UUID) error
	FindByParam(param model.UserParam) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user entity.User) error {
	if err := r.db.Create(map[string]interface{}{
		"phone_number": user.PhoneNumber,
		"name":         user.Name,
		"email":        user.Email,
		"password":     user.Password,
		"photo_link":   user.PhotoLink,
		"updated_at":   user.CreatedAt,
		"created_at":   user.CreatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) UpdateProfile(user entity.User) error {
	if err := r.db.Model(&user).Where("id", user.ID).Updates(map[string]interface{}{
		"name":         user.Name,
		"phone_number": user.PhoneNumber,
		"updated_at":   user.UpdatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdatePassword(user entity.User) error {
	if err := r.db.Model(&user).Updates(map[string]interface{}{
		"password":   user.Password,
		"updated_at": user.UpdatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdatePhoto(user entity.User) error {
	if err := r.db.Model(&user).Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"photo_link": user.PhotoLink,
			"updated_at": user.UpdatedAt,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByParam(param model.UserParam) (entity.User, error) {
	var user entity.User

	if err := r.db.Where(&param).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
