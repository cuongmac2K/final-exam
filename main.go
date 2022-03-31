package main

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/middleware"
	"finnal-exam/modules/comment/commenttransport/gincomment"
	"finnal-exam/modules/commentlike/commentliketransport/gincommentlike"
	"finnal-exam/modules/post/posttransport/ginpost"
	"finnal-exam/modules/postlike/postliketransport/ginpostlike"
	"finnal-exam/modules/user/usertransport/ginuser"
	"finnal-exam/modules/userfollow/userfollowtransport/ginuserfollow"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SYSTEM_SECRET")
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db, secretKey, emailUsername, emailPassword); err != nil {
		log.Fatalln(err)
	}

}
func runService(db *gorm.DB, secretKey string, emailUsername string, emailPassword string) error {
	appCtx := component.NewAppContext(db, secretKey, emailUsername, emailPassword)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// CRUD
	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.PUT("/users/verified_email", ginuser.VerifyEmail(appCtx))
	v1.PUT("/users/send_otp/:email", ginuser.SendOtpCode(appCtx))
	v1.POST("/users/forgot_password/:email", ginuser.ForgotPassword(appCtx))
	v1.PUT("/users/reset_password/", middleware.RequiredAuthEmail(appCtx), ginuser.ResetPassword(appCtx))
	users := v1.Group("/users", middleware.RequiredAuth(appCtx))
	{
		users.POST("/update", ginuser.UpdateUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.GET("/get/:id", ginuser.GetUserbyID(appCtx))
		users.GET("/trending", ginuser.ListUserTrend(appCtx))
		users.POST("/follow/:id", ginuserfollow.FollowUser(appCtx))
		users.DELETE("/unfollow/:id", ginuserfollow.UnFollowUser(appCtx))
		users.GET("/follow/list/:id", ginuserfollow.ListUserFollow(appCtx))
		users.GET("/following/list/:id", ginuserfollow.ListUserFollowing(appCtx))
	}

	posts := v1.Group("/posts", middleware.RequiredAuth(appCtx))
	{
		posts.POST("/new", ginpost.CreatePost(appCtx))
		posts.POST("/delete/:id", ginpost.DeletePost(appCtx))
		posts.POST("/like/:id", ginpostlike.UserLikePost(appCtx))
		posts.POST("/update/:id", ginpost.UpdatePost(appCtx))
		posts.GET("", ginpost.ListPost(appCtx))
		posts.DELETE("/unlike/:id", ginpostlike.UserUnLikePost(appCtx))
		posts.GET("/trend", ginpost.TrendPost(appCtx))
		posts.GET("/get/:id", ginpost.GetPost(appCtx))
	}
	comments := v1.Group("/comments", middleware.RequiredAuth(appCtx))
	{
		comments.POST("/new", gincomment.CreateComment(appCtx))
		comments.DELETE("/unlike/:id", gincommentlike.UnlikeComment(appCtx))
		comments.POST("/like/:id", gincommentlike.UserLikeComment(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	return r.Run()
}
