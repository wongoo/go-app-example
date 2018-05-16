package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/wongoo/go-app-example/user"
)

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	u, err := user.FindUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})
}
func PostUser(c *gin.Context) {
	u := &user.User{}

	// Bind using query & form & post parameter
	err := c.Bind(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user data required", "error": err})
		return
	}

	err = user.AddUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to add user", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	err := user.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to delete user", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
