package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/shopping-basket/db"
	"github.com/mr-amirfazel/shopping-basket/handlers/auth"
	"github.com/mr-amirfazel/shopping-basket/models"
)

func GetAllUsers(c echo.Context) error {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	type LoginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	login := new(LoginInfo)
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	var user models.User
	if err := db.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	if user.Password != login.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	token, err := auth.GenerateToken(user.UserID) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
func DeleteUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func ChangePassword(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	type PasswordUpdate struct {
		Password string `json:"password"`
	}

	update := new(PasswordUpdate)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if update.Password != "" {
		user.Password = update.Password

		if err := db.DB.Save(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, user)
}
