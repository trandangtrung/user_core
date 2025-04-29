package postgres

import (
	"strongbody-api/internal/entity"
	utils "strongbody-api/utility"
)

func (d *Database) Seed() error {
	db := d.Connection

	// 1. Tạo User
	hashedPassword, err := utils.HashPassword("hashed-password")
	if err != nil {
		return err
	}
	user := entity.User{
		Email:          "test@example.com",
		PasswordHashed: hashedPassword,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// 2. Tạo Platform
	app := entity.App{
		Name:   "network",
		Config: `{"theme":"dark","version":"1.0.0"}`,
	}
	if err := db.Create(&app).Error; err != nil {
		return err
	}

	// 3. Tạo Role
	role := entity.Role{
		AppID:       app.ID,
		Name:        "Admin",
		Description: "Quyền quản trị",
	}
	if err := db.Create(&role).Error; err != nil {
		return err
	}

	// 4. Gán Role cho User
	userRole := entity.UserRoles{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := db.Create(&userRole).Error; err != nil {
		return err
	}

	// 5. Gán Platform cho User
	userPlatform := entity.UserApps{
		UserID: user.ID,
		AppID:  uint(app.ID),
	}
	if err := db.Create(&userPlatform).Error; err != nil {
		return err
	}

	// 6. Tạo Refresh Token
	token := entity.Token{
		UserID:       user.ID,
		RefreshToken: "sample-refresh-token-123456",
		Scope:        "read write",
	}
	if err := db.Create(&token).Error; err != nil {
		return err
	}

	return nil
}
