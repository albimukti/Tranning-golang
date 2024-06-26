package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Age            int       `json:"age"`
	RiskScore      int       `json:"risk_score"`
	RiskCategory   string    `json:"risk_category"`
	RiskDefinition string    `json:"risk_definition"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Submission struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	Answers      string    `json:"answers"`
	RiskScore    int       `json:"risk_score"`
	RiskCategory string    `json:"risk_category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `gorm:"foreignKey:UserID"`
}

func main() {
	// Database connection string
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "albi123"
		dbname   = "users"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&User{}, &Submission{})
	if err != nil {
		panic("failed to auto migrate database schema")
	}

	// Set up Gin
	r := gin.Default()

	// Routes
	r.GET("/users/:id", getUserByID)
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.GET("/submissions/:id", getSubmissionByID)
	r.GET("/submissions", getSubmissions)
	r.POST("/submissions", createSubmission)
	r.DELETE("/submissions/:id", deleteSubmission)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		panic("failed to start server")
	}
}

// User Handlers
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var users []User
	DB.Offset((page - 1) * limit).Limit(limit).Find(&users)

	var total int64
	DB.Model(&User{}).Count(&total)
	c.JSON(http.StatusOK, gin.H{
		"users":        users,
		"total_pages":  (total + int64(limit) - 1) / int64(limit),
		"current_page": page,
	})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user_id": user.ID})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UpdatedAt = time.Now()
	DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Submission Handlers
func getSubmissionByID(c *gin.Context) {
	id := c.Param("id")
	var submission Submission
	if err := DB.Preload("User").First(&submission, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func getSubmissions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var submissions []Submission
	DB.Preload("User").Offset((page - 1) * limit).Limit(limit).Find(&submissions)

	var total int64
	DB.Model(&Submission{}).Count(&total)
	c.JSON(http.StatusOK, gin.H{
		"submissions":  submissions,
		"total_pages":  (total + int64(limit) - 1) / int64(limit),
		"current_page": page,
	})
}

func createSubmission(c *gin.Context) {
	var submission Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission.CreatedAt = time.Now()
	submission.UpdatedAt = time.Now()
	DB.Create(&submission)
	c.JSON(http.StatusOK, gin.H{"message": "Submission created successfully", "submission_id": submission.ID})
}

func deleteSubmission(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&Submission{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submission deleted successfully"})
}
