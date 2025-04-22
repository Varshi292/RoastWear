package config

type AppConfig struct {
	BackendPort     string `mapstructure:"BACKEND_PORT" validate:"required"`
	FrontendPort    string `mapstructure:"FRONTEND_PORT" validate:"required"`
	Mode            string `mapstructure:"MODE" validate:"required"`
	Domain          string `validate:"required"`
	StaticFilesPath string `mapstructure:"STATIC_FILES_PATH" validate:"required"`
	UserDBPath      string `mapstructure:"USER_DB_PATH" validate:"required"`
	SessionDBPath   string `mapstructure:"SESSION_DB_PATH" validate:"required"`
	UploadDBPath    string `mapstructure:"UPLOAD_DB_PATH" validate:"required"`
	CartDBPath      string `mapstructure:"CART_DB_PATH" validate:"required"`
	PurchaseDBPath  string `mapstructure:"PURCHASE_DB_PATH" validate:"required"`
}
