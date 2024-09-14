package main

import (
	"github.com/gin-gonic/gin"
	ctrlUser "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/controller/user"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/user"
	srvUser "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/service/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupV1UserRoutes(db *mongo.Database, routerGroup *gin.RouterGroup) {
	userCtrl := ctrlUser.New(srvUser.New(user.New(db)))
	userGroup := routerGroup.Group("/user")
	{
		userGroup.POST("", userCtrl.CreateUser)
		userGroup.GET("", userCtrl.GetUsers)
		userGroup.GET("/:id", userCtrl.GetUser)
		userGroup.PUT("/:id", userCtrl.UpdateUser)
		userGroup.DELETE("/:id", userCtrl.DeleteUser)
	}
}
