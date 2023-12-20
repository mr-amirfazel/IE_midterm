package handlers

import (
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/shopping-basket/db"
	"github.com/mr-amirfazel/shopping-basket/handlers/auth"
	"github.com/mr-amirfazel/shopping-basket/models"
	"gorm.io/datatypes"
)

func GetBaskets(c echo.Context) error {
	_, err := auth.ExtractUserIDFromToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	var baskets []models.Basket
	if err := db.DB.Find(&baskets).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, baskets)
}

func CreateBasket(c echo.Context) error {
	userID, err := auth.ExtractUserIDFromToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	basket := new(models.Basket)
	if err := c.Bind(basket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	basket.UserID = userID

	if err := db.DB.Create(&basket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, basket)
}

func GetBasketByID(c echo.Context) error {
	_, err := auth.ExtractUserIDFromToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var basket models.Basket
	if err := db.DB.First(&basket, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}
	return c.JSON(http.StatusOK, basket)
}

func UpdateBasket(c echo.Context) error {
	_, err := auth.ExtractUserIDFromToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var basket models.Basket
	if err := db.DB.First(&basket, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	if err := c.Bind(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := db.DB.Save(&basket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, basket)
}

func DeleteBasket(c echo.Context) error {
	_, err := auth.ExtractUserIDFromToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var basket models.Basket
	if err := db.DB.First(&basket, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	if err := db.DB.Delete(&basket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Item with id: " + strconv.Itoa(id) + " Deleted from basket Successfully"})
}

func StoreJsonData(baskets []models.Basket, jsonData []datatypes.JSON) []map[string]interface{} {
	var inInterfaces []map[string]interface{}

	if len(baskets) == 1 {
		var inInterface map[string]interface{}

		inrec, _ := json.Marshal(baskets[0])
		json.Unmarshal(inrec, &inInterface)
		inInterface["json_data"] = jsonData[0]

		inInterfaces = append(inInterfaces, inInterface)
	} else {
		for i, _ := range baskets {
			var inInterface map[string]interface{}

			inrec, _ := json.Marshal(baskets[i])
			json.Unmarshal(inrec, &inInterface)

			inInterface["json_data"] = jsonData[i]
			inInterfaces = append(inInterfaces, inInterface)
		}
	}
	return inInterfaces
}
