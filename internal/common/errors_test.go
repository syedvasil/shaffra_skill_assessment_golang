package common_test

import (
	"bytes"
	"errors"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/common"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func getContext() *gin.Context {
	context, _ := createTestContext("", "/test/sa", nil)

	return context
}

func createTestContext(method, url string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	// Create a new HTTP request and response recorder
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Create a new Gin context
	c, _ := gin.CreateTestContext(w)

	// Assign the request to the context
	c.Request = req

	return c, w
}

func TestHandleError(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx *gin.Context
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "not allowed",
			args: args{
				ctx: getContext(),
				err: common.ForbiddenError,
			},
		},
		{
			name: "default",
			args: args{
				ctx: getContext(),
				err: errors.New("bla blue"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			common.HandleError(tt.args.ctx, tt.args.err)
		})
	}
}
