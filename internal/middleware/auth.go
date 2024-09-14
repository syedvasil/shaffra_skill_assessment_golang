package middleware

//func Auth(db *mongo.Database) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		username, password, hasAuth := c.Request.BasicAuth()
//		if !hasAuth {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//
//		// Check username and password in MongoDB
//		var user repoModels.User
//		filter := bson.M{"username": username}
//		err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
//		if err != nil || user.Password != password {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//
//		// User authenticated
//		c.Set("Username", user.Username)
//		c.Set("ID", user.ID.Hex())
//		c.Next()
//	}
//}
