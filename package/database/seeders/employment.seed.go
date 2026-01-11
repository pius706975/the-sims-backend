package seeders

import (
	"time"

	"github.com/pius706975/the-sims-backend/package/database/models"
)

var EmployeeTypeSeed = models.EmployeeTypes{
	{
		ID:               "GR",
		EmployeeTypeName: "Guru",
		CreatedBy:        "Super Admin Pius",
	},
	{
		ID:               "KS",
		EmployeeTypeName: "Kepala Sekolah",
		CreatedBy:        "Super Admin Pius",
	},
	{
		ID:               "WKS",
		EmployeeTypeName: "Wakil Kepala Sekolah",
		CreatedBy:        "Super Admin Pius",
	},
	{
		ID:               "TU",
		EmployeeTypeName: "Tata Usaha",
		CreatedBy:        "Super Admin Pius",
	},
	{
		ID:               "BD",
		EmployeeTypeName: "Bendahara",
		CreatedBy:        "Super Admin Pius",
	},
}

var EmploymentStatusSeed = models.EmploymentStatuses{
	{
		ID:                   "GTY",
		EmploymentStatusName: "Guru Tetap Yayasan",
		CreatedBy:            "Super Admin Pius",
	},
	{
		ID:                   "KT",
		EmploymentStatusName: "Kontrak",
		CreatedBy:            "Super Admin Pius",
	},
	{
		ID:                   "HN",
		EmploymentStatusName: "Honorer",
		CreatedBy:            "Super Admin Pius",
	},
}

func TimePtr(date string) *time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return &t
}

var EmployeeSeed = models.Employees{
	{
		EmployeeNumber:     "EMP001",
		FullName:           "Budi Santoso",
		Gender:             "Male",
		BirthPlace:         "Jakarta",
		BirthDate:          "1990-05-01",
		Religion:           "Islam",
		MaritalStatus:      "Single",
		Address:            "Jl. Merdeka No.1",
		Phone:              "08123456789",
		Email:              "budi@example.com",
		IdentifyCardNumber: 1234567890123456,
		JoinDate:           TimePtr("2024-01-15"),
		EndDate:            nil,
		IsActivated:        true,
		EmployeeTypeID:     "GR",
		EmploymentStatusID: "GTY",
		CreatedBy:          "Super Admin Pius",
	},

	{
		EmployeeNumber:     "EMP002",
		FullName:           "Prabowo Subianto",
		Gender:             "Male",
		BirthPlace:         "Jakarta",
		BirthDate:          "1958-08-12",
		Religion:           "Buddha",
		MaritalStatus:      "Menikah",
		Address:            "Jl. Belum Merdeka No.45",
		Phone:              "08123456789",
		Email:              "prasub@example.com",
		IdentifyCardNumber: 98765432123456,
		JoinDate:           TimePtr("2018-04-15"),
		EndDate:            nil,
		IsActivated:        true,
		EmployeeTypeID:     "GR",
		EmploymentStatusID: "HN",
		CreatedBy:          "Super Admin Pius",
	},

	{
		EmployeeNumber:     "EMP003",
		FullName:           "Purbaya Yudhi Sadewa",
		Gender:             "Male",
		BirthPlace:         "Bali",
		BirthDate:          "1964-07-07",
		Religion:           "Islam",
		MaritalStatus:      "Menikah",
		Address:            "Jl. Menuju Merdeka No.45",
		Phone:              "08123456789",
		Email:              "purbaya@example.com",
		IdentifyCardNumber: 91734983647138432,
		JoinDate:           TimePtr("2015-11-15"),
		EndDate:            nil,
		IsActivated:        true,
		EmployeeTypeID:     "BD",
		EmploymentStatusID: "KT",
		CreatedBy:          "Super Admin Pius",
	},
}

var PositionSeed = models.Positions{
	{
		ID: "GR",
		PositionName: "Guru",
		CreatedBy: "Super Admin Pius",
	},
	{
		ID: "PO",
		PositionName: "Pembina Osis",
		CreatedBy: "Super Admin Pius",
	},
	{
		ID: "TU",
		PositionName: "Tata Usaha",
		CreatedBy: "Super Admin Pius",
	},
	{
		ID: "BD",
		PositionName: "Bendahara",
		CreatedBy: "Super Admin Pius",
	},
	{
		ID: "WKS",
		PositionName: "Wakil Kepala Sekolah",
		CreatedBy: "Super Admin Pius",
	},
	{
		ID: "KS",
		PositionName: "Kepala Sekolah",
		CreatedBy: "Super Admin Pius",
	},
}