package auth_handlers

import (
	"fdms/src/infrastructure/keycloak"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CookieDomain = "localhost"

const PGOSB_ACCESS_TOKEN_COOKIE string = "PGOSB_ACCESS_TOKEN"
const PGOSB_REFRESH_TOKEN_COOKIE string = "PGOSB_REFRESH_TOKEN"
const PGOSB_SESSION_STATE_COOKIE string = "PGOSB_SESSION_STATE"

func SetCookies(jwt string, refreshJwt string, sessionState string, expiresIn int, c *gin.Context) {
	c.SetCookie(PGOSB_ACCESS_TOKEN_COOKIE, jwt, expiresIn, "/", CookieDomain, false, true)
	c.SetCookie(PGOSB_REFRESH_TOKEN_COOKIE, refreshJwt, expiresIn, "/", CookieDomain, false, true)
	c.SetCookie(PGOSB_SESSION_STATE_COOKIE, sessionState, expiresIn, "/", CookieDomain, false, true)
}

func ClearCookies(c *gin.Context) {
	// c.SetCookie(PGOSB_ACCESS_TOKEN_COOKIE, "", 0, "/", CookieDomain, false, false)
	// c.SetCookie(PGOSB_REFRESH_TOKEN_COOKIE, "", 0, "/", CookieDomain, false, false)
	// c.SetCookie(PGOSB_SESSION_STATE_COOKIE, "", 0, "/", CookieDomain, false, false)
}

type AuthController struct {
	authService *keycloak.KeycloakAuthenticationService
}

func NewAuthController(s *keycloak.KeycloakAuthenticationService) *AuthController {
	return &AuthController{
		authService: s,
	}
}

func (controller *AuthController) Login(c *gin.Context) {

	var loginData LoginData
	err := c.ShouldBindBodyWithJSON(&loginData)

	if err != nil {
		errMessage := LoginErr{
			Code:    "BadCredendials",
			Message: err.Error(),
		}
		//ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	jwt, err := controller.authService.LoginUser(c.Request.Context(), loginData.UserName, loginData.Password)

	if err != nil {
		errMessage := LoginErr{
			Code:    "NotJWT",
			Message: err.Error(),
		}
		// ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	loginResponse := LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		SessionState: jwt.SessionState,
		ExpiresIn:    jwt.ExpiresIn,
	}

	// SetCookies(jwt.AccessToken, jwt.RefreshToken, jwt.SessionState, jwt.ExpiresIn, c)
	c.JSON(http.StatusOK, loginResponse)

}

func (controller *AuthController) LogOut(c *gin.Context) {

	sessionId, err := c.Cookie(PGOSB_SESSION_STATE_COOKIE)

	if err != nil {
		fmt.Println(err)
		//ClearCookies(c)
		c.Status(http.StatusOK)
		return
	}

	err = controller.authService.LogOutUser(c.Request.Context(), sessionId)
	fmt.Println(err)

	// ClearCookies(c)
	c.Status(http.StatusOK)

}

func (controller *AuthController) RefreshSession(c *gin.Context) {

	refreshToken, err := c.Cookie(PGOSB_REFRESH_TOKEN_COOKIE)

	if err != nil {
		fmt.Println(err)
		//ClearCookies(c)
		c.Status(http.StatusBadRequest)
		return
	}

	jwt, err := controller.authService.RefreshToken(c.Request.Context(), refreshToken)

	if err != nil {
		errMessage := LoginErr{
			Code:    "NotJWT",
			Message: err.Error(),
		}
		//ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	//SetCookies(jwt.AccessToken, jwt.RefreshToken, jwt.SessionState, jwt.ExpiresIn, c)

	loginResponse := LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		SessionState: jwt.SessionState,
		ExpiresIn:    jwt.ExpiresIn,
	}

	// SetCookies(jwt.AccessToken, jwt.RefreshToken, jwt.SessionState, jwt.ExpiresIn, c)
	c.JSON(http.StatusOK, loginResponse)
}
