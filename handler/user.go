package handler

import (
	"bileygr/components"
	"bileygr/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(ctx echo.Context) error {
	username := ctx.Param("username")
	if username == "" {
		return ctx.JSON(http.StatusBadGateway, map[string]string{
			"error": "username is required",
		})
	}

	var userID string
	err := db.DevDB.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "user not found",
		})
	}

	tokenUserID := ctx.Get("user_id").(string)
	if userID != tokenUserID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "unauthorized access",
		})
	}
	return Render(ctx, http.StatusOK, components.Profile(username))
}
func GetUserInfo(ctx echo.Context) error {
	username := ctx.Param("username")
	if username == "" {
		return ctx.JSON(http.StatusBadGateway, map[string]string{
			"error": "username is required",
		})
	}

	var userID string
	var email sql.NullString
	var profileImage sql.NullString
	err := db.DevDB.QueryRow(`
	SELECT id, email, profile_image 
	FROM users 
	WHERE username = $1`, username).Scan(&userID, &email, &profileImage)

	if err != nil {
		log.Printf("GetUserInfo: Database error: %v", err)
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "user not found",
		})
	}

	tokenUserID := ctx.Get("user_id").(string)
	if userID != tokenUserID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "unauthorized access",
		})
	}
	emailDisplay := "Not provided"
	if email.Valid {
		emailDisplay = email.String
	}

	var profileImageHTML string
	if profileImage.Valid && profileImage.String != "" {
		profileImageHTML = fmt.Sprintf(`
            <div class="mt-4">
                <img src="%s" alt="Profile" class="w-32 h-32 rounded-full"/>
            </div>
        `, profileImage.String)
	}

	return ctx.HTML(http.StatusOK, fmt.Sprintf(`
        <div class="bg-white shadow rounded-lg p-6">
            <div class="space-y-4">
                <div class="flex items-center space-x-2">
                    <span class="font-semibold text-gray-700">Username:</span>
                    <span>%s</span>
                </div>
                <div class="flex items-center space-x-2">
                    <span class="font-semibold text-gray-700">Email:</span>
                    <span>%s</span>
                </div>
                %s
            </div>
        </div>
    `, username, emailDisplay, profileImageHTML))
}

func UpdateUser(ctx echo.Context) error {
	return errors.New("error not implemented update users function")
}

func DeleteUser(ctx echo.Context) error {
	return errors.New("error not implemented delete users function")
}
