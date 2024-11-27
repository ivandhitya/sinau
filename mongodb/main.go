package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ivandhitya/sinau/mongodb/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sales struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Item string             `bson:"item"`
}

func main() {
	conn := mongodb.ConnectMongoDB("mongodb://mongo:gTNiJdQlwtBOHwZBJBQWDZFoVTlBGxVF@junction.proxy.rlwy.net:24268")

	collection := getCollection(conn, "sinau_mongo", "sales")

	doc := map[string]interface{}{
		"item":     "mouse",
		"price":    30000,
		"quantity": 1,
		"category": "elektronik",
	}
	err := insertDocument(collection, doc)
	if err != nil {
		log.Fatal(err)
	}
	results, err := findDocuments(collection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dokumen:")
	for _, v := range results {
		fmt.Println(v.ID, v.Item)
	}

	// Filter dokumen yang akan diupdate
	filter := Sales{Item: "mouse"}
	err = updateDocument(collection, filter)
	if err != nil {
		log.Fatal(err)
	}

	results, err = findDocuments(collection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dokumen setelah update:")
	for _, v := range results {
		fmt.Println(v.ID, v.Item)
	}

	filter = Sales{Item: "headphone"}
	err = deleteDocument(collection, filter)
	if err != nil {
		log.Fatal(err)
	}

	results, err = findDocuments(collection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dokumen setelah delete:")
	for _, v := range results {
		fmt.Println(v.ID, v.Item)
	}

}

func getCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func findDocuments(collection *mongo.Collection) ([]Sales, error) {
	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cari semua dokumen
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("Gagal membaca dokumen: %v", err)
	}
	defer cursor.Close(ctx)

	var results []Sales

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("Iterasi Error: %v", err)
		return nil, err
	}
	return results, nil
}

func insertDocument(collection *mongo.Collection, data map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert ke MongoDB
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Printf("Dokumen berhasil disimpan dengan ID: %v\n", result.InsertedID)
	return nil
}

func updateDocument(collection *mongo.Collection, filter Sales) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Data baru untuk update
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"item": "headphone",
		},
	}

	// Update dokumen
	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("Banyak dokumen yang diupdate: %d\n", result.ModifiedCount)
	return nil
}

func deleteDocument(collection *mongo.Collection, filter Sales) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Hapus beberapa dokumen
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	fmt.Printf("Banyak dokumen yang dihapus: %d\n", result.DeletedCount)
	return nil
}
