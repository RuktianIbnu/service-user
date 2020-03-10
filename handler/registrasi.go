package handler

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
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
	user.Password 	= password
	user.Role 		= role

	if nama == "" || email == "" || password == "" {
		result = gin.H {
			"status": "warning!",
				"pesan": "Silahkan lengkapi lembar isian",
		}
	} else {
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
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) LoginUser(c gin.Context) {
	var (
		user structs.User
		result gin.H
	)

	email := c.PostForm("email")
	password :=  c.PostForm("password")

	cekUser := idb.CekUser(email, password)
	if cekUser == false {
		result = gin.H{
			"status": "warning",
			"pesan": "email dan password tidak sesuai",
		}
	} else {
		result = gin.H{
			"status": "success",
			"pesan": "Berhasil login",
		}
	}
	c.JSON(http.StatusOK, result)
}