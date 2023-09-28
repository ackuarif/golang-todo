package configs

import "todo/models"

func initMigrate(){
	DB.AutoMigrate(&models.Todo{})
}