package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string) *mongo.Client {

	// Context dengan timeout untuk koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Koneksi ke MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Gagal terhubung ke MongoDB: %v", err)
	}

	// Tes koneksi
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Ping MongoDB Gagal: %v", err)
	}

	fmt.Println("Berhasil terhubung ke MongoDB!")
	return client
}
