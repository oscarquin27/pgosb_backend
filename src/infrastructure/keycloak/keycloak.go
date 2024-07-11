package keycloak

import (
	"context"
	"fdms/infra/config"
	"fmt"
	"time"

	"github.com/Nerzal/gocloak/v13"
)

var keycloakConfig *config.Keycloak

type KeycloakAuthenticationService struct {
	client *gocloak.GoCloak
}

func NewService(c *gocloak.GoCloak) *KeycloakAuthenticationService {

	return &KeycloakAuthenticationService{client: c}
}

// ... (Custom error types remain the same) ...

// 1. Create User
func (s *KeycloakAuthenticationService) CreateUser(ctx context.Context, username, email, userID, password string) (string, error) {

	token, err := s.client.LoginAdmin(ctx,
		keycloakConfig.AdminUser,
		keycloakConfig.AdminPassword,
		"master",
	)
	if err != nil {
		return "", fmt.Errorf("error getting token: %w", err)
	}

	user := gocloak.User{
		Username:    gocloak.StringP(username),
		Email:       gocloak.StringP(email), // Email can be nil
		Enabled:     gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{{Type: gocloak.StringP("password"), Value: gocloak.StringP(password)}},
		Attributes:  &map[string][]string{"pgosb_id": {userID}},
	}

	userId, err := s.client.CreateUser(ctx, token.AccessToken, keycloakConfig.Realm, user)

	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	return userId, nil
}

// 2. Inspect Token
func (s *KeycloakAuthenticationService) InspectToken(ctx context.Context, accessToken string) error {

	result, err := s.client.RetrospectToken(ctx, accessToken,
		keycloakConfig.ClientId,
		keycloakConfig.ClientSecret,
		keycloakConfig.Realm) // Using master realm for introspection

	fmt.Println("Result Instropect", result)

	if err != nil {
		return err
	}

	if result == nil {
		return fmt.Errorf("no se obtuvo result")
	}

	if !*result.Active {

		if result.Exp != nil && int64(*result.Exp) < time.Now().Unix() {
			return fmt.Errorf("token expirado")
		}
		return fmt.Errorf("invalid or inactive token: %w", err)
	}

	return nil
}

// 3. Login User
func (s *KeycloakAuthenticationService) LoginUser(ctx context.Context, username, password string) (*gocloak.JWT, error) {
	token, err := s.client.Login(
		ctx,
		keycloakConfig.ClientId,
		keycloakConfig.ClientSecret,
		keycloakConfig.Realm,
		username,
		password,
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// 3. Logout User
func (s *KeycloakAuthenticationService) LogOutUser(ctx context.Context, sessionId string) error {

	token, err := s.client.LoginAdmin(ctx,
		keycloakConfig.AdminUser,
		keycloakConfig.AdminPassword,
		"master",
	)
	if err != nil {
		return fmt.Errorf("error getting token: %w", err)
	}

	err = s.client.LogoutUserSession(
		ctx,
		token.AccessToken,
		keycloakConfig.Realm,
		sessionId)

	if err != nil {
		return err
	}

	return nil
}

// 4. Refresh Token
func (s *KeycloakAuthenticationService) RefreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error) {
	token, err := s.client.RefreshToken(
		ctx,
		refreshToken,
		keycloakConfig.ClientId,
		keycloakConfig.ClientSecret,
		keycloakConfig.Realm,
	)
	if err != nil {
		return nil, fmt.Errorf("error refreshing token: %w", err)
	}

	return token, nil
}

func (s *KeycloakAuthenticationService) GetCerts(ctx context.Context) (*gocloak.CertResponse, error) {

	return s.client.GetCerts(ctx, keycloakConfig.Realm)
}

func (s *KeycloakAuthenticationService) DeleteUser(ctx context.Context, userId string) error {

	token, err := s.client.LoginAdmin(ctx,
		keycloakConfig.AdminUser,
		keycloakConfig.AdminPassword,
		"master",
	)

	if err != nil {
		return fmt.Errorf("error getting token: %w", err)
	}

	err = s.client.DeleteUser(ctx, token.AccessToken, "pgosb", userId)

	return err
}

func init() {
	fmt.Println("Inicio Package Authentication")
	keycloakConfig = &config.Get().Keycloak
}
