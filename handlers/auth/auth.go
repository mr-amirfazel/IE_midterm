package auth

import (
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "net/http"
    "strings"
    "time"
)

var jwtKey = []byte("6w5F94c9Uc17CVPuWPrZRV4V3hm7CByZ") 

func GenerateToken(userID uint) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

    signedToken, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

func ExtractToken(c echo.Context) (string, error) {
    authHeader := c.Request().Header.Get("Authorization")
    if authHeader == "" {
        return "", echo.NewHTTPError(http.StatusUnauthorized, "No JWT token found")
    }

    
    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
    }

    return parts[1], nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
        }
        return []byte(jwtKey), nil 
    })
    if err != nil {
        return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
    }

    return token, nil 
}

func ExtractUserIDFromToken(c echo.Context) (uint, error) {
    tokenString, err := ExtractToken(c)
    if err != nil {
        return 0, err
    }

    token, err := ValidateToken(tokenString)
    if err != nil {
        return 0, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID in token claims")
    }

    userID := uint(userIDFloat)
    return userID, nil
}
