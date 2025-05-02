package entity

type Category struct {
	BaseEntity
	Name        string `json:"name" gorm:"type:varchar(100);not null;unique"`
	Description string `json:"description" gorm:"type:text"`
	Slug        string `json:"slug" gorm:"type:varchar(100);not null;unique"`
	IsActive    bool   `json:"is_active" gorm:"not null;default:true"`

	// Parent-child relationship for category hierarchy
	ParentID *uint      `json:"parent_id" gorm:"type:bigint"`
	Parent   *Category  `json:"parent" gorm:"foreignKey:ParentID"`
	Children []Category `json:"children" gorm:"foreignKey:ParentID"`

	// Products in this category
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
