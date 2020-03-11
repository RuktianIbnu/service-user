package handler

import (
	"net/http"
	"../structs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) UserRegistration(c *gin.Context){
	var (
		user structs.User
		result gin.H
	)

	nama 	 := c.PostForm("nama")
	email	 := c.PostForm("email")
	nip 	 := c.PostForm("nip")
	jabatan  := c.PostForm("jabatan")
	instansi := c.PostForm("instansi")
	kontak	 := c.PostForm("kontak")
	password := c.PostForm("password")
	role	 := c.PostForm("role")

	user.Email		= email
	user.Nama 		= nama
	user.Nip 		= nip
	user.Jabatan 	= jabatan
	user.Instansi 	= instansi
	user.Kontak 	= kontak
	user.Role 		= role

	if nama == "" || email == "" || password == "" {
		result = gin.H {
			"status": "warning!",
				"pesan": "Silahkan lengkapi lembar isian",
		}
	} else {
		cekEmail := idb.CekEmail(email)
		if cekEmail == true {
			passbyte := GetPwd(password)
			hashpass := HashAndSalt(passbyte)
			user.Password = hashpass
			status := idb.DB.Create(&user).Error
			if status != nil {
				result = gin.H {
					"status": "Error",
					"pesan": "Gagal Registrasi",
				}
			} else {
				result = gin.H {
					"status": "Success",
					"pesan": "Registrasi berhasil, silahkan login",
				}
			}
		} else  {
			result = gin.H {
				"cekEmai": cekEmail,
				"status": "warning",
				"pesan": "Email sudah digunakan",
			}
		}
	}
	c.JSON(http.StatusOK, result)
}