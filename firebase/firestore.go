package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	U "github.com/off-chain-storage/comparing-offchain-firebase/utils"
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

func CreateDoc() {
	ctx := context.Background()
	client := DBClient()

	_, err := client.Doc("User/User13818").Create(ctx, map[string]interface{}{
		"Index":  1,
		"Gender": "M",
		"UserID": "bum0448",
		"Name":   "jinbumjinbum",
	})
	U.CheckErr(err)
}

func UpdateDoc() {
	ctx := context.Background()
	client := DBClient()

	if _, err := client.Doc("User/User13818").
		Update(ctx, []firestore.Update{
			{"FlagColor", nil, "Red"},
			{Path: "Location", Value: "Middle"}}); err != nil {
		log.Fatalf("Update error: %s\n", err)
	}
}

func ReadDoc() {
	ctx := context.Background()
	client := DBClient()

	state, err := client.Doc("User/User13818").Get(ctx)
	U.CheckErr(err)

	_, err = json.MarshalIndent(state.Data(), "", "  ")
	U.CheckErr(err)
}

func DeleteDoc() {
	ctx := context.Background()
	client := DBClient()

	_, err := client.Doc("User/User13818").Delete(ctx)
	U.CheckErr(err)
}
