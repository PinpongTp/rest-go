package deliveries

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pinpong.co/rest-go/models"
	"pinpong.co/rest-go/usecases"
)

type NoteHandler struct {
	noteUseCase usecases.NoteUseCase
}

func NewNoteHandler(noteUseCase usecases.NoteUseCase) *NoteHandler {
	return &NoteHandler{noteUseCase}
}

func (t *NoteHandler) FindAll(c *gin.Context) {
	res, err := t.noteUseCase.FindAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (t *NoteHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	res, err := t.noteUseCase.FindById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (t *NoteHandler) Create(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(">> test detect error")
		return
	}

	res, err := t.noteUseCase.Create(note)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (t *NoteHandler) Update(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(">> test detect error")
		return
	}

	res, err := t.noteUseCase.Update(note)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (t *NoteHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := t.noteUseCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusNoContent)
	}
}
