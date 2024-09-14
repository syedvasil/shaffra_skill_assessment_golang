package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email" gorm:"unique;not null;index"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
	Age      int                `bson:"age" json:"age"`
	//Role      string             `bson:"role" json:"role"`
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type UserUpdateReq struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty" binding:"email"`
	Age      int    `bson:"age,omitempty" json:"age,omitempty" binding:"gte=0,lte=130"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Password string `bson:"password,omitempty" json:"-"`
}

func (u *User) ForUpdate(newUser UserUpdateReq) {
	if len(newUser.Password) > 0 {
		u.Password = newUser.Password
	}
	if newUser.Email != newUser.Email {
		u.Email = newUser.Email
	}
	if newUser.Age != newUser.Age {
		u.Age = newUser.Age
	}
	if newUser.Username != newUser.Username {
		u.Username = newUser.Username
	}
}
