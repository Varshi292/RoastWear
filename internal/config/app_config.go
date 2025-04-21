package config

type AppConfig struct {
	BackendPort     string `mapstructure:"BACKEND_PORT" validate:"required,port"`
	FrontendPort    string `mapstructure:"FRONTEND_PORT" validate:"required,port"`
	Mode            string `mapstructure:"MODE" validate:"required"`
	Domain          string `mapstructure:"DOMAIN" validate:"required"`
	StaticFilesPath string `mapstructure:"STATIC_FILES_PATH" validate:"required,filepath"`
	UserDBPath      string `mapstructure:"USER_DB_PATH" validate:"required,filepath"`
	SessionDBPath   string `mapstructure:"SESSION_DB_PATH" validate:"required,filepath"`
	UploadDBPath    string `mapstructure:"UPLOAD_DB_PATH" validate:"required,filepath"`
	CartDBPath      string `mapstructure:"CART_DB_PATH" validate:"required,filepath"`
}
