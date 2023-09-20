package initializers

import "github.com/joacof98/epicprompts/pkg/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Prompt{})
	DB.AutoMigrate(&models.Response{})
}
