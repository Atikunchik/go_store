package Controller

import (
	_ "GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	Validate(ctx *gin.Context)
}

func (c *Controller) GetUser(ctx *gin.Context) {
	var user model.User
	UserLogin := ctx.Param("login")
	ctx.ShouldBindJSON(&user)
	prod, err := c.service.GetUser(UserLogin)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this Login"})
		return
	}
	ctx.JSON(200, gin.H{"message": prod})
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var cur_user model.User
	err := ctx.ShouldBindJSON(&cur_user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!"})
		return
	}
	fmt.Println(cur_user.Login)
	fmt.Println(cur_user.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(cur_user.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!"})
		return
	}
	user := model.User{Login: cur_user.Login, Password: string(hash)}
	c.service.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user created"})

}

func (c *Controller) Login(ctx *gin.Context) {
	var cur_user model.User
	err := ctx.ShouldBindJSON(&cur_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!"})
		return
	}

	var user model.User
	user, err = c.service.GetUser(cur_user.Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt find"})
		return

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cur_user.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "compare"})
		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.Login,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (c *Controller) Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userlogin, _ := ctx.Get("userlogin")
	str_login := userlogin.(string)
	fmt.Println(str_login)
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": user,
	})
}
