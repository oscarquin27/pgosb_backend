package auth_handlers

import (
	"fdms/src/infrastructure/config"
	"fdms/src/infrastructure/keycloak"
	logger "fdms/src/infrastructure/log"
	"fdms/src/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// const PGOSB_USER_ID_COOKIE string = "PGOSB_USER_ID_COOKIE"
// const PGOSB_EXPIRES_IN_COOKIE = "PGOSB_EXPIRES_IN"
// const PGOSB_EXPIRES_DATE = "PGOSB_EXPIRES_DATE"

const secureCookie = true
const SameSiteConfig = http.SameSiteNoneMode

type UserIdentification struct {
	UserId         string    `json:"user_id"`
	ExpirationDate time.Time `json:"expire_in_date"`
	ExpiresInValue int       `json:"expire_in"`
}

func SetCookies(jwtToken string, refreshJwt string, sessionState string, expiresIn int, c *gin.Context) {

	// token, err := ReadJwt(jwtToken)

	// if err != nil {
	// 	logger.Error().Err(err).Msg("no se pudo leer el token")
	// 	return

	// }
	// claims, ok := token.Claims.(jwt.MapClaims)

	// if !ok {
	// 	logger.Error().Msg("no se pudieron obtener los claims del token")
	// 	return
	// }

	// userId, ok := claims["pgosb_id"].(string)

	// if !ok {
	// 	logger.Error().Msg("no se pudo obtener el user id")
	// 	return
	// }

	// loc, err := time.LoadLocation("America/Caracas")

	// if err != nil {
	// 	logger.Error().Err(err).Msg("no se puedo colocar zona horaria")
	// 	return

	// }

	c.SetCookie(utils.PGOSB_ACCESS_TOKEN_COOKIE, jwtToken, expiresIn, "/", config.Get().Http.MainDomain, secureCookie, true)
	c.SetCookie(utils.PGOSB_REFRESH_TOKEN_COOKIE, refreshJwt, expiresIn, "/", config.Get().Http.MainDomain, secureCookie, true)
	c.SetCookie(utils.PGOSB_SESSION_STATE_COOKIE, sessionState, expiresIn, "/", config.Get().Http.MainDomain, secureCookie, true)

	// c.SetCookie(PGOSB_USER_ID_COOKIE, userId, expiresIn, "/", config.Get().Http.MainDomain, secureCookie, false)
	// c.SetCookie(PGOSB_EXPIRES_IN_COOKIE, strconv.Itoa(expiresIn), expiresIn, "/", config.Get().Http.MainDomain, secureCookie, false)
	// expiresDate := time.Now().In(loc).Add(time.Duration(expiresIn) * time.Second)
	// c.SetCookie(PGOSB_EXPIRES_DATE,
	// 	expiresDate.String(), expiresIn, "/", config.Get().Http.MainDomain, secureCookie, false)

}

func ClearCookies(c *gin.Context) {
	c.SetCookie(utils.PGOSB_ACCESS_TOKEN_COOKIE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, true)
	c.SetCookie(utils.PGOSB_REFRESH_TOKEN_COOKIE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, true)
	c.SetCookie(utils.PGOSB_SESSION_STATE_COOKIE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, true)

	// c.SetCookie(PGOSB_USER_ID_COOKIE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, false)
	// c.SetCookie(PGOSB_EXPIRES_IN_COOKIE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, false)
	// c.SetCookie(PGOSB_EXPIRES_DATE, "", 0, "/", config.Get().Http.MainDomain, secureCookie, false)

}

type AuthController struct {
	authService *keycloak.KeycloakAuthenticationService
}

func NewAuthController(s *keycloak.KeycloakAuthenticationService) *AuthController {
	return &AuthController{
		authService: s,
	}
}

func GetUserData(tokenString string, expiresIn int) (*UserIdentification, error) {

	jwtToken, err := utils.ReadJwt(tokenString)

	if err != nil {
		logger.Warn().Err(err).Msg("no se pudo leer el token de la cookie")
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {

		return nil, fmt.Errorf("no se pudo leer el id de los claims del usuario")
	}

	userId, ok := claims["pgosb_id"].(string)

	if !ok {
		logger.Error().Msg("no se pudo obtener el user id")
		return nil, fmt.Errorf("no se pudo leer el id de los claims del usuario")
	}

	loc, _ := time.LoadLocation("America/Caracas")

	userIdentification := UserIdentification{
		UserId:         userId,
		ExpirationDate: time.Now().In(loc).Add(time.Duration(expiresIn) * time.Second),
		ExpiresInValue: expiresIn,
	}

	return &userIdentification, nil
}

func (controller *AuthController) Login(c *gin.Context) {

	var loginData LoginData
	err := c.ShouldBindBodyWithJSON(&loginData)

	if err != nil {
		errMessage := LoginErr{
			Code:    "BadCredendials",
			Message: err.Error(),
		}
		ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	jwt, err := controller.authService.LoginUser(c.Request.Context(), loginData.UserName, loginData.Password)

	if err != nil {
		errMessage := LoginErr{
			Code:    "NotJWT",
			Message: err.Error(),
		}
		ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	// loginResponse := LoginResponse{
	// 	AccessToken:  jwt.AccessToken,
	// 	RefreshToken: jwt.RefreshToken,
	// 	SessionState: jwt.SessionState,
	// 	ExpiresIn:    jwt.ExpiresIn,
	// }

	userData, err := GetUserData(jwt.AccessToken, jwt.ExpiresIn)

	if err != nil {

		c.Status(500)

		return
	}

	c.SetSameSite(SameSiteConfig)
	SetCookies(jwt.AccessToken, jwt.RefreshToken, jwt.SessionState, jwt.ExpiresIn, c)

	c.JSON(200, userData)

}

func (controller *AuthController) LogOut(c *gin.Context) {

	sessionId, err := c.Cookie(utils.PGOSB_SESSION_STATE_COOKIE)

	if err != nil {
		logger.Warn().Err(err).Msg("no se obtuvo token durante el login")
		ClearCookies(c)
		c.Status(http.StatusOK)
		return
	}

	err = controller.authService.LogOutUser(c.Request.Context(), sessionId)

	if err != nil {

		logger.Error().Err(err).Msg("no se obtuvo token durante el login")

		c.Status(500)

		return
	}

	c.SetSameSite(SameSiteConfig)
	ClearCookies(c)
	c.Status(http.StatusOK)

}

func (controller *AuthController) RefreshSession(c *gin.Context) {

	refreshToken, err := c.Cookie(utils.PGOSB_REFRESH_TOKEN_COOKIE)

	if err != nil {
		fmt.Println(err)
		ClearCookies(c)
		c.Status(http.StatusBadRequest)
		return
	}

	jwt, err := controller.authService.RefreshToken(c.Request.Context(), refreshToken)

	if err != nil {
		errMessage := LoginErr{
			Code:    "NotJWT",
			Message: err.Error(),
		}
		ClearCookies(c)
		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	userData, err := GetUserData(jwt.AccessToken, jwt.ExpiresIn)

	if err != nil {

		c.Status(500)

		return
	}
	c.SetSameSite(SameSiteConfig)

	SetCookies(jwt.AccessToken, jwt.RefreshToken, jwt.SessionState, jwt.ExpiresIn, c)

	c.JSON(200, userData)
}

func (controller *AuthController) LoginTest(c *gin.Context) {

	// one, _ := c.Cookie(PGOSB_ACCESS_TOKEN_COOKIE)
	// two, _ := c.Cookie(PGOSB_REFRESH_TOKEN_COOKIE)
	// three, _ := c.Cookie(PGOSB_SESSION_STATE_COOKIE)

	// four, _ := c.Cookie(PGOSB_USER_ID_COOKIE)
	// five, _ := c.Cookie(PGOSB_EXPIRES_DATE)
	// six, _ := c.Cookie(PGOSB_EXPIRES_IN_COOKIE)

	// fmt.Println(one, two, three, four, five, six)

	c.Status(200)
}
