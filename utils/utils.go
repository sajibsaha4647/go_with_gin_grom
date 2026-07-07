package utils

import (
	"ecommerce/domain"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"

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

func SendLoginSuccess(c *gin.Context, statusCode int, message string, data any, token string) {
	c.JSON(statusCode, domain.ResponseLogin{
		Message: message,
		Status:  statusCode,
		Data:    data,
		Token:   token,
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

func BindAndValidateFromData[T any](c *gin.Context, input *T) bool {
	if err := c.ShouldBind(input); err != nil {
		// Explicitly calling it from the utils package
		SendError(c, http.StatusBadRequest, err.Error())
		return false
	}
	return true
}

func UploadImageFile(file *multipart.FileHeader, folder string) (string, error) {

	//uploads/product
	uploadPath := filepath.Join("uploads", folder)

	// Automatically create folder if it doesn't exist

	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return "", fmt.Errorf("only jpg, jpeg, png and webp are allowed")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	// uploads/products/xxxxx.jpg
	dst := filepath.Join(uploadPath, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return filename, nil

}
