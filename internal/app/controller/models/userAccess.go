package models

import (
	"time"

	repoModels "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email" binding:"email"`
	Age      int                `bson:"age" json:"age" binding:"gte=0,lte=130"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
}

func CreateUserFromReq(req User) repoModels.User {
	now := time.Now()
	return repoModels.User{
		ID:        primitive.NewObjectID(),
		Username:  req.Username,
		Password:  req.Password,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
	}
}
