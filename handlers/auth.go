package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxonbejenari/testWebApp/models"
	"github.com/maxonbejenari/testWebApp/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}
		if !utils.CheckPasswordHash(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}
		token, err := utils.GenerateJWT(user.Username, string(user.Role))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not generate token",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"role":    user.Role,
			"expires": time.Now().Add(time.Hour * 24),
		})
	}
}
