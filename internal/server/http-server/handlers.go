package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// handlerFindPerson finds person.
func (s *HTTPServer) handlerFindPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// here is service that finds person by id
	person, err := s.service.FindPersonByID(c.Request.Context(), uint64(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"person": person})
}

// handlerDeletePerson deletes person.
func (s *HTTPServer) handlerDeletePerson(c *gin.Context) {
	personID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// here is service that deletes person by id
	id, err := s.service.DeletePerson(c.Request.Context(), uint64(personID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("person with id = %d has been deleted", id)})
}

// handlerUpdatePerson edits person.
func (s *HTTPServer) handlerUpdatePerson(c *gin.Context) {
	personID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// here is service that updates person
	var newParams entities.Person
	if err := c.ShouldBindJSON(&newParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := s.service.UpdatePerson(c.Request.Context(), uint64(personID), newParams)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("person with id = %d has been updated", personID), "person": person})
}

// handlerAddPerson creates person.
func (s *HTTPServer) handlerAddPerson(c *gin.Context) {
	person := entities.Person{}
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.service.CreatePerson(c.Request.Context(), person)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	person.ID = id

	c.JSON(http.StatusCreated, gin.H{"person": person})
}
