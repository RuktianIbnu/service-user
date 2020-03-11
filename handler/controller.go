package handler

import (
	"../structs"
    "log"
    "fmt"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) CekEmailLogin(email string) bool{
	var (
		user structs.User
	)
	idb.DB.Where("email = ?", email).First(&user)
	if user.Email == "" {
		return false	
	} else {
		return true
	}
}

func (idb *InDB) CekEmail(email string) bool{
	var (
		user structs.User
	)
	isExsist := idb.DB.Where("email = ?", email).First(&user).Error
	if isExsist == nil {
		return false	
	} else {
		return true
	}
}

func (idb *InDB) CekLogin(email string, password string) bool{
	var (
		user structs.User
	)
	isExsist := idb.DB.Where("email = ? AND password = ?", email, password).Find(&user).Error
	if isExsist == nil {
		return false	
	} else {
		return true
	}
}

func GetPwd(pass string) []byte {
    return []byte(pass)
}

func HashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func ComparePasswords(PassFromDb string, EncryptPassInput []byte) bool {
    byteHash := []byte(PassFromDb)
    err := bcrypt.CompareHashAndPassword(byteHash, EncryptPassInput)
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}

func Auth(c *gin.Context) {
	var result gin.H
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("hS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		result = gin.H{
			"status": "success",
			"pesan": "token verified",
		}
		c.JSON(http.StatusOK, result)
		c.Abort()
	} else {
		result = gin.H {
			"status": "error " + err.Error(),
			"pesan": "not authorized",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}