package repository

import (
	"context"
	"errors"

	v1 "github.com/quannv/strongbody-api/api/user/v1"
	"github.com/quannv/strongbody-api/internal/entity"
	utils "github.com/quannv/strongbody-api/utility"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
		GetTokenByID(ctx context.Context, id uint) (*entity.Token, error)
		UpdateToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
		DeleteToken(ctx context.Context, id uint) error

		CreateUserApp(ctx context.Context, user *entity.User, app *entity.App) error
		GetAppsByUserID(ctx context.Context, userID uint) ([]*entity.App, error)

		CreateUserRole(ctx context.Context, user *entity.User, role *entity.Role) error
		GetRolesByUserID(ctx context.Context, userID uint) ([]*entity.Role, error)

		CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error)
		CreateUserByAdmin(ctx context.Context, req *v1.CreateReq) (*entity.User, error)
		GetUserByID(ctx context.Context, id uint) (*entity.User, error)
		GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
		DeleteUser(ctx context.Context, id uint) error

		CreateEmailOTP(ctx context.Context, tx *gorm.DB, emailOtp *entity.EmailOTP) (*entity.EmailOTP, error)
		GetEmailOTPByEmail(ctx context.Context, tx *gorm.DB, email string) (*entity.EmailOTP, error)
		UpdateEmailOTP(ctx context.Context, tx *gorm.DB, emailOtp *entity.EmailOTP) (*entity.EmailOTP, error)
		IsOtpValid(ctx context.Context, tx *gorm.DB, email string, otp string) (bool, error)

		CreatePendingUser(ctx context.Context, tx *gorm.DB, pendingUser *entity.PendingUser) (*entity.PendingUser, error)
		GetPendingUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*entity.PendingUser, error)
		UpdatePendingUser(ctx context.Context, tx *gorm.DB, pendingUser *entity.PendingUser) (*entity.PendingUser, error)
		DeletePendingUser(ctx context.Context, tx *gorm.DB, id uint) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
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

func (r *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error) {
	if err := tx.WithContext(ctx).Create(user).Error; err != nil {
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

func (u *userRepository) CreateUserByAdmin(ctx context.Context, req *v1.CreateReq) (*entity.User, error) {
	var createdUser *entity.User

	err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Check if email exists
		var count int64
		if err := tx.Model(&entity.User{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("email is already exist")
		}

		// 2. Hash password
		passwordHashed, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}

		// 3. Create user
		user := &entity.User{
			UserName:       req.UserName,
			Email:          req.Email,
			PasswordHashed: passwordHashed,
			Mobile:         req.Mobile,
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 4. Get role by ID
		var role entity.Role
		if err := tx.First(&role, req.Role).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("role is not exist")
			}
			return err
		}

		// 5. Attach role to user
		if err := tx.Model(user).Association("Roles").Append(&role); err != nil {
			return err
		}

		// 6. Attach apps
		var apps []*entity.App
		for _, appID := range req.Apps {
			var app entity.App
			if err := tx.First(&app, appID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("app is not exist")
				}
				return err
			}
			apps = append(apps, &app)
		}
		if err := tx.Model(user).Association("Apps").Append(apps); err != nil {
			return err
		}

		createdUser = user
		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (r *userRepository) CreateEmailOTP(ctx context.Context, tx *gorm.DB, emailOtp *entity.EmailOTP) (*entity.EmailOTP, error) {
	if err := tx.WithContext(ctx).Create(emailOtp).Error; err != nil {
		return nil, err
	}
	return emailOtp, nil
}

func (r *userRepository) GetEmailOTPByEmail(ctx context.Context, tx *gorm.DB, email string) (*entity.EmailOTP, error) {
	var emailOtp entity.EmailOTP
	if err := tx.WithContext(ctx).Where("email = ?", email).First(&emailOtp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &emailOtp, nil
}

func (r *userRepository) UpdateEmailOTP(ctx context.Context, tx *gorm.DB, emailOtp *entity.EmailOTP) (*entity.EmailOTP, error) {
	if err := tx.WithContext(ctx).Where("id = ?", emailOtp.ID).Updates(emailOtp).Error; err != nil {
		return nil, err
	}
	return emailOtp, nil
}

// SELECT * FROM email_otps
// WHERE email = $1 AND otp = $2 AND is_used = FALSE AND expires_at > NOW()
// ORDER BY created_at DESC
// LIMIT 1;
func (r *userRepository) IsOtpValid(ctx context.Context, tx *gorm.DB, email string, otp string) (bool, error) {
	var emailOtp entity.EmailOTP
	if err := tx.WithContext(ctx).Where("email = ? AND otp = ? AND used = FALSE AND expire_at > NOW()", email, otp).Order("created_at DESC").First(&emailOtp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	// Mark the OTP as used
	emailOtp.Used = true
	if err := tx.WithContext(ctx).Save(&emailOtp).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepository) CreatePendingUser(ctx context.Context, tx *gorm.DB, pendingUser *entity.PendingUser) (*entity.PendingUser, error) {
	if err := tx.WithContext(ctx).Create(pendingUser).Error; err != nil {
		return nil, err
	}
	return pendingUser, nil
}

func (r *userRepository) GetPendingUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*entity.PendingUser, error) {
	var pendingUser entity.PendingUser
	if err := tx.WithContext(ctx).Where("email = ?", email).First(&pendingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &pendingUser, nil
}

func (r *userRepository) UpdatePendingUser(ctx context.Context, tx *gorm.DB, pendingUser *entity.PendingUser) (*entity.PendingUser, error) {
	if err := tx.WithContext(ctx).Where("id = ?", pendingUser.ID).Updates(pendingUser).Error; err != nil {
		return nil, err
	}
	return pendingUser, nil
}

func (r *userRepository) DeletePendingUser(ctx context.Context, tx *gorm.DB, id uint) error {
	if err := tx.WithContext(ctx).Delete(&entity.PendingUser{}, id).Error; err != nil {
		return err
	}
	return nil
}
