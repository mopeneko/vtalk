package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
)

type FirestoreHandler struct {
	conn *firestore.Client
}

func NewFirestoreHandler(projectID string) (*FirestoreHandler, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	handler := &FirestoreHandler{client}
	return handler, nil
}
