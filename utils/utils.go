package utils

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// HashPassword generate hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword compare password
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetPagination from query params
func GetPagination(c *gin.Context) (page int, limit int) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ = strconv.Atoi(pageStr)
	limit, _ = strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return
}
