package main

import (
	"backend/cmd/api/config"
	"backend/cmd/api/infrastructure/persistence"
	"backend/cmd/api/interface/handler"
	"backend/cmd/api/interface/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	diaryPersistence := persistence.NewDiaryPersistence(db)
	diaryUseCase := usecase.NewDiaryUseCase(diaryPersistence, gamePersistence)
	diaryHandler := handler.NewDiaryHandler(diaryUseCase, gameUseCase)

	r.POST("/startplaying", diaryHandler.StartPlaying)
	r.POST("/endplaying", diaryHandler.EndPlaying)

	r.Run(":8000")
}
