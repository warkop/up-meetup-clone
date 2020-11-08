package cmd

import (
	"net/http"
	"up-meetup-clone/repository"

	"github.com/gin-gonic/gin"
)

func defaultStatus() map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": "Ok",
	}
}

// ListTitle is get all data from db
func ListTitle(c *gin.Context) {
	title := repository.AllTitle()

	c.JSON(http.StatusOK, gin.H{"status": defaultStatus(), "data": title})
}

// DetailTitle is showing detail of title
func DetailTitle(c *gin.Context) {
	id := c.Param("id")
	title := repository.FindByID(id)

	c.JSON(http.StatusOK, gin.H{"status": defaultStatus(), "data": title})
}
