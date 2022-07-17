package fb

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firebase struct {
	Client *firestore.Client
	Ctx    context.Context
}

func NewFirebase() *Firebase {
	ctx := context.Background()
	opt := option.WithCredentialsFile("util/my-test-firebase-e0b24-firebase-adminsdk-88aai-3b0e645857.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return &Firebase{
		Client: client,
		Ctx:    ctx,
	}
}
