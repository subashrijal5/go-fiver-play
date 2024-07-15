package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/subashrijal5/go-fiber-boilerplate/app/model"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/database"
)

var MigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDB()
		fmt.Println("Starting Migration")
		migrateErr := db.AutoMigrate(&model.User{}, &model.Role{})
		if migrateErr != nil {
			log.Panicf("failed database migration. error: %v", migrateErr)
		}
		fmt.Println("Completed migration")
	},
}

func init() {
	rootCommand.AddCommand(MigrateCommand)
}
