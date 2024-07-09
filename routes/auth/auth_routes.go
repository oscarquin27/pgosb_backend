package auth_routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	SessionState string `json:"session_state"`
	ExpiresIn    int    `json:"expires_in"`
}

type LoginErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

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

func Login(c *gin.Context) {

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

	jwt, err := keycloakAuthService.LoginUser(c.Request.Context(), loginData.UserName, loginData.Password)

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

func LogOut(c *gin.Context) {

	sessionId, err := c.Cookie(PGOSB_SESSION_STATE_COOKIE)

	if err != nil {
		fmt.Println(err)
		//ClearCookies(c)
		c.Status(http.StatusOK)
		return
	}

	err = keycloakAuthService.LogOutUser(c.Request.Context(), sessionId)
	fmt.Println(err)

	// ClearCookies(c)
	c.Status(http.StatusOK)

}

func RefreshSession(c *gin.Context) {

	refreshToken, err := c.Cookie(PGOSB_REFRESH_TOKEN_COOKIE)

	if err != nil {
		fmt.Println(err)
		//ClearCookies(c)
		c.Status(http.StatusBadRequest)
		return
	}

	jwt, err := keycloakAuthService.RefreshToken(c.Request.Context(), refreshToken)

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
