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

	emailORnip := c.PostForm("emailORnip")
	password := c.PostForm("password") 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"pesan": "tidak bisa bind Structs",
		})
	}

	EmailOrNipIsExist := idb.CekEmailOrNipLogin(emailORnip)
	Pass := GetPwd(password) 

	idb.DB.Where(" email = ?", emailORnip).First(&user)
	if user.Email == "" {
		idb.DB.Where(" nip = ?", emailORnip).First(&user)
	}

	passToString := user.Password 
	matchPass := ComparePasswords(passToString, Pass) 

	if EmailOrNipIsExist == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Email Atau Nip Tidak Terdaftar",
		})
	} else if matchPass == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Password Salah",
		})
	} else if matchPass == true && EmailOrNipIsExist == true {
		sign := jwt.New(jwt.GetSigningMethod("HS256")) // hs256 
		
		claims := sign.Claims.(jwt.MapClaims)
		claims["ID"]		= user.ID
		claims["nama"]		= user.Nama
		claims["email"]		= user.Email
		claims["nip"]		= user.Nip
		claims["jabatan"]	= user.Jabatan
		claims["role"]		= user.Role
		claims["instansi"]	= user.Instansi
		
		token, err := sign.SignedString([]byte("secret"))

		newUser.Token = token
		update := idb.DB.Model(&user).Where("id = ?", user.ID).Updates(newUser).Error
		
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
			"session_user": claims,
			
			"status": "success",
			"pesan": "Login Berhasil",
			"token": token,
		})
	}
}