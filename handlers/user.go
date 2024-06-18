package handlers

import (
	"my-gin-api/logger"
	"my-gin-api/models"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUsers(c *gin.Context) {
	log := logger.GetLogger()
	requestID := c.MustGet("request_id").(string)
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("Failed to fetch users: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, users)
	log.WithFields(logrus.Fields{
		"request-id": requestID,
		"endpoint":   "/users",
		"method":     "GET",
	}).Info("Handling POST /login request")
}

func GetUserByID(c *gin.Context) {
	log := logger.GetLogger()
	requestID := c.MustGet("request_id").(string)
	id := c.Param("id")

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		log.Printf("Invalid user ID: %s\n", id)
		return
	}

	var user models.User
	if err := DB.First(&user, uintID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "GET",
		}).Error("Failed to fetch user with ID: %d\n", uintID)
		return
	}
	c.JSON(http.StatusOK, user)
	log.WithFields(logrus.Fields{
		"request-id": requestID,
		"endpoint":   "/users/:id",
		"method":     "GET",
	}).Info("Fetched user with ID: %d\n", user.ID)
}

func CreateUser(c *gin.Context) {
	log := logger.GetLogger()
	requestID := c.MustGet("request_id").(string)
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Failed to bind user data: %v\n", err)
		return
	}

	if err := newUser.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users",
			"method":     "POST",
		}).Error("User validation failed: %v\n", err)
		return
	}

	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users",
			"method":     "POST",
		}).Error("Failed to create user: %v\n", err)
		return
	}

	c.JSON(http.StatusCreated, newUser)
	log.WithFields(logrus.Fields{
		"request-id": requestID,
		"endpoint":   "/users",
		"method":     "POST",
	}).Info("Created new user with ID: %d\n", newUser.ID)
}

func UpdateUser(c *gin.Context) {
	log := logger.GetLogger()
	requestID := c.MustGet("request_id").(string)
	id := c.Param("id")

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("Invalid user ID: %s\n", id)
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("Failed to bind user data: %v\n", err)
		return
	}

	if err := updatedUser.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("User validation failed: %v\n", err)
		return
	}

	if err := DB.Model(&models.User{}).Where("id = ?", uintID).Updates(updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("Failed to update user with ID: %d\n", uintID)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
	log.WithFields(logrus.Fields{
		"request-id": requestID,
		"endpoint":   "/users/:id",
		"method":     "PUT",
	}).Info("Updated user with ID: %d\n", uintID)
}

func DeleteUser(c *gin.Context) {
	log := logger.GetLogger()
	requestID := c.MustGet("request_id").(string)
	id := c.Param("id")

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("Invalid user ID: %s\n", id)
		return
	}

	if err := DB.Delete(&models.User{}, uintID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.WithFields(logrus.Fields{
			"request-id": requestID,
			"endpoint":   "/users/:id",
			"method":     "PUT",
		}).Error("Failed to delete user with ID: %d\n", uintID)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	log.WithFields(logrus.Fields{
		"request-id": requestID,
		"endpoint":   "/users/:id",
		"method":     "PUT",
	}).Info("Deleted user with ID: %d\n", uintID)
}
