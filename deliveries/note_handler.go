package deliveries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pinpong.co/rest-go/usecases"
)

type NoteHandler struct {
	noteUseCase usecases.NoteUseCase
}

func NewNoteHandler() *NoteHandler {
	noteUseCase := usecases.NewNoteUseCase()
	return &NoteHandler{*noteUseCase}
}

func (t *NoteHandler) GetAll(c *gin.Context) {
	res, err := t.noteUseCase.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
