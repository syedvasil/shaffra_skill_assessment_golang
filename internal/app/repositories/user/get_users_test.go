package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/user"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/user/mocks"
	"go.mongodb.org/mongo-driver/bson"
)

//func Test_GetUsers_Success(t *testing.T) {
//	t.Parallel()
//
//	// Create mock objects
//	mockCollection := new(mocks.Collection)
//	mockCursor := new(mocks.Cursor)
//
//	// Define filter and repository
//	filter := bson.M{}
//	repo := &user.Repository{Db: mockCollection}
//
//	// Mock cursor behavior
//	mockCursor.On("Next", mock.Anything).Return(true).Once() // First call returns true
//	mockCursor.On("Decode", mock.AnythingOfType("*models.User")).Run(func(args mock.Arguments) {
//		// Populate the user with test data
//		userMock := args.Get(0).(*repoModels.User)
//		*userMock = repoModels.User{
//			ID:       primitive.NewObjectID(),
//			Username: "Test User",
//			Email:    "test@example.com",
//		}
//	}).Return(nil).Once()
//
//	mockCursor.On("Next", mock.Anything).Return(false).Once() // Second call returns false
//	mockCursor.On("Err").Return(nil)                          // No error from cursor
//
//	// Mock the collection's Find method
//	mockCollection.On("Find", mock.Anything, filter, mock.Anything).Return(mockCursor, nil)
//
//	// Call the GetUsers method
//	users, err := repo.GetUsers(filter, 0, 10)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Len(t, users, 1)
//	assert.Equal(t, "Test User", users[0].Username)
//
//	// Verify that expectations were met
//	mockCollection.AssertExpectations(t)
//	mockCursor.AssertExpectations(t)
//}

func Test_GetUsers_FindError(t *testing.T) {
	t.Parallel()

	mockCollection := new(mocks.Collection)
	repo := &user.Repository{Db: mockCollection}

	filter := bson.M{}
	findError := assert.AnError

	// Mock the collection's Find method to return an error
	mockCollection.On("Find", mock.Anything, filter, mock.Anything).Return(nil, findError)

	// Call the GetUsers method
	users, err := repo.GetUsers(filter, 0, 10)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, findError, err)

	// Verify expectations
	mockCollection.AssertExpectations(t)
}

//func Test_GetUsers_CursorError(t *testing.T) {
//	t.Parallel()
//
//	// Create mock objects
//	mockCollection := new(mocks.Collection)
//	mockCursor := new(mocks.Cursor)
//
//	// Define filter and repository
//	filter := bson.M{}
//	repo := &user.Repository{Db: mockCollection}
//
//	// Mock cursor behavior
//	mockCursor.On("Next", mock.Anything).Return(true).Once()                                   // First call returns true
//	mockCursor.On("Decode", mock.AnythingOfType("*models.User")).Return(assert.AnError).Once() // Error on Decode
//
//	// Mock the collection's Find method
//	mockCollection.On("Find", mock.Anything, filter, mock.Anything).Return(mockCursor, nil)
//
//	// Call the GetUsers method
//	users, err := repo.GetUsers(filter, 0, 10)
//
//	// Assertions
//	assert.Error(t, err)
//	assert.Nil(t, users)
//
//	// Verify expectations
//	mockCollection.AssertExpectations(t)
//	mockCursor.AssertExpectations(t)
//}
