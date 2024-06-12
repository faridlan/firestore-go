package config

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/helper"
	"google.golang.org/api/option"
)

func NewDatabase() (*firestore.Client, error) {

	ctx := context.Background()
	config, err := helper.GetEnv()
	if err != nil {
		return nil, err
	}

	key := config.GetString("FIRESTORE_KEY")
	collection := config.GetString("COLLECTION")
	sa := option.WithCredentialsFile(key)

	client, err := firestore.NewClient(ctx, collection, sa)
	if err != nil {
		return nil, err
	}

	return client, nil

}
