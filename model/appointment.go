package model

// patient type
type AppointmentKind uint8

const (
	NEWPATIENT AppointmentKind = 0 // new patient
	FOLLOWUP   AppointmentKind = 1 // follow-up
)

// Appointment info
type Appointment struct {
	UID              string          `json:"uid"`
	DoctorUID        string          `json:"doctoruid"`
	PatientFirstName string          `json:"patientfirstname"`
	PatientLastName  string          `json:"patientlastname"`
	DateTime         string          `json:"datetime"`
	Kind             AppointmentKind `json:"kind"`
}
