package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *Handler) listNotesHandler(c *gin.Context) {
	var notes []Note
	if result := h.db.Find(&notes); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &notes)
}

func (h *Handler) createNoteHandler(c *gin.Context) {
	var note Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if result := h.db.Create(&note); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &note)
}

func (h *Handler) deleteNoteHandler(c *gin.Context) {
	id := c.Param("id")
	if result := h.db.Delete(&Note{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Note{})

	handler := newHandler(db)

	r := gin.New()
	r.GET("/notes", handler.listNotesHandler)
	r.POST("/notes", handler.createNoteHandler)
	r.DELETE("/notes/:id", handler.deleteNoteHandler)
	r.Run()
}
