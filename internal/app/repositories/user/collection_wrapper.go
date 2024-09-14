package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockery --name=Cursor --case underscore
type Cursor interface {
	Next(ctx context.Context) bool
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
}

//go:generate mockery --name=SingleResult --case underscore
type SingleResult interface {
	Decode(val interface{}) error
}

// CollectionWrapper wraps the actual *mongo.Collection and implements the CollectionAPI interface
type CollectionWrapper struct {
	collection *mongo.Collection
}

// NewCollectionWrapper creates a new CollectionWrapper
func NewCollectionWrapper(collection *mongo.Collection) *CollectionWrapper {
	return &CollectionWrapper{collection: collection}
}

// InsertOne Implement the CollectionAPI methods by calling the corresponding methods on the wrapped *mongo.Collection
func (m *CollectionWrapper) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, document)
}

func (m *CollectionWrapper) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.collection.Find(ctx, filter, opts...)
}

func (m *CollectionWrapper) FindOne(ctx context.Context, filter interface{}) SingleResult {
	return m.collection.FindOne(ctx, filter)
}

func (m *CollectionWrapper) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}) (*mongo.UpdateResult, error) {
	return m.collection.ReplaceOne(ctx, filter, replacement)
}

func (m *CollectionWrapper) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(ctx, filter, update)
}
