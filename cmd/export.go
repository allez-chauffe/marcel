package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/api/db"
	"github.com/Zenika/MARCEL/api/db/users"
	"github.com/Zenika/MARCEL/config"
)

var usersWPassword bool

type userPassword struct {
	users.User
	PasswordHash string `json:"passwordHash"`
	PasswordSalt string `json:"passwordSalt"`
}

func init() {
	exportCmd.PersistentFlags().String("dbFile", config.Config.DBFile, "Database file name")
	viper.BindPFlag("dbFile", exportCmd.PersistentFlags().Lookup("dbFile"))

	Marcel.AddCommand(exportCmd)

	usersCmd.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(usersCmd)
}

var (
	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports data from Marcel's database",
		Args:  cobra.NoArgs,
	}

	usersCmd = &cobra.Command{
		Use:   "users FILE",
		Short: "Exports users from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := db.OpenRO(); err != nil {
				return err
			}
			defer db.Close()

			f, err := os.Create(args[0])
			if err != nil {
				return err
			}
			defer f.Close()

			users, err := users.List()
			if err != nil {
				return err
			}

			if usersWPassword {
				var usersPassword = make([]userPassword, 0, len(users))

				for _, user := range users {
					usersPassword = append(usersPassword, userPassword{
						User:         user,
						PasswordHash: user.PasswordHash,
						PasswordSalt: user.PasswordSalt,
					})
				}

				return json.NewEncoder(f).Encode(usersPassword)
			}

			return json.NewEncoder(f).Encode(users)
		},
	}
)
