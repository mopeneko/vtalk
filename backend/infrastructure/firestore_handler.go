package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type FirestoreHandler struct {
	conn *firestore.Client
}

func NewFirestoreHandler(projectID string) (*FirestoreHandler, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	handler := &FirestoreHandler{client}
	return handler, nil
}
