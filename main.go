package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	conf := &firebase.Config{
		StorageBucket: "securely-374619.appspot.com",
	}
	opt := option.WithCredentialsFile("securely-374619-firebase-adminsdk-yd74p-2ca4c4a753.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Get a client to interact with Firebase Cloud Storage
	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Cloud Storage client: %v", err)
	}

	// Upload a file to Firebase Cloud Storage
	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default Firebase Cloud Storage bucket: %v", err)
	}
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	obj := bucket.Object("example.txt")
	wc := obj.NewWriter(ctx)
	if _, err = wc.Write(data); err != nil {
		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
	}
	fmt.Println("File uploaded successfully!")

}
