package model

// Doctor info
type Doctor struct {
	UID       string `json:"uid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
