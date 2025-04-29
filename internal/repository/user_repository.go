package repository

import (
	"context"
	"errors"
	"fmt"
	"strongbody-api/internal/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		BeginTx(ctx context.Context) *gorm.DB
		Create(ctx context.Context, user *entity.User) (*entity.User, error)
		GetByID(ctx context.Context, id uint) (*entity.User, error)
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		Update(ctx context.Context, user *entity.User) (*entity.User, error)
		Delete(ctx context.Context, id uint) error
		AssignRole(ctx context.Context, tx *gorm.DB, userID uint, roleID int) error
		AssignApps(ctx context.Context, tx *gorm.DB, userID uint, appIDs []int) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) BeginTx(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) AssignRole(ctx context.Context, tx *gorm.DB, userID uint, roleID int) error {
	// Kiểm tra role có tồn tại không
	var role entity.Role
	if err := tx.WithContext(ctx).First(&role, roleID).Error; err != nil {
		return fmt.Errorf("role not found: %v", err)
	}

	// Gán role cho user
	userRole := entity.UserRoles{
		UserID: userID,
		RoleID: uint(roleID),
	}

	if err := tx.WithContext(ctx).Create(&userRole).Error; err != nil {
		return fmt.Errorf("failed to assign role to user: %v", err)
	}

	return nil
}

func (r *userRepository) AssignApps(ctx context.Context, tx *gorm.DB, userID uint, appIDs []int) error {
	// Kiểm tra apps có tồn tại không
	var count int64
	if err := tx.WithContext(ctx).Model(&entity.App{}).Where("id IN ?", appIDs).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to verify apps: %v", err)
	}
	if int(count) != len(appIDs) {
		return errors.New("some apps not found")
	}

	// Tạo records user_app
	var userApps []entity.UserApps
	for _, appID := range appIDs {
		userApps = append(userApps, entity.UserApps{
			UserID: userID,
			AppID:  uint(appID),
		})
	}

	if err := tx.WithContext(ctx).Create(&userApps).Error; err != nil {
		return fmt.Errorf("failed to assign apps to user: %v", err)
	}

	return nil
}
