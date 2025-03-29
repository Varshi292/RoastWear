package config

type AppConfig struct {
	Port            string `mapstructure:"PORT" validate:"required"`
	StaticFilesPath string `mapstructure:"STATIC_FILES_PATH" validate:"required"`
	DBPath          string `mapstructure:"DB_PATH" validate:"required"`
}
