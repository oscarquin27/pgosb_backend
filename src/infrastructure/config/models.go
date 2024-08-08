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
	MainDomain string `yaml:"main_domain"`
	Port       int    `yaml:"port"`
	EnabledSsl bool   `yaml:"enable_ssl"`
	SslCert    string `yaml:"ssl_cert"`
	SslKey     string `yaml:"ssl_key"`
}

type LogSettings struct {
	Console           bool   `yaml:"console"`
	BeutifyConsoleLog bool   `yaml:"beutify_console"`
	File              bool   `yaml:"file"`
	Ruta              string `yaml:"route"`
	MinLevel          string `yaml:"min_level"`
	RotationMaxSizeMB int    `yaml:"rotation_max_size_mb"`
	MaxAgeDay         int    `yaml:"max_age_day"`
	MaxBackups        int    `yaml:"max_backups"`
	Compress          bool   `yaml:"compress"`
}

type Config struct {
	Database    Database    `yaml:"database"`
	Keycloak    Keycloak    `yaml:"keycloak"`
	Http        Http        `yaml:"http"`
	LogSettings LogSettings `yaml:"log_settings"`
}
