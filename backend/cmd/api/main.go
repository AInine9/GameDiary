package main

import (
	"backend/cmd/api/config"
	"backend/cmd/api/infrastructure/persistence"
	"backend/cmd/api/interface/handler"
	"backend/cmd/api/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	db := config.Connect()
	defer db.Close()

	gamePersistence := persistence.NewGamePersistence(db)
	gameUseCase := usecase.NewGameUseCase(gamePersistence)

	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := usecase.NewUserUseCase(userPersistence)

	diaryPersistence := persistence.NewDiaryPersistence(db)
	diaryUseCase := usecase.NewDiaryUseCase(diaryPersistence, gamePersistence)
	diaryHandler := handler.NewDiaryHandler(diaryUseCase, gameUseCase, userUseCase)

	authHandler := handler.NewAuthHandler(userUseCase)

	r.GET("/auth", authHandler.BeginAuth)
	r.GET("/auth/callback", authHandler.CompleteAuth)
	r.GET("/logout", authHandler.Logout)
	r.POST("/startplaying", diaryHandler.StartPlaying)
	r.POST("/endplaying", diaryHandler.EndPlaying)
	r.Run(":8000")
}

func init() {
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "discord", nil
	}
	goth.UseProviders(
		discord.New(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_SECRET"), os.Getenv("CALLBACK_URL"), discord.ScopeIdentify),
	)
}
