package postgres

import (
	"time"

	"github.com/quannv/strongbody-api/internal/entity"
	utils "github.com/quannv/strongbody-api/utility"
)

// Seed inserts initial data for testing
func (d *Database) Seed() error {
	db := d.Connection

	// Avoid duplicate seed
	var count int64
	db.Model(&entity.App{}).Where("name = ?", "network").Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	// Create App
	app := &entity.App{
		Name:   "network",
		Key:    "network-key",
		Config: `{"key": "value"}`,
	}
	if err := db.Create(app).Error; err != nil {
		return err
	}

	app1 := &entity.App{
		Name:   "shop",
		Key:    "network-key",
		Config: `{"key": "value"}`,
	}
	if err := db.Create(app1).Error; err != nil {
		return err
	}

	// Create Role
	role := &entity.Role{
		Name:        "admin",
		Key:         "admin",
		Description: "Administrator role",
		AppID:       app.ID,
	}
	if err := db.Create(role).Error; err != nil {
		return err
	}

	role1 := &entity.Role{
		Name:        "bảo vệ",
		Key:         "admin",
		Description: "Administrator role",
		AppID:       app1.ID,
	}
	if err := db.Create(role1).Error; err != nil {
		return err
	}

	// hash password
	hashedPassword, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}

	hashedPassword1, err := utils.HashPassword("user")
	if err != nil {
		return err
	}

	// Create User
	user := &entity.User{
		Email:          "admin@example.com",
		PasswordHashed: hashedPassword,
		UserName:       "adminuser",
		FirstName:      "Admin",
		LastName:       "User",
		BirthDate:      ptrTime(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Language:       "en",
		Currency:       "USD",
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}

	user1 := &entity.User{
		Email:          "user@example.com",
		PasswordHashed: hashedPassword1,
		UserName:       "user",
		FirstName:      "User",
		LastName:       "User",
		BirthDate:      ptrTime(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Language:       "en",
		Currency:       "USD",
	}
	if err := db.Create(user1).Error; err != nil {
		return err
	}

	// Attach App to User (user_apps)
	if err := db.Model(&user).Association("Apps").Append(app); err != nil {
		return err
	}

	if err := db.Model(&user).Association("Apps").Append(app1); err != nil {
		return err
	}

	if err := db.Model(&user1).Association("Apps").Append(app); err != nil {
		return err
	}

	// Attach Role to User (user_roles)
	if err := db.Model(&user).Association("Roles").Append(role); err != nil {
		return err
	}

	if err := db.Model(&user).Association("Roles").Append(role1); err != nil {
		return err
	}

	if err := db.Model(&user1).Association("Roles").Append(role); err != nil {
		return err
	}

	return nil
}

// Helper to get pointer to time
func ptrTime(t time.Time) *time.Time {
	return &t
}
