package auth

import (
	"context"
	"go-api/config"
	"net/http"

	"gorm.io/gorm"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/labstack/echo/v4"
)

type FirebaseAuthService struct {
	db *gorm.DB
	appConfig *config.Config

}

func NewFirebaseAuth(db *gorm.DB) AuthService {
	return &FirebaseAuthService{db: db, appConfig: config.GetConfig()}
}

func (auth *FirebaseAuthService) AuthenticateWithThirdParty(
	idToken string,
) (interface{}) {
	ctx := context.Background()
	conf := &firebase.Config{
	    ProjectID: auth.appConfig.Firebase.ProjectId,
	}
	opt := option.WithCredentialsFile(auth.appConfig.Firebase.Credentials)
	app, err := firebase.NewApp(ctx, conf, opt)

	if err != nil {
	    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	client, err := app.Auth(ctx)

	if err != nil {
	    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data, err := client.VerifyIDToken(ctx, idToken)

	if err != nil {
	    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return data
}
