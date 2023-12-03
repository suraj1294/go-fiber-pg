package app

import (
	"time"

	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
)

func (app *Application) GetAppAuth() *models.Auth {

	return &models.Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecrete,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Session", //"__Host-refresh_token",
		//CookieDomain:  app.CookieDomain,
	}

}
