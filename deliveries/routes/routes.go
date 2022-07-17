package routes

import (
	"github.com/gin-gonic/gin"
	"pinpong.co/rest-go/deliveries"
	"pinpong.co/rest-go/repositories"
	"pinpong.co/rest-go/usecases"
)

func SetupRouter() *gin.Engine {
	noteRepo := repositories.NewNoteRepository()
	noteUseCase := usecases.NewNoteUseCase(*noteRepo)
	noteHandler := deliveries.NewNoteHandler(*noteUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("note", noteHandler.FindAll)
		v1.GET("note/:id", noteHandler.FindById)
		v1.POST("note", noteHandler.Create)
		v1.PUT("note", noteHandler.Update)
		v1.DELETE("note/:id", noteHandler.Delete)
	}
	return r
}
