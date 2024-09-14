package user_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	"github.com/stretchr/testify/assert"
	repoModels "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/service/user"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/service/user/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestService_CreateUser(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userT := repoModels.User{
		ID:       primitive.NewObjectID(),
		Username: "John Doe",
	}

	mockRepo.On("CreateUser", userT).Return(nil)

	err := svc.CreateUser(userT)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_GetUsers(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	filter := bson.M{"deleted_at": bson.M{"$exists": false}}
	users := []repoModels.User{
		{ID: primitive.NewObjectID(), Username: "John"},
		{ID: primitive.NewObjectID(), Username: "Jane"},
	}

	mockRepo.On("GetUsers", filter, 0, 10).Return(users, nil)

	result, err := svc.GetUsers(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockRepo.AssertExpectations(t)
}

func TestService_GetUserByID(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userID := primitive.NewObjectID()
	userT := repoModels.User{ID: userID, Username: "John Doe"}

	mockRepo.On("GetUserByID", userID).Return(userT, nil)

	result, err := svc.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, userT, result)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateUser(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userID := primitive.NewObjectID()
	oldUser := repoModels.User{ID: userID, Username: "John Doe"}
	updateReq := repoModels.UserUpdateReq{Username: "John Updated"}

	updatedUser := oldUser
	updatedUser.ForUpdate(updateReq)

	mockRepo.On("GetUserByID", userID).Return(oldUser, nil)
	mockRepo.On("UpdateUser", userID, updatedUser).Return(nil)

	err := svc.UpdateUser(userID, updateReq)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteUser(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userID := primitive.NewObjectID()

	mockRepo.On("DeleteUser", userID).Return(nil)

	err := svc.DeleteUser(userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test for error handling when repository methods fail

func TestService_CreateUser_Error(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userT := repoModels.User{
		ID:       primitive.NewObjectID(),
		Username: "John Doe",
	}

	mockRepo.On("CreateUser", userT).Return(errors.New("failed to create userT"))

	err := svc.CreateUser(userT)

	assert.Error(t, err)
	assert.Equal(t, "failed to create userT", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestService_GetUserByID_Error(t *testing.T) {
	t.Parallel()
	mockRepo := new(mocks.Repository)
	svc := user.New(mockRepo)

	userID := primitive.NewObjectID()

	mockRepo.On("GetUserByID", userID).Return(repoModels.User{}, errors.New("user not found"))

	_, err := svc.GetUserByID(userID)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}
