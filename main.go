package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes = []Note{
	{ID: "1", Title: "test first time", Content: "นี้คือบทความแรกของฉัน"},
	{ID: "2", Title: "บทความที่ 2", Content: "this is second note."},
}

func listNotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, notes)
}

func createNoteHandler(c *gin.Context) {
	var note Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	notes = append(notes, note)

	c.JSON(http.StatusCreated, note)
}

func deleteNoteHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range notes {
		if a.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/notes", listNotesHandler)
	r.POST("/notes", createNoteHandler)
	r.DELETE("/notes/:id", deleteNoteHandler)
	r.Run()
}
