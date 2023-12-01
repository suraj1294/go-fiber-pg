package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/suraj1294/go-fiber-pg-auth/cmd/api/response"
	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
)

type AuthHandler struct {
	AppStore
}

func (ah *AuthHandler) Authenticate(c *fiber.Ctx) error {

	// read the json payload
	authUser := &models.AuthUser{}

	if err := c.BodyParser(authUser); err != nil {
		return err
	}

	//validate if user is present in db
	user, err := ah.UsersStore.GetUserByEmail(authUser.Email)

	if err != nil {
		return response.NotFoundRequestHandler(c, "invalid username or password")
	}

	fmt.Print(user.FirstName)

	//check if password is correct

	valid, err := user.PasswordMatches(authUser.Password)

	if err != nil || !valid {
		return response.NotFoundRequestHandler(c, "invalid username or password")

	}

	//create a jwt user

	u := models.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	//generate token
	tokens, err := ah.AuthStore.GenerateTokenPair(&u)

	if err != nil {
		fmt.Print(err)
		return response.UnexpectedErrorHandler(c, "error generating tokens")
	}

	refreshCookie := ah.AuthStore.GetRefreshCookie(tokens.RefreshToken)

	c.Cookie(refreshCookie)

	return response.SuccessResponseHandler(c, fiber.Map{"accessToken": tokens.Token})

}

func (rh *AuthHandler) AuthProfile(c *fiber.Ctx) error {
	refreshToken := c.Cookies(rh.AuthStore.CookieName)

	if refreshToken == "" {
		return response.UnAuthorizedRequestHandler(c, "unauthorized")
	}

	claims := &models.Claims{}

	// parse the token to get the claims
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(rh.AuthStore.Secret), nil
	})

	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unauthorized")
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unknown user")
	}

	user, err := rh.UsersStore.GetUserById(userID)

	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unknown user")
	}

	u := models.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := rh.AuthStore.GenerateTokenPair(&u)
	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "error generating tokens")
	}

	refreshCookie := rh.AuthStore.GetRefreshCookie(tokens.RefreshToken)

	c.Cookie(refreshCookie)

	return response.SuccessResponseHandler(c, fiber.Map{"accessToken": tokens.Token})

}

func (rh *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies(rh.AuthStore.CookieName)

	claims := &models.Claims{}

	// parse the token to get the claims
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(rh.AuthStore.Secret), nil
	})

	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unauthorized")
	}

	//get the user id from token claims

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unknown user")
	}

	user, err := rh.UsersStore.GetUserById(userID)

	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "unknown user")
	}

	u := models.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := rh.AuthStore.GenerateTokenPair(&u)
	if err != nil {
		return response.UnAuthorizedRequestHandler(c, "error generating tokens")
	}

	refreshCookie := rh.AuthStore.GetRefreshCookie(tokens.RefreshToken)

	c.Cookie(refreshCookie)

	return response.SuccessResponseHandler(c, fiber.Map{"accessToken": tokens.Token})

}

func (rh *AuthHandler) Logout(c *fiber.Ctx) error {

	c.Cookie(rh.AuthStore.GetExpiredRefreshCookie())

	return response.SuccessResponseHandler(c, "ok")
}
