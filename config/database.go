package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewDatabase() (*firestore.Client, error) {

	ctx := context.Background()

	serviceAccountKey := os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY")
	if serviceAccountKey == "" {
		log.Fatal("FIREBASE_SERVICE_ACCOUNT_KEY environment variable not set")
	}

	collection := os.Getenv("COLLECTION")

	var sa map[string]any

	err := json.Unmarshal([]byte(serviceAccountKey), &sa)
	if err != nil {
		return nil, fmt.Errorf("error parsing service account key: %v", err)
	}

	opt := option.WithCredentialsJSON([]byte(serviceAccountKey))

	client, err := firestore.NewClient(ctx, collection, opt)
	if err != nil {
		return nil, err
	}

	return client, nil

}
