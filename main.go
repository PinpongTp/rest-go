package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"pinpong.co/rest-go/models"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func (route *App) createNoteHandler(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.New().String()
	note.Id = id

	var err error
	_, _, err = route.client.Collection("notes").Add(route.ctx, &note)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, &note)
}

func (route *App) findById(c *gin.Context) {
	id := c.Param("id")
	NotesData := []models.Note{}

	fmt.Printf(">> find by id: %v", id)
	fmt.Println(id)
	item := route.client.Collection("notes").Where("Id", "==", id).Documents(route.ctx)
	for {
		NoteData := models.Note{}
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		mapstructure.Decode(doc.Data(), &NoteData)
		NotesData = append(NotesData, NoteData)
	}
	c.JSON(http.StatusOK, NotesData)
}

type App struct {
	//Router *mux.Router
	client *firestore.Client
	ctx    context.Context
}

func (route *App) Init() {
	route.ctx = context.Background()
	opt := option.WithCredentialsFile("my-test-firebase-e0b24-firebase-adminsdk-88aai-3b0e645857.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	route.client, err = app.Firestore(route.ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func (route *App) get(c *gin.Context) {
	NotesData := []models.Note{}
	item := route.client.Collection("notes").Documents(route.ctx)
	for {
		NoteData := models.Note{}
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		mapstructure.Decode(doc.Data(), &NoteData)
		NotesData = append(NotesData, NoteData)
	}
	c.JSON(http.StatusOK, NotesData)
}

func (route *App) updateNoteHandler(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(note.Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not have id",
		})
		return
	}

	var refId string

	// get Ref.ID for update
	item := route.client.Collection("notes").Where("Id", "==", note.Id).Documents(route.ctx)
	for {
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		refId = doc.Ref.ID
	}

	// update
	var err error
	_, err = route.client.Collection("notes").Doc(refId).Set(route.ctx, note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, note)
	return
}

func (route *App) deleteNoteHandler(c *gin.Context) {
	id := c.Param("id")
	var refId string

	// get Ref.ID for delete
	item := route.client.Collection("notes").Where("Id", "==", id).Documents(route.ctx)
	for {
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		refId = doc.Ref.ID
	}

	// delete
	var err error
	_, err = route.client.Collection("notes").Doc(refId).Delete(route.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, "")
	return
}

func main() {
	route := App{}
	route.Init()

	r := gin.New()
	r.GET("/notes", route.get)
	r.GET("/notes/:id", route.findById)
	r.POST("/notes", route.createNoteHandler)
	r.PUT("/notes", route.updateNoteHandler)
	r.DELETE("/notes/:id", route.deleteNoteHandler)
	r.Run()
}
