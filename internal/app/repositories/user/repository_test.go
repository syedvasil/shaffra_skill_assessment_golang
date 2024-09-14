package user_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	repoModels "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/user"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/user/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	mockCollection := new(mocks.Collection)

	repo := &user.Repository{Db: mockCollection}

	mockUser := repoModels.User{
		ID:       primitive.NewObjectID(),
		Username: "Test User",
		Email:    "test@example.com",
		Age:      12,
		Password: "usele",
	}

	mockCollection.On("InsertOne", mock.Anything, mockUser).Return(&mongo.InsertOneResult{}, nil)

	err := repo.CreateUser(mockUser)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

//func Test_GetUsers(t *testing.T) {
//	t.Parallel()
//	mockCollection := new(mocks.Collection)
//	repo := &user.Repository{Db: mockCollection}
//
//	filter := bson.M{}
//	mockCursor := new(mocks.Cursor)
//
//	mockCursor.On("Next", mock.Anything).Return(true).Once()
//	mockCursor.On("Decode", mock.AnythingOfType("*repoModels.User")).Run(func(args mock.Arguments) {
//		userMock := args.Get(0).(*repoModels.User)
//		*userMock = repoModels.User{
//			ID:       primitive.NewObjectID(),
//			Username: "Test User",
//			Email:    "test@example.com",
//		}
//	}).Return(nil)
//	mockCursor.On("Next", mock.Anything).Return(false).Once()
//
//	mockCollection.On("Find", mock.Anything, filter, mock.Anything).Return(mockCursor, nil)
//
//	users, err := repo.GetUsers(filter, 0, 10)
//	assert.NoError(t, err)
//	assert.Len(t, users, 1)
//	assert.Equal(t, "Test User", users[0].Username)
//
//	mockCollection.AssertExpectations(t)
//	mockCursor.AssertExpectations(t)
//}
//
//func Test_GetUserByID(t *testing.T) {
//	t.Parallel()
//
//	mockCollection := new(mocks.Collection)
//	mockSingleResult := new(mocks.SingleResult)
//
//	repo := &user.Repository{Db: mockCollection}
//
//	userID := primitive.NewObjectID()
//	expectedUser := repoModels.User{
//		ID:       userID,
//		Username: "Test User",
//		Email:    "test@example.com",
//		Age:      12,
//	}
//
//	mockCollection.On("FindOne", mock.Anything, bson.M{"_id": userID}).Return(mockSingleResult)
//
//	mockSingleResult.On("Decode", mock.AnythingOfType("*repoModels.User")).Run(func(args mock.Arguments) {
//		arg := args.Get(0).(*repoModels.User)
//		*arg = expectedUser
//	}).Return(nil)
//
//	mockSingleResult.On("Decode", mock.AnythingOfType("*repoModels.User")).Run(func(args mock.Arguments) {
//		arg := args.Get(0).(*repoModels.User)
//		*arg = expectedUser
//	}).Return(nil)
//
//	result, err := repo.GetUserByID(userID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expectedUser, result)
//	mockCollection.AssertExpectations(t)
//	mockSingleResult.AssertExpectations(t)
//}

func Test_UpdateUser(t *testing.T) {
	t.Parallel()
	mockCollection := new(mocks.Collection)
	repo := &user.Repository{Db: mockCollection}

	userID := primitive.NewObjectID()
	mockUser := repoModels.User{
		ID:       userID,
		Username: "Updated User",
		Email:    "updated@example.com",
	}

	mockCollection.On("ReplaceOne", mock.Anything, bson.M{"_id": userID}, mockUser).Return(&mongo.UpdateResult{}, nil)

	err := repo.UpdateUser(userID, mockUser)
	assert.NoError(t, err)

	mockCollection.AssertExpectations(t)
}

func normalizeTimestamp(ts time.Time) time.Time {
	return ts.Truncate(time.Second) // Adjust precision as needed
}

func Test_DeleteUser(t *testing.T) {
	t.Parallel()
	mockCollection := new(mocks.Collection)
	repo := &user.Repository{Db: mockCollection}

	userID := primitive.NewObjectID()
	now := normalizeTimestamp(time.Now())

	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": userID}, mock.Anything).Return(&mongo.UpdateResult{}, nil)

	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": userID}, mock.Anything).Run(func(args mock.Arguments) {
		updateDoc := args.Get(2).(bson.M)
		setField, ok := updateDoc["$set"].(bson.M)
		assert.True(t, ok)
		deletedAt, ok := setField["deleted_at"].(time.Time)
		assert.True(t, ok)
		// Normalize and assert as time changes while the code is run
		assert.Equal(t, normalizeTimestamp(now), normalizeTimestamp(deletedAt))
	}).Return(&mongo.UpdateResult{}, nil)

	err := repo.DeleteUser(userID)
	assert.NoError(t, err)

	mockCollection.AssertExpectations(t)
}
