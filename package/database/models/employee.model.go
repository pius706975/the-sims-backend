package models

import "time"

type Employee struct {
	ID                 string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"employee_id,omitempty" valid:"-"`
	EmployeeNumber     string `gorm:"not null" json:"employee_number,omitempty" valid:"type(string), required~EmployeeNumber is required"`
	FullName           string `gorm:"not null" json:"full_name,omitempty" valid:"type(string), required~FullName is required"`
	Gender             string `gorm:"not null" json:"gender,omitempty" valid:"type(string), required~Gender is required"`
	BirthPlace         string `gorm:"not null" json:"birth_place,omitempty" valid:"type(string), required~BirthPlace is required"`
	BirthDate          string `gorm:"not null" json:"birth_date,omitempty" valid:"type(string), required~BirthDate is required"`
	Religion           string `gorm:"not null" json:"religion,omitempty" valid:"type(string), required~Religion is required"`
	MaritalStatus      string `gorm:"not null" json:"marital_status,omitempty" valid:"type(string), required~MaritalStatus is required"`
	Address            string `gorm:"not null" json:"address,omitempty" valid:"type(string), required~Address is required"`
	Phone              string `gorm:"not null" json:"phone,omitempty" valid:"type(string), required~Phone is required"`
	Email              string `gorm:"not null" json:"email,omitempty" valid:"type(string), required~Email is required"`
	IdentifyCardNumber int64 `gorm:"not null" json:"identify_card_number,omitempty" valid:"type(bigserial), required~IdentifyCardNumber is required"`

	JoinDate time.Time `json:"join_date" valid:"-"`
	EndDate  time.Time `json:"end_date" valid:"-"`

	IsActivated bool      `gorm:"default: true" json:"is_activated,omitempty" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	CreatedBy   string    `gorm:"null" json:"created_by,omitempty" valid:"type(string)"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
	UpdatedBy   string    `gorm:"null" json:"updated_by,omitempty" valid:"type(string)"`
}

type Employees []Employee

func (Employee) TableName() string {
	return "employees"
}

// type EmployeeType struct {
// 	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"employee_type_id,omitempty" valid:"-"`
// }
