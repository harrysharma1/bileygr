package routes

import (
	"bileygr/db"
	"bileygr/handler"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Run(app *echo.Echo) {
	app.Static("/static", "static")
	app.GET("/", handler.Home)

	users(app)
	reading(app)
	testUserRoutes(app)

	app.Logger.Fatal(app.Start(":6969"))
}

func users(app *echo.Echo) {
	app.GET("/users/:id", handler.GetUser)
	app.POST("/users", handler.SaveUser)
	app.PUT("/users/:id", handler.UpdateUser)
	app.DELETE("/users/:id", handler.DeleteUser)
}

func reading(app *echo.Echo) {
	app.GET("/:readingType/:id", handler.GetReading)
	app.POST("/reading", handler.SaveReading)
	app.PUT("/:readingType/:id", handler.UpdateReading)
	app.DELETE("/:readingType/:id", handler.DeleteReading)
}

type Creds struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}

func testUserRoutes(app *echo.Echo) {
	app.POST("/signup", signup)
}

func signup(ctx echo.Context) error {
	creds := &Creds{}
	err := json.NewDecoder(ctx.Request().Body).Decode(creds)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request format",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "processing password",
		})
	}
	id := uuid.New()
	timestamp := time.Now()
	log.Printf("Username: %s\n Password: %s", creds.Username, string(hashedPassword))

	_, err = db.DevDB.Exec("INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)",
		id.String(), creds.Username, string(hashedPassword), timestamp.String())
	if err != nil {
		log.Printf("Database error: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating user",
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "user created successfully",
	})

}
