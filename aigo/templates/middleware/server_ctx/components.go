package server_ctx

import (
	"chocolate/configs"
	"chocolate/init/database"

	"gorm.io/gorm"
)

type Components struct {
	GDB *gorm.DB
}

func initComponents(conf configs.Conf) *Components {
	gdb := database.InitGROMClientForMySQL(conf)
	return &Components{
		GDB: gdb,
	}

}
