package structs

import "time"

type User struct{
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Email		string
	Nama 		string
	Nip 		string
	Jabatan 	string
	Instansi 	string
	Kontak 		string
	Password 	string
	Role 		string
	Token		string
}

func (User) User()string{
	return "user"
}