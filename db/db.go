package db

import (
	"fmt"
	"github.com/micahli/notable-take-home/model"
	"github.com/micahli/notable-take-home/utils"
)

// DB is the where all the data stores
type DB struct {
	allDoctors      []model.Doctor
	allAppointments []model.Appointment
}

var dbInstance DB

func init() {
	// init the data here
	dbInstance.allDoctors = []model.Doctor{
		model.Doctor{
			UID:       "001",
			FirstName: "Hibbert",
			LastName:  "Julius",
		},
		model.Doctor{
			UID:       "002",
			FirstName: "Algernop",
			LastName:  "Krieger",
		},
		model.Doctor{
			UID:       "003",
			FirstName: "Riviera",
			LastName:  "Nick",
		},
	}

	// dbInstance.allAppointments = []model.Appointment {
	// 	model.Appointment{
	// 		UID: "",

	// 	}
	// }
}

func NewDB() *DB {
	return &dbInstance
}

func (db *DB) GetAllDoctors() ([]model.Doctor, error) {
	return db.allDoctors, nil
}

func (db *DB) GetDoctorAppointments(uid string) ([]model.Appointment, error) {
	var rst []model.Appointment
	for _, val := range db.allAppointments {
		if val.DoctorUID == uid {
			rst = append(rst, val)
		}
	}

	return rst, nil
}

func (db *DB) AddApointment(ap model.Appointment) (model.Appointment, error) {
	// get all the doctor's appointment with same time
	var sameTimeCnt int
	for _, val := range db.allAppointments {
		if val.DoctorUID == ap.DoctorUID && val.DateTime == ap.DateTime {
			sameTimeCnt++
		}
	}

	if sameTimeCnt >= 3 {
		return ap, fmt.Errorf("More than three at the sametime")
	}

	ap.UID = utils.NewUUID()

	dbInstance.allAppointments = append(dbInstance.allAppointments, ap)

	return ap, nil
}

func (db *DB) CancelAppointment(ap model.Appointment) error {
	for idx, val := range db.allAppointments {
		if val.UID == ap.UID {
			db.allAppointments = append(db.allAppointments[:idx], db.allAppointments[idx+1:]...)
			break
		}
	}

	return nil
}
