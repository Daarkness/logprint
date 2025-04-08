package web

import (
	"fmt"
	"github.com/Daarkness/logprint/web-basic/webook/internal/domain"
	"github.com/Daarkness/logprint/web-basic/webook/internal/repository"
	"github.com/Daarkness/logprint/web-basic/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	//svc "github.com/Daarkness/logprint/web-basic/webook/internal/service"
	"net/http"
	//"regexp"
)

//type UserHandler struct {
//}

type UserHandler struct {
	emailRegex    *regexp.Regexp
	passwordRegex *regexp.Regexp
	svc           service.UserService
}

const (
	emailRegex    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	passwordRegex = `^[a-zA-Z0-9._%+\-]+$`
)

func NewUserHandler() *UserHandler {
	//return &UserHandler{
	//	emailRegex:    regexp.MustCompile(emailRegex, regexp.None),
	//	passwordRegex: regexp.MustCompile(passwordRegex, regexp.None),
	//}
	return &UserHandler{
		emailRegex:    regexp.MustCompile(emailRegex, regexp.None),
		passwordRegex: regexp.MustCompile(passwordRegex, regexp.None),
		svc:           service.NewUserService(repository.NewUserRepository()),
	}
}

// func (u *UserHandler) RegisetrRoutes(r *gin.Engine) {
//
//		r.GET("/user/signup", u.SignUp)
//		r.POST("/user/login", u.Login)
//		r.POST("/user/edit", u.Edit)
//		r.POST("/user/profile", u.Profile)
//	}

func (u *UserHandler) RegisetrRoutes(r *gin.Engine) {

	ug := r.Group("/user")
	{
		ug.POST("/signup", u.SignUp)
		ug.POST("/login", u.Login)
		ug.POST("/edit", u.Edit)
		ug.POST("/profile", u.Profile)
	}

}

type U struct {
	Username        string `json:"username" binding:"required,min=4"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"ConfirmPassword"`
	//Email    string `json:"email" binding:"required,email"`
	Email string `json:"email" `
	//req   domain.User
}

func (u *UserHandler) SignUp(c *gin.Context) {
	req := NewUserHandler()
	fmt.Println("userSignUp ")
	req := U{
		Username: c.PostForm("username"),
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟业务逻辑（如保存到数据库）
	//fmt.Printf("user: %#v\n", req.Username)
	fmt.Printf("password: %s\n", req.Password)
	fmt.Printf("password: %s\n", req.Password)
	fmt.Printf("email: %s\n", req.Email)
	//emailExp := regexp.MustCompile(emailRegex, regexp.None)

	//ok, err := emailExp.FindString(req.Email), err
	//ok, err := regexp.MatchString(emailRegex, req.Email)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//if req.ConfirmPassword != req.ConfirmPassword {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Password confirmation does not match"})
	//}
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱格式错误"})
	//	return
	//}
	//_, err := c.emailRegex.MatchString(req.Email)
	//u := NewUserHandler()
	ok, err := userhaader.emailRegex.MatchString(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱格式错误"})
		return
	}
	type User struct {
		svc repository.UserRepository
	}
	u := NewUserHandler()
	err := u.svc.SignUp(c, domain.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data":    req,
	})
}
func (u *UserHandler) Edit(c *gin.Context) {
	fmt.Println("SignUp")
}
func (u *UserHandler) Login(c *gin.Context) {
	fmt.Println("SignUp")
}
func (u *UserHandler) Profile(c *gin.Context) {
	fmt.Println("SignUp")
}
