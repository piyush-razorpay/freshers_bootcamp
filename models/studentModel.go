package models

type Student struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment""`
	FirstName string `json:"firstName" gorm:"type:varchar(50)"`
	LastName  string `json:"lastName" gorm:"type:varchar(50)"`
	DOB       string `json:"dob" gorm:"type:varchar(50)"`
	Address   string `json:"address" gorm:"type:varchar(100)"`
	Subject   string `json:"subject" gorm:"type:varchar(50)"`
	Marks     uint   `json:"marks""`
}
