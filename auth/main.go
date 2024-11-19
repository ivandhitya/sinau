package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/ivandhitya/sinau/auth/docs"
	"github.com/ivandhitya/sinau/auth/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// generateToken menghasilkan token JWT untuk autentikasi
// @Summary      Generate JWT Token
// @Description  Menghasilkan token JWT dengan klaim username, role, dan waktu kedaluwarsa (expiry) yang dapat digunakan untuk autentikasi pada endpoint yang dilindungi
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string "token" "Token JWT berhasil dihasilkan"
// @Failure      500 {object} map[string]string "message" "Could not generate token"
// @Router       /login [get]
func generateToken(c echo.Context) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "user123",
		"role":     []string{"admin"},
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

// @Summary     Get Salary Data
// @Description Ini adalah endpoint untuk mendapatkan data salary yang hanya bisa diakses dengan token JWT
// @Tags        Secure
// @Security    ApiKeyAuth
// @Produce     json
// @Success     200 {string} string "Data aman"
// @Failure     401 {string} string "Token tidak valid atau tidak disediakan"
// @Router      /salary [get]
func salaryData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]int{"salary": 40000000})
}

// @title           Contoh Auth API Menggunakan
// @version         1.0
// @description     Ini adalah contoh implementasi autentikasi menggunakan Echo dan JWT.

// @contact.name   Support Team
// @contact.email  ivandhitya@gmail.com

// @host      localhost:7777
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	webServer := echo.New()

	webServer.GET("/login", generateToken)
	webServer.GET("/salary", salaryData, middleware.AuthMiddleware)
	webServer.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	webServer.Logger.Fatal(webServer.Start(":7777"))
}
