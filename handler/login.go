package handler

import (
	"net/http"
	"../structs"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) LoginUser(c *gin.Context) {
	var (
		user structs.User
		newUser structs.User
		err = c.Bind(&user)
	)

	email := c.PostForm("email")
	password := c.PostForm("password") 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"pesan": "tidak bisa bind Structs",
		})
	}

	EmailIsExist := idb.CekEmail(email)
	Pass := GetPwd(password) 
	//EncryptPass := HashAndSalt(Pass)

	idb.DB.Select("password").Where(" email = ?", email).Find(&user)
	passToString := user.Password
	 
	matchPass := ComparePasswords(passToString, Pass) 

	if EmailIsExist == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Nip Tidak Terdaftar",
			"passToString": Pass,
		})
	} else if matchPass == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Password Salah",
		})
	} else if matchPass == true && EmailIsExist == true {
		sign := jwt.New(jwt.GetSigningMethod("HS256")) // hs256 
		
		claims := sign.Claims.(jwt.MapClaims)
		claims["nama"] 		= user.Nama
		claims["email"] 	= user.Email
		claims["jabatan"]	= user.Jabatan
		claims["role"]		= user.Role

		token, err := sign.SignedString([]byte("secret"))

		newUser.Token = token
		update := idb.DB.Model(&user).Where("email = ?", email).Updates(newUser).Error

		if update != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"pesan": err.Error(),
				"status": "error",
			})
			c.Abort()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"pesan": err.Error(),
			})
			c.Abort()
		}
		
		c.JSON(http.StatusOK, gin.H {
			"token": token,
			"data_user": claims,
			"status": "success",
			"pesan": "Login Berhasil",
			"detail_user": user,
		})
	}
}