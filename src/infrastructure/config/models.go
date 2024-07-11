package config

type Database struct {
	ConnectionString string `yaml:"connection_string"`
}
type Keycloak struct {
	Address       string `yaml:"address"`
	AdminUser     string `yaml:"admin_user"`
	AdminPassword string `yaml:"admin_password"`
	ClientId      string `yaml:"client_id"`
	ClientSecret  string `yaml:"client_secret"`
	Realm         string `yaml:"realm"`
}
type Http struct {
	Port string `yaml:"port"`
}

type Config struct {
	Database Database
	Keycloak Keycloak
	Http     Http
}
