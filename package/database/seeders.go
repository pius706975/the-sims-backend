package database

import (
	"fmt"
	"log"

	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/database/seeders"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type seederData struct {
	name  string
	model interface{}
	size  int
}

var SeedCMD = &cobra.Command{
	Use:   "seed",
	Short: "For running db seeder",
	RunE:  Seed,
}

var seedUP bool
var seedDOWN bool

func init() {
	SeedCMD.Flags().BoolVarP(&seedUP, "seedUP", "u", true, "run seed up")

	SeedCMD.Flags().BoolVarP(&seedDOWN, "seedDOWN", "d", false, "run seed down")
}

func seedDown(db *gorm.DB) error {

	var err error

	var seedModel = []seederData{
		{
			name:  models.EmployeeType{}.TableName(),
			model: models.EmployeeType{},
		},

		{
			name:  models.EmploymentStatus{}.TableName(),
			model: models.EmploymentStatus{},
		},

		{
			name:  models.Employee{}.TableName(),
			model: models.Employee{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for ", data.name)
		sql := fmt.Sprintf("DELETE FROM %v ", data.name)
		err = db.Exec(sql).Error
	}

	return err
}

func seedUp(db *gorm.DB) error {

	var err error

	var seedModel = []seederData{
		{
			name:  "employee type",
			model: seeders.EmployeeTypeSeed,
			size:  cap(seeders.EmployeeTypeSeed),
		},

		{
			name:  "employment status",
			model: seeders.EmploymentStatusSeed,
			size:  cap(seeders.EmploymentStatusSeed),
		},

		{
			name:  "employee",
			model: seeders.EmployeeSeed,
			size:  cap(seeders.EmployeeSeed),
		},
	}

	for _, data := range seedModel {

		log.Println("create seeding data for ", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err
}

func Seed(cmd *cobra.Command, args []string) error {

	var err error

	db, err := NewDB()
	if err != nil {
		return err
	}

	if seedDOWN {
		err = seedDown(db)
		return err
	}

	if seedUP {
		err = seedUp(db)
	}

	return err
}
