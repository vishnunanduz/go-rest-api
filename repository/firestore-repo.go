package repository

import (
	"context"
	"log"

	"github.com/vishnunanduz/go-rest-api/entity"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

const (
	collectionName string = "post"
)

// NewPostrepository

func NewFireStoreRepository() PostRepo {
	return &repo{}
}

func createClient(ctx context.Context) *firestore.Client {
	// Use a service account
	sa := option.WithCredentialsFile("/home/vishnujith/Downloads/restapi-eb168-firebase-adminsdk-h258y-0941baa340.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client := createClient(ctx)

	defer client.Close()

	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return nil, err
	}
	return post, nil

}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	client := createClient(ctx)
	defer client.Close()

	var posts []entity.Post

	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (*repo) Delete(post *entity.Post) error {
	return nil
}

//FindByID: TODO
func (r *repo) FindByID(id string) (*entity.Post, error) {
	return nil, nil
}
