package repository

import (
	"context"
	"errors"

	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateApp(ctx context.Context, app *entity.App) (*entity.App, error)
		GetAppByID(ctx context.Context, id uint) (*entity.App, error)
		UpdateApp(ctx context.Context, app *entity.App) (*entity.App, error)
		DeleteApp(ctx context.Context, id uint) error

		CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error)
		GetRoleByID(ctx context.Context, id uint) (*entity.Role, error)
		GetRolesByUserIDAndAppName(ctx context.Context, userID uint, appName string) ([]*entity.Role, error)
		UpdateRole(ctx context.Context, role *entity.Role) (*entity.Role, error)
		DeleteRole(ctx context.Context, id uint) error

		CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
		GetTokenByID(ctx context.Context, id uint) (*entity.Token, error)
		UpdateToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
		DeleteToken(ctx context.Context, id uint) error

		CreateUserApp(ctx context.Context, user *entity.User, app *entity.App) error
		GetAppsByUserID(ctx context.Context, userID uint) ([]*entity.App, error)

		CreateUserRole(ctx context.Context, user *entity.User, role *entity.Role) error
		GetRolesByUserID(ctx context.Context, userID uint) ([]*entity.Role, error)

		CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
		GetUserByID(ctx context.Context, id uint) (*entity.User, error)
		GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
		DeleteUser(ctx context.Context, id uint) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateApp(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := r.db.WithContext(ctx).Create(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (r *userRepository) GetAppByID(ctx context.Context, id uint) (*entity.App, error) {
	var app entity.App
	if err := r.db.WithContext(ctx).First(&app, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &app, nil
}

func (r *userRepository) UpdateApp(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := r.db.WithContext(ctx).Save(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (r *userRepository) DeleteApp(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.App{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *userRepository) GetRoleByID(ctx context.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *userRepository) GetRolesByUserIDAndAppName(ctx context.Context, userID uint, appName string) ([]*entity.Role, error) {
	var roles []*entity.Role

	err := r.db.WithContext(ctx).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Joins("JOIN apps ON apps.id = roles.app_id").
		Where("user_roles.user_id = ? AND apps.name = ?", userID, appName).
		Find(&roles).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return roles, nil
}

func (r *userRepository) UpdateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *userRepository) DeleteRole(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	if err := r.db.WithContext(ctx).Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *userRepository) GetTokenByID(ctx context.Context, id uint) (*entity.Token, error) {
	var token entity.Token
	if err := r.db.WithContext(ctx).First(&token, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}

func (r *userRepository) UpdateToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	if err := r.db.WithContext(ctx).Save(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *userRepository) DeleteToken(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Token{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateUserApp(ctx context.Context, user *entity.User, app *entity.App) error {
	if err := r.db.WithContext(ctx).Model(user).Association("Apps").Append(app); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAppsByUserID(ctx context.Context, userID uint) ([]*entity.App, error) {
	var apps []*entity.App
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Association("Apps").Find(&apps); err != nil {
		return nil, err
	}
	return apps, nil
}

func (r *userRepository) CreateUserRole(ctx context.Context, user *entity.User, role *entity.Role) error {
	if err := r.db.WithContext(ctx).Model(user).Association("Roles").Append(role); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetRolesByUserID(ctx context.Context, userID uint) ([]*entity.Role, error) {
	var roles []*entity.Role
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Association("Roles").Find(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
