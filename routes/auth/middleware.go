package auth_routes

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	roles "fdms/domain/roles"
	users "fdms/domain/users"
	"fdms/infra/config"
	authentication "fdms/infra/keycloak"
	"fdms/routes/auth/permission"
	"fdms/utils"
	"strings"

	"fmt"
	"math/big"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var keycloakAuthService authentication.KeycloakAuthenticationService
var keycloakClient gocloak.GoCloak

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		accessTokenCookie, err := c.Request.Cookie("PGOSB_ACCESS_TOKEN")

		fmt.Println("TOKEN:", accessTokenCookie)

		if err != nil || accessTokenCookie.Value == "" {
			fmt.Println("No se pudo obtener cookie de token", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		sessionState, err := c.Request.Cookie("PGOSB_SESSION_STATE")

		if err != nil {
			fmt.Println("No se pudo obtener cookie de session", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx := context.TODO()

		err = keycloakAuthService.InspectToken(ctx, accessTokenCookie.Value)

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token not pass Inspect"})
			return
		}

		certs, err := keycloakAuthService.GetCerts(ctx)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem Getting Certs"})
			return
		}

		token, err := jwt.Parse(accessTokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, errors.New("expecting JWT header to have a string kid")
			}

			for _, key := range *certs.Keys {
				if *key.Kid == kid {
					// Decode modulus and exponent from base64
					nBytes, err := base64.RawURLEncoding.DecodeString(*key.N)
					if err != nil {
						return nil, fmt.Errorf("error decoding modulus: %v", err)
					}
					eBytes, err := base64.RawURLEncoding.DecodeString(*key.E)
					if err != nil {
						return nil, fmt.Errorf("error decoding exponent: %v", err)
					}

					// Construct the RSA public key
					publicKey := &rsa.PublicKey{
						N: new(big.Int).SetBytes(nBytes),
						E: int(new(big.Int).SetBytes(eBytes).Int64()),
					}

					return publicKey, nil
				}
			}

			return nil, errors.New("unable to find matching key")
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		fmt.Println(token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			userId := claims["pgosb_id"].(string)

			c.Set("access_token", accessTokenCookie.Value)
			c.Set("session_state", sessionState.Value)
			c.Set("user_id", userId)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}

func PermissionAuthMiddleware(moduleName string, perm string,
	userService users.UserRepository,
	roleService roles.RoleRepository,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		userId := strings.TrimSpace(c.GetString("user_id"))

		if userId == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		id := utils.ParseInt(userId)

		user, err := userService.GetUser(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		rolSchema, err := roleService.GetRoleSchema(int64(user.Id_role))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		userPermissions, err := permission.UserPermissionFromJSONString(*rolSchema)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if value, ok := userPermissions[moduleName]; ok {

			if perm == permission.Read {
				c.Next()
				return
			}

			if (value.Write && perm == permission.Write) ||
				(value.Update && perm == permission.Update) ||
				(value.Delete && perm == permission.Delete) ||
				(value.Export && perm == permission.Export) ||
				(value.Print && perm == permission.Print) {
				c.Next()
				return
			}

		}

		c.Status(http.StatusUnauthorized)
	}
}

func init() {
	keycloakClient = *gocloak.NewClient(config.Get().Keycloak.Address)
	keycloakAuthService = *authentication.NewService(&keycloakClient)
}
