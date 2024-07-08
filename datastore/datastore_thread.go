package datastore

import (
	"context"
	"github.com/google/uuid"
	"github.com/pzierahn/chatbot_services/llm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Thread struct {
	// ID of the message
	Id uuid.UUID `bson:"_id,omitempty"`

	// Thread ID
	ThreadId uuid.UUID `bson:"thread_id,omitempty"`

	// User ID
	UserId string `bson:"user_id,omitempty"`

	// Collection ID
	CollectionId uuid.UUID `bson:"collection_id,omitempty"`

	// Timestamp of the last message
	Timestamp time.Time `bson:"timestamp,omitempty"`

	// Messages
	Messages []*llm.Message `bson:"messages,omitempty"`
}

// StoreThread stores a thread
func (service *Service) StoreThread(ctx context.Context, thread *Thread) error {
	coll := service.mongo.Database(DatabaseName).Collection(CollectionMessages)

	_, err := coll.InsertOne(ctx, thread)
	if err != nil {
		return err
	}

	return nil
}

// GetThread returns all messages of a thread
func (service *Service) GetThread(ctx context.Context, userId string, threadId uuid.UUID) (*Thread, error) {
	coll := service.mongo.Database(DatabaseName).Collection(CollectionMessages)

	filter := bson.M{
		"thread_id": threadId,
		"user_id":   userId,
	}

	var thread Thread
	err := coll.FindOne(ctx, filter).Decode(&thread)
	if err != nil {
		return nil, err
	}

	return &thread, nil
}

// GetThreadIDs returns all thread IDs of a user
func (service *Service) GetThreadIDs(ctx context.Context, userId string) ([]uuid.UUID, error) {
	coll := service.mongo.Database(DatabaseName).Collection(CollectionMessages)

	filter := bson.M{
		"user_id": userId,
	}

	ops := &options.FindOptions{
		Projection: bson.M{
			"thread_id": 1,
		},
	}

	cursor, err := coll.Find(ctx, filter, ops)
	if err != nil {
		return nil, err
	}

	defer func() { _ = cursor.Close(ctx) }()

	var result []uuid.UUID
	for cursor.Next(ctx) {
		var thread Thread
		err = cursor.Decode(&thread)
		if err != nil {
			return nil, err
		}

		result = append(result, thread.ThreadId)
	}

	return result, nil
}
