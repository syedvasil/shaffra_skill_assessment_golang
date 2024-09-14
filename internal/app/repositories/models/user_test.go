package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUser_ForUpdate(t *testing.T) {
	t.Parallel()

	var (
		id        = primitive.ObjectID{}
		Email     = "A@B.IO"
		Username  = "asd"
		Password  = "nopwd"
		Age       = 20
		CreatedAt = time.Now()
	)
	type fields struct {
		ID        primitive.ObjectID
		Email     string
		Username  string
		Password  string
		Age       int
		CreatedAt time.Time
		DeletedAt *time.Time
	}
	type args struct {
		newUser UserUpdateReq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				ID:        id,
				Email:     Email,
				Username:  Username,
				Password:  Password,
				Age:       Age,
				CreatedAt: CreatedAt,
				DeletedAt: nil,
			},
			args: args{
				newUser: UserUpdateReq{
					Email:    Email,
					Username: Username,
					Password: Password,
					Age:      Age,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Email:     tt.fields.Email,
				Username:  tt.fields.Username,
				Password:  tt.fields.Password,
				Age:       tt.fields.Age,
				CreatedAt: tt.fields.CreatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			u.ForUpdate(tt.args.newUser)
		})
	}
}
