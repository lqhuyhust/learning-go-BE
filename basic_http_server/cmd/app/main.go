package main

import (
	"fmt"
	"httpServer/controllers"
	"httpServer/database"
	"httpServer/models"
	"httpServer/repositories"
	"httpServer/routes"
	"httpServer/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	godotenv.Load()

	// Connect to database
	database.ConnectDatabase()

	// Run automigration
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.ReactionType{},
		&models.Reaction{},
		&models.Post{},
		&models.Comment{},
		&models.RefreshToken{},
	)
	if err != nil {
		fmt.Println("Error migrating database:", err)
	}

	// Run seed
	database.SeedReactions()

	// Init repositories
	userRepository := repositories.NewUserRepository(database.DB)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(database.DB)
	postRepository := repositories.NewPostRepository(database.DB)
	reactionRepository := repositories.NewReactionRepository(database.DB)
	commentRepository := repositories.NewCommentRepository(database.DB)

	// Init services
	authService := services.NewAuthService(userRepository, refreshTokenRepository)
	userService := services.NewUserService(userRepository)
	postService := services.NewPostService(postRepository)
	reactionService := services.NewReactionService(reactionRepository)
	commentService := services.NewCommentService(commentRepository)

	// Init controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	postController := controllers.NewPostController(postService)
	reactionController := controllers.NewReactionController(reactionService)
	commentController := controllers.NewCommentController(commentService)

	// Init routes
	router := routes.InitRoutes(
		userController,
		authController,
		postController,
		reactionController,
		commentController,
	)

	fmt.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
