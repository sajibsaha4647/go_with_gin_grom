package utils

import (
	"ecommerce/domain"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidEmail(email string) bool { //if true, then valid email address
	_, err := mail.ParseAddress(email)
	return err == nil
}

// 2. SendError lives here!
func SendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, domain.Response{
		Message: message,
		Status:  statusCode,
		Data:    nil,
	})
}

// 3. SendSuccess lives here too
func SendSuccess(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, domain.Response{
		Message: message,
		Status:  statusCode,
		Data:    data,
	})
}

func BindAndValidate[T any](c *gin.Context, input *T) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		// Explicitly calling it from the utils package
		SendError(c, http.StatusBadRequest, err.Error())
		return false
	}
	return true
}
