package user

import (
	"context"
	"time"

	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "users"

//go:generate mockery --name=Collection --case underscore
type Collection interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}) SingleResult
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
}

type Repository struct {
	Db Collection
}

func New(db *mongo.Database) *Repository {
	return &Repository{Db: NewCollectionWrapper(db.Collection(collectionName))}
}

func (r *Repository) CreateUser(user models.User) error {
	_, err := r.Db.InsertOne(context.Background(), user)
	return err
}

func (r *Repository) GetUsers(filter interface{}, offset, limit int) ([]models.User, error) {
	var users []models.User
	ctx := context.Background()

	cursor, err := r.Db.Find(ctx, filter, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, cursor.Err()
}

func (r *Repository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	result := r.Db.FindOne(context.Background(), bson.M{"_id": id})
	err := result.Decode(&user)
	return user, err
}

func (r *Repository) UpdateUser(id primitive.ObjectID, user models.User) error {
	_, err := r.Db.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	return err
}

func (r *Repository) DeleteUser(id primitive.ObjectID) error {
	now := time.Now()
	_, err := r.Db.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": &now}})
	return err
}
