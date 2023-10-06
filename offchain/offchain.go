package offchain

import (
	"go.mongodb.org/mongo-driver/bson"

	U "github.com/off-chain-storage/go-off-chain-storage/utils"

	clusterclient "github.com/off-chain-storage/go-off-chain-storage/clusterclient"
)

type UserInfo struct {
	Index  int
	Gender string
	UserID string
	Name   string
}

func CreateDoc_off_chain() {
	data := UserInfo{
		Index:  1,
		Gender: "M",
		UserID: "bum0448",
		Name:   "jinbumjinbum",
	}

	client, ctx, cancel, err := clusterclient.GetMongoClient()
	U.CheckErr(err)

	defer client.Disconnect(ctx)
	defer cancel()

	db := client.Database("myDB")
	collection := db.Collection("fs.files")

	_, err = collection.InsertOne(ctx, data)
	U.CheckErr(err)
}

func ReadDoc_off_chain() {
	client, ctx, cancel, err := clusterclient.GetMongoClient()
	U.CheckErr(err)

	defer client.Disconnect(ctx)
	defer cancel()

	db := client.Database("myDB")
	collection := db.Collection("fs.files")

	var result UserInfo
	filter := bson.M{"name": "jinbumjinbum"}

	err = collection.FindOne(ctx, filter).Decode(&result)
	U.CheckErr(err)
}

func UpdateDoc_off_chain() {
	client, ctx, cancel, err := clusterclient.GetMongoClient()
	U.CheckErr(err)

	defer client.Disconnect(ctx)
	defer cancel()

	db := client.Database("myDB")
	collection := db.Collection("fs.files")

	filter := bson.M{"name": "jinbumjinbum"}
	update := bson.M{"$set": bson.M{
		"flag": "Red",
		"role": "magicion",
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	U.CheckErr(err)
}

func DeleteDoc_off_chain() {
	client, ctx, cancel, err := clusterclient.GetMongoClient()
	U.CheckErr(err)

	defer client.Disconnect(ctx)
	defer cancel()

	db := client.Database("myDB")
	collection := db.Collection("fs.files")

	filter := bson.M{"name": "jinbumjinbum"}

	_, err = collection.DeleteOne(ctx, filter)
	U.CheckErr(err)
}
