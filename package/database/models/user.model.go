package models

import "time"

type User struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4(); column:user_id" json:"user_id,omitempty" valid:"-"`

	RoleID *string `gorm:"type:varchar(50);default:null" json:"role_id,omitempty"`
	Role   *Role   `gorm:"foreignKey:RoleID;references:ID"`

	Name        string    `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Name is required"`
	Username    string    `gorm:"not null" json:"username,omitempty" valid:"type(string)"`
	Email       string    `gorm:"not null;unique" json:"email" valid:"email, required~Email is required"`
	Password    string    `gorm:"not null" json:"password,omitempty" valid:"type(string), required~Password is required"`
	IsActivated bool      `gorm:"default: false" json:"is_activated,omitempty" valid:"-"`
	IsSuperUser bool      `gorm:"column:is_superuser;default: false" json:"is_superuser,omitempty" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	CreatedBy   string    `gorm:"default:null" json:"created_by,omitempty" valid:"type(string)"`
	UpdatedAt   time.Time `gorm:"default:null" json:"updated_at" valid:"-"`
	UpdatedBy   string    `gorm:"default:null" json:"updated_by,omitempty" valid:"type(string)"`
}

type Users []User

func (User) TableName() string {
	return "users"
}

type UserRegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateNewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
