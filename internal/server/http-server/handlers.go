package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handlerFindPerson finds person.
func (s *HTTPServer) handlerFindPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: here add service that finds person by id

	c.JSON(http.StatusOK, gin.H{"person": fmt.Sprintf("person with id = %d was found", id)})
}

// handlerDeletePerson deletes person.
func (s *HTTPServer) handlerDeletePerson(c *gin.Context) {
	personID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: here add service that deletes person by id

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("person with id = %d has been deleted", personID)})
}

// handlerUpdatePerson edits person.
func (s *HTTPServer) handlerUpdatePerson(c *gin.Context) {
	personID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: here add service that updates person

	c.JSON(http.StatusOK, gin.H{"person": fmt.Sprintf("person with id = %d has been updated", personID)})
}

// handlerAddPerson creates person.
func (s *HTTPServer) handlerAddPerson(c *gin.Context) {
	var input PersonBodyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: here add service that create person

	c.JSON(http.StatusCreated, gin.H{"person": input})
}
