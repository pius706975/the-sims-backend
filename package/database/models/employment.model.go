package models

import "time"

type Employee struct {
	ID                 string `gorm:"primarykey; type:uuid; default:uuid_generate_v4(); column:employee_id" json:"employee_id,omitempty" valid:"-"`
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
	IdentifyCardNumber int64  `gorm:"not null" json:"identify_card_number,omitempty" valid:"type(bigserial), required~IdentifyCardNumber is required"`

	EmployeeTypeID string       `gorm:"type:varchar(50);not null" json:"employee_type_id,omitempty"`
	EmployeeType   EmployeeType `gorm:"foreignKey:EmployeeTypeID;references:ID"`

	EmploymentStatusID string           `gorm:"type:varchar(50);not null" json:"employment_status_id,omitempty"`
	EmploymentStatus   EmploymentStatus `gorm:"foreignKey:EmploymentStatusID;references:ID"`

	JoinDate *time.Time `json:"join_date" valid:"-"`
	EndDate  *time.Time `json:"end_date" valid:"-"`

	IsActivated bool      `gorm:"default: true" json:"is_activated,omitempty" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	CreatedBy   string    `gorm:"default:null" json:"created_by,omitempty" valid:"type(string)"`
	UpdatedAt   time.Time `gorm:"default:null" json:"updated_at" valid:"-"`
	UpdatedBy   string    `gorm:"default:null" json:"updated_by,omitempty" valid:"type(string)"`
}

type Employees []Employee

func (Employee) TableName() string {
	return "employees"
}

type EmployeeRawResponse struct {
	EmployeeID         string    `json:"employee_id"`
	EmployeeNumber     string    `json:"employee_number"`
	FullName           string    `json:"full_name"`
	Gender             string    `json:"gender"`
	BirthPlace         string    `json:"birth_place"`
	BirthDate          time.Time `json:"birth_date"`
	Religion           string    `json:"religion"`
	MaritalStatus      string    `json:"marital_status"`
	Address            string    `json:"address"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	IdentifyCardNumber int64     `json:"identify_card_number"`

	EmployeeTypeID   string `json:"employee_type_id"`
	EmployeeTypeName string `json:"employee_type_name"`

	EmploymentStatusID   string `json:"employment_status_id"`
	EmploymentStatusName string `json:"employment_status_name"`

	JoinDate    *time.Time `json:"join_date"`
	EndDate     *time.Time `json:"end_date"`
	IsActivated bool       `json:"is_activated"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

// Employee Type
type EmployeeType struct {
	ID               string    `gorm:"type:varchar(50);primaryKey;column:employee_type_id" json:"employee_type_id,omitempty"`
	EmployeeTypeName string    `gorm:"not null" json:"employee_type_name,omitempty" valid:"type(string), required~EmployeeTypeName is required"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	CreatedBy        string    `gorm:"default:null" json:"created_by,omitempty" valid:"type(string)"`
	UpdatedAt        time.Time `gorm:"default:null" json:"updated_at" valid:"-"`
	UpdatedBy        string    `gorm:"default:null" json:"updated_by,omitempty" valid:"type(string)"`
}

type EmployeeTypes []EmployeeType

func (EmployeeType) TableName() string {
	return "employee_types"
}

// Employment Status
type EmploymentStatus struct {
	ID                   string    `gorm:"type:varchar(50);primaryKey;column:employment_status_id" json:"employment_status_id,omitempty"`
	EmploymentStatusName string    `gorm:"not null" json:"employment_status_name,omitempty" valid:"type(string), required~EmploymentStatusName is required"`
	CreatedAt            time.Time `json:"created_at" valid:"-"`
	CreatedBy            string    `gorm:"default:null" json:"created_by,omitempty" valid:"type(string)"`
	UpdatedAt            time.Time `gorm:"default:null" json:"updated_at" valid:"-"`
	UpdatedBy            string    `gorm:"default:null" json:"updated_by,omitempty" valid:"type(string)"`
}

type EmploymentStatuses []EmploymentStatus

func (EmploymentStatus) TableName() string {
	return "employment_statuses"
}

// Struct for swagger
type CreateEmployeeTypeRequest struct {
	ID               string `json:"employee_type_id"`
	EmployeeTypeName string `json:"employee_type_name"`
}

type CreateEmploymentStatusRequest struct {
	ID                   string `json:"employment_status_id"`
	EmploymentStatusName string `json:"employment_status_name"`
}

type CreateEmployeeRequest struct {
	EmployeeNumber     string `json:"employee_number" binding:"required"`
	FullName           string `json:"full_name" binding:"required"`
	Gender             string `json:"gender" binding:"required"`
	BirthPlace         string `json:"birth_place" binding:"required"`
	BirthDate          string `json:"birth_date" binding:"required"`
	Religion           string `json:"religion" binding:"required"`
	MaritalStatus      string `json:"marital_status" binding:"required"`
	Address            string `json:"address" binding:"required"`
	Phone              string `json:"phone" binding:"required"`
	Email              string `json:"email" binding:"required"`
	IdentifyCardNumber int64  `json:"identify_card_number" binding:"required"`

	EmployeeTypeID     string `json:"employee_type_id" binding:"required"`
	EmploymentStatusID string `json:"employment_status_id" binding:"required"`

	JoinDate string `json:"join_date" binding:"required"`
	EndDate  string `json:"end_date"`
}
