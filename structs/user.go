package structs

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
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