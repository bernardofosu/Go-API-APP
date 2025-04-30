// package controllers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetUsers(c *gin.Context) {
// 	users := []string{"Alice", "Bob"}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 		"data":   users,
// 	})
// }

package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"my-go-gin-app/config"
)

func GetUsers(c *gin.Context) {
	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(ctx)

	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"users": users,
	})
}
