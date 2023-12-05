package migrate

import (
	"github.com/logan1x/fiber_go/initializers"
	"github.com/logan1x/fiber_go/models"
)

func init() {
	// Load env variables
	initializers.LoadEnvVariables()
	// Connect to DB
	initializers.ConnectToDB()
}

func Migrate() {
	// Migrate the schema

	initializers.DB.AutoMigrate(&models.User{})
}
