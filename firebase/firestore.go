package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreDB struct {
	client *firestore.Client
}

var dbInstance *FirestoreDB
var once sync.Once

func DBClient() *firestore.Client {
	once.Do(func() {
		dbInstance = new(FirestoreDB)
		makeDBClient(dbInstance)
	})
	return dbInstance.client
}

// makeDBClient Client 주입
func makeDBClient(rtb *FirestoreDB) {
	// Init FireStore SDK
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	ctx := context.Background()
	client, err := firestore.NewClient(
		ctx, os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"), opt)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	rtb.client = client
}
