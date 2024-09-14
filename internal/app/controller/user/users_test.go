package user_test

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/controller/models"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/controller/user"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/controller/user/mocks"
	repoModels "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser_Success(t *testing.T) {
	t.Parallel()

	mockService := new(mocks.Service)
	controller := user.New(mockService)

	mockObjectID := primitive.NewObjectID()
	newUser := models.User{
		ID:       mockObjectID, // Use the mock ObjectID
		Username: "John Doe",
		Email:    "johndoe@gmail.com",
		Age:      20,
	}

	matchUser := func(u repoModels.User) bool {
		return u.Username == newUser.Username &&
			u.Email == newUser.Email &&
			u.Age == newUser.Age
	}
	mockService.On("CreateUser", mock.MatchedBy(matchUser)).Return(nil)

	// Mock request and response
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{
		"email": "johndoe@gmail.com",
		"age": 20,
		"username": "John Doe"
	}`))

	controller.CreateUser(ctx)
	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertCalled(t, "CreateUser", mock.MatchedBy(matchUser))
}

func TestCreateUser_BadRequest(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`invalid json`))

	controller.CreateUser(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "CreateUser")
}

func TestGetUsers_Success(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	users := []repoModels.User{
		{
			Username: "John Doe",
			Email:    "johndoe@gmail.com",
			Age:      20,
			Password: "Asdasd",
		},
	}
	mockService.On("GetUsers", 1, 10).Return(users, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/users?page=1&limit=10", nil)

	controller.GetUsers(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertCalled(t, "GetUsers", 1, 10)
}

func TestGetUsers_InternalServerError(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	mockService.On("GetUsers", 1, 10).Return(nil, errors.New("error"))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/users?page=1&limit=10", nil)

	controller.GetUsers(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertCalled(t, "GetUsers", 1, 10)
}

func TestGetUser_Success(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	userID := primitive.NewObjectID()
	userT := repoModels.User{
		Username: "John Doe",
		Email:    "johndoe@gmail.com",
		Age:      20,
		Password: "Asdasd",
	}
	mockService.On("GetUserByID", userID).Return(userT, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID.Hex()}}

	controller.GetUser(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertCalled(t, "GetUserByID", userID)
}

func TestGetUser_InvalidID(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "invalid-id"}}

	controller.GetUser(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "GetUserByID")
}

func TestUpdateUser_Success(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	userID := primitive.NewObjectID()
	userUpdate := repoModels.UserUpdateReq{
		Username: "John Doe",
		Email:    "johndoe@gmail.com",
		Age:      20,
	}
	mockService.On("UpdateUser", userID, userUpdate).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID.Hex()}}
	ctx.Request, _ = http.NewRequest(http.MethodPut, "/users/"+userID.Hex(), bytes.NewBufferString(`{
    "Username": "John Doe",
    "Email": "johndoe@gmail.com",
    "Age": 20,
    "Password": "Asdasd"
}`))

	controller.UpdateUser(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertCalled(t, "UpdateUser", userID, userUpdate)
}

func TestUpdateUser_InvalidID(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "invalid-id"}}

	controller.UpdateUser(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "UpdateUser")
}

func TestDeleteUser_Success(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	userID := primitive.NewObjectID()
	mockService.On("DeleteUser", userID).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID.Hex()}}

	controller.DeleteUser(ctx)

	mockService.AssertCalled(t, "DeleteUser", userID)
}

func TestDeleteUser_InvalidID(t *testing.T) {
	t.Parallel()
	mockService := new(mocks.Service)
	controller := user.New(mockService)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "invalid-id"}}

	controller.DeleteUser(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "DeleteUser")
}
