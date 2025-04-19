package config

type AppConfig struct {
	BackendPort     string `mapstructure:"BACKEND_PORT" validate:"required"`
	FrontendPort    string `mapstructure:"FRONTEND_PORT" validate:"required"`
	Mode            string `mapstructure:"MODE" validate:"required"`
	Domain          string `validate:"required"`
	StaticFilesPath string `mapstructure:"STATIC_FILES_PATH" validate:"required"`
	DBPath          string `mapstructure:"DB_PATH" validate:"required"`
}
