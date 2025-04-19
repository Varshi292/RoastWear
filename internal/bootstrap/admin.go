package bootstrap

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gofiber"
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/tests/tables"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/Varshi292/RoastWear/internal/admin/models"
	"github.com/Varshi292/RoastWear/internal/admin/pages"
	"github.com/gofiber/fiber/v2"
)

func initializeAdmin(app *fiber.App) {
	template.AddComp(chartjs.NewChart())
	eng := engine.Default()
	if err := eng.AddConfigFromYAML("./internal/admin/config.yml").
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}
	eng.HTML("GET", "/admin", pages.GetDashBoard)
	models.Init(eng.SqliteConnection())
}
