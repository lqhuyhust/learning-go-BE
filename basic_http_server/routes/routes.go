package routes

import (
	"httpServer/controllers"
	"httpServer/middleware"
	"httpServer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userController *controllers.UserController,
	authController *controllers.AuthController,
	postController *controllers.PostController,
	reactionController *controllers.ReactionController,
	commentController *controllers.CommentController,
) *gin.Engine {
	router := gin.Default()
	authService := &services.AuthService{}

	// API Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// authentication routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", middleware.AuthMiddleware(authService), authController.Logout)
		authRoutes.POST("/refresh-token", authController.RefreshToken)
	}

	// user routes
	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(authService))
	{
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
	}

	// post routes
	postRoutes := router.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware(authService))
	{
		postRoutes.GET("/:id", postController.ShowPostByID)
		postRoutes.POST("/", postController.CreatePost)
		postRoutes.PUT("/:id", postController.UpdatePostByID)
		postRoutes.DELETE("/:id", postController.DeletePostByID)
	}

	// reaction routes
	reactionRoutes := router.Group("/reaction")
	reactionRoutes.Use(middleware.AuthMiddleware(authService))
	{
		reactionRoutes.POST("/", reactionController.MakeReaction)
		reactionRoutes.DELETE("/:id", reactionController.DeleteReaction)
	}

	// comment routes
	commentRoutes := router.Group("/comment")
	commentRoutes.Use(middleware.AuthMiddleware(authService))
	{
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.GET("/", commentController.GetComments)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
	}

	return router
}
