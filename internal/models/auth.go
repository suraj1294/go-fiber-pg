package models

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type JwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *JwtUser) (TokenPairs, error) {
	//Create Token and set claims
	signedAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"sub":  fmt.Sprint(user.ID),
		"iss":  j.Issuer,
		"iat":  time.Now().UTC().Unix(),
		"typ":  "JWT",
		//set the expiry for jwt
		"exp": time.Now().UTC().Add(j.TokenExpiry).Unix(),
	}).SignedString([]byte(j.Secret))

	if err != nil {
		return TokenPairs{}, err
	}

	//Create a refresh token and set claims
	signedRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprint(user.ID),
		"iat": time.Now().UTC().Unix(),
		//Set the expiry for the expiry token
		"exp": time.Now().UTC().Add(j.RefreshExpiry).Unix(),
	}).SignedString([]byte(j.Secret))

	if err != nil {
		return TokenPairs{}, err
	}

	//Create TokenPairs and populate with signed tokens
	tokenPairs := TokenPairs{Token: signedAccessToken, RefreshToken: signedRefreshToken}

	//return token pairs
	return tokenPairs, nil

}

func (j *Auth) GetRefreshCookie(refreshToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: "Lax",
		//Domain:   j.CookieDomain,
		HTTPOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: "Lax",
		//Domain:   j.CookieDomain,
		HTTPOnly: true,
		Secure:   true,
	}
}
