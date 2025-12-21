package models

import "time"

type Role struct {
	ID        string    `gorm:"type:varchar(50);primaryKey" json:"role_id,omitempty"`
	Name      string    `gorm:"unique" json:"name,omitempty" valid:"type(string)"`
	CreatedAt time.Time `json:"created_at"  valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Roles []Role

func (Role) TableName() string {
	return "roles"
}

type AddRoleRequest struct {
	Name string `json:"name" validate:"required"`
}
